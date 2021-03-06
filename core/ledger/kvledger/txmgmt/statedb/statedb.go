/*
Copyright IBM Corp. 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package statedb

import (
	"sort"

	"github.com/hyperledger/fabric/core/ledger/kvledger/txmgmt/version"
	"github.com/hyperledger/fabric/core/ledger/util"
)

// VersionedDBProvider provides an instance of an versioned DB
type VersionedDBProvider interface {
	// GetDBHandle returns a handle to a VersionedDB
	GetDBHandle(id string) (VersionedDB, error)
	// Close closes all the VersionedDB instances and releases any resources held by VersionedDBProvider
	Close()
}

// VersionedDB lists methods that a db is supposed to implement
type VersionedDB interface {
	// GetState gets the value for given namespace and key. For a chaincode, the namespace corresponds to the chaincodeId
	GetState(namespace string, key string) (*VersionedValue, error)
	// GetStateMultipleKeys gets the values for multiple keys in a single call
	GetStateMultipleKeys(namespace string, keys []string) ([]*VersionedValue, error)
	// GetStateRangeScanIterator returns an iterator that contains all the key-values between given key ranges.
	// startKey is inclusive
	// endKey is exclusive
	// The returned ResultsIterator contains results of type *VersionedKV
	GetStateRangeScanIterator(namespace string, startKey string, endKey string) (ResultsIterator, error)
	// ExecuteQuery executes the given query and returns an iterator that contains results of type *VersionedKV.
	ExecuteQuery(namespace, query string) (ResultsIterator, error)
	// ExecuteUpdate verifies the given update query and returns true if it can be applied or false otherwise
	// it doesnt execute the update, but returns the current value and key the update modifies
	ExecuteUpdate(namespace, query string) (bool, *VersionedValue, string, error)
	// ApplyUpdates applies the batch to the underlying db.
	// height is the height of the highest transaction in the Batch that
	// a state db implementation is expected to ues as a save point
	ApplyUpdates(batch *UpdateBatch, height *version.Height) error
	// GetLatestSavePoint returns the height of the highest transaction upto which
	// the state db is consistent
	GetLatestSavePoint() (*version.Height, error)
	// ValidateKey tests whether the key is supported by the db implementation.
	// For instance, leveldb supports any bytes for the key while the couchdb supports only valid utf-8 string
	ValidateKey(key string) error
	// Open opens the db
	Open() error
	// Close closes the db
	Close()
}

// CompositeKey encloses Namespace and Key components
type CompositeKey struct {
	Namespace string
	Key       string
}

// VersionedValue encloses value and corresponding version
type VersionedValue struct {
	Value   []byte
	Version *version.Height
}

// VersionedKV encloses key and corresponding VersionedValue
type VersionedKV struct {
	CompositeKey
	VersionedValue
}

// ResultsIterator hepls in iterates over query results
type ResultsIterator interface {
	Next() (QueryResult, error)
	Close()
}

// QueryResult - a general interface for supporting different types of query results. Actual types differ for different queries
type QueryResult interface{}

type VersionedValueIncDelta struct {
	VersionedValue
	IsDelta bool
}

// With delta, its possible that there are multiple updates to a key
// so this holds a ordered list of updates
type nsUpdates struct {
	m map[string][]*VersionedValueIncDelta
}

func newNsUpdates() *nsUpdates {
	return &nsUpdates{make(map[string][]*VersionedValueIncDelta)}
}

// UpdateBatch encloses the details of multiple `updates`
type UpdateBatch struct {
	updates map[string]*nsUpdates
}

// NewUpdateBatch constructs an instance of a Batch
func NewUpdateBatch() *UpdateBatch {
	return &UpdateBatch{make(map[string]*nsUpdates)}
}

// Get returns the VersionedValue for the given namespace and key
func (batch *UpdateBatch) Get(ns string, key string) *VersionedValue {
	nsUpdates, ok := batch.updates[ns]
	if !ok {
		return nil
	}
	vv, ok := nsUpdates.m[key]
	if !ok {
		return nil
	}
	if len(vv) > 1 {
		panic("Multiple updates for key")
	}
	return &vv[0].VersionedValue
}

// Put adds a VersionedKV
func (batch *UpdateBatch) Put(ns string, key string, value []byte, isDelta bool, version *version.Height) {
	if value == nil {
		panic("Nil value not allowed")
	}
	nsUpdates := batch.getOrCreateNsUpdates(ns)
	if !isDelta {
		// any full write will overwrite all previous deltas
		nsUpdates.m[key] = make([]*VersionedValueIncDelta, 0)
	}
	nsUpdates.m[key] = append(nsUpdates.m[key], &VersionedValueIncDelta{VersionedValue{value, version}, isDelta})
}

// Delete deletes a Key and associated value
func (batch *UpdateBatch) Delete(ns string, key string, version *version.Height) {
	nsUpdates := batch.getOrCreateNsUpdates(ns)
	// all deltas,put before delete do not matter
	nsUpdates.m[key] = make([]*VersionedValueIncDelta, 0)
	nsUpdates.m[key] = append(nsUpdates.m[key], &VersionedValueIncDelta{VersionedValue{nil, version}, false})
}

// Exists checks whether the given key exists in the batch
func (batch *UpdateBatch) Exists(ns string, key string) bool {
	nsUpdates, ok := batch.updates[ns]
	if !ok {
		return false
	}
	_, ok = nsUpdates.m[key]
	return ok
}

// GetUpdatedNamespaces returns the names of the namespaces that are updated
func (batch *UpdateBatch) GetUpdatedNamespaces() []string {
	namespaces := make([]string, len(batch.updates))
	i := 0
	for ns := range batch.updates {
		namespaces[i] = ns
		i++
	}
	return namespaces
}

func (batch *UpdateBatch) GetUpdatesRaw(ns string) map[string][]*VersionedValueIncDelta {
	nsUpdates, ok := batch.updates[ns]
	if !ok {
		return nil
	}
	return nsUpdates.m
}

// GetUpdates returns all the updates for a namespace
func (batch *UpdateBatch) GetUpdates(ns string) map[string]*VersionedValue {
	nsUpdates, ok := batch.updates[ns]
	if !ok {
		return nil
	}
	retUpdates := make(map[string]*VersionedValue)
	for k, v := range nsUpdates.m {
		retUpdates[k] = &v[0].VersionedValue
	}

	return retUpdates
}

// GetRangeScanIterator returns an iterator that iterates over keys of a specific namespace in sorted order
// In other word this gives the same functionality over the contents in the `UpdateBatch` as
// `VersionedDB.GetStateRangeScanIterator()` method gives over the contents in the statedb
// This function can be used for querying the contents in the updateBatch before they are committed to the statedb.
// For instance, a validator implementation can used this to verify the validity of a range query of a transaction
// where the UpdateBatch represents the union of the modifications performed by the preceding valid transactions in the same block
// (Assuming Group commit approach where we commit all the updates caused by a block together).
func (batch *UpdateBatch) GetRangeScanIterator(ns string, startKey string, endKey string) ResultsIterator {
	return newNsIterator(ns, startKey, endKey, batch)
}

func (batch *UpdateBatch) getOrCreateNsUpdates(ns string) *nsUpdates {
	nsUpdates := batch.updates[ns]
	if nsUpdates == nil {
		nsUpdates = newNsUpdates()
		batch.updates[ns] = nsUpdates
	}
	return nsUpdates
}

type nsIterator struct {
	ns         string
	nsUpdates  *nsUpdates
	sortedKeys []string
	nextIndex  int
	lastIndex  int
}

func newNsIterator(ns string, startKey string, endKey string, batch *UpdateBatch) *nsIterator {
	nsUpdates, ok := batch.updates[ns]
	if !ok {
		return &nsIterator{}
	}
	sortedKeys := util.GetSortedKeys(nsUpdates.m)
	var nextIndex int
	var lastIndex int
	if startKey == "" {
		nextIndex = 0
	} else {
		nextIndex = sort.SearchStrings(sortedKeys, startKey)
	}
	if endKey == "" {
		lastIndex = len(sortedKeys)
	} else {
		lastIndex = sort.SearchStrings(sortedKeys, endKey)
	}
	return &nsIterator{ns, nsUpdates, sortedKeys, nextIndex, lastIndex}
}

// Next gives next key and versioned value. It returns a nil when exhausted
func (itr *nsIterator) Next() (QueryResult, error) {
	if itr.nextIndex >= itr.lastIndex {
		return nil, nil
	}
	key := itr.sortedKeys[itr.nextIndex]
	vv := itr.nsUpdates.m[key]
	itr.nextIndex++
	return &VersionedKV{CompositeKey{itr.ns, key}, VersionedValue{vv[0].VersionedValue.Value, vv[0].VersionedValue.Version}}, nil
}

// Close implements the method from QueryResult interface
func (itr *nsIterator) Close() {
	// do nothing
}
