// Code generated by cdpgen. DO NOT EDIT.

package indexeddb

import (
	"github.com/icc-fathom/cdp/protocol/runtime"
)

// DatabaseWithObjectStores Database with an array of object stores.
type DatabaseWithObjectStores struct {
	Name         string        `json:"name"`         // Database name.
	Version      float64       `json:"version"`      // Database version (type is not 'integer', as the standard requires the version number to be 'unsigned long long')
	ObjectStores []ObjectStore `json:"objectStores"` // Object stores in this database.
}

// ObjectStore Object store.
type ObjectStore struct {
	Name          string             `json:"name"`          // Object store name.
	KeyPath       KeyPath            `json:"keyPath"`       // Object store key path.
	AutoIncrement bool               `json:"autoIncrement"` // If true, object store has auto increment flag set.
	Indexes       []ObjectStoreIndex `json:"indexes"`       // Indexes in this object store.
}

// ObjectStoreIndex Object store index.
type ObjectStoreIndex struct {
	Name       string  `json:"name"`       // Index name.
	KeyPath    KeyPath `json:"keyPath"`    // Index key path.
	Unique     bool    `json:"unique"`     // If true, index is unique.
	MultiEntry bool    `json:"multiEntry"` // If true, index allows multiple entries for a key.
}

// Key Key.
type Key struct {
	// Type Key type.
	//
	// Values: "number", "string", "date", "array".
	Type   string   `json:"type"`
	Number *float64 `json:"number,omitempty"` // Number value.
	String *string  `json:"string,omitempty"` // String value.
	Date   *float64 `json:"date,omitempty"`   // Date value.
	Array  []Key    `json:"array,omitempty"`  // Array value.
}

// KeyRange Key range.
type KeyRange struct {
	Lower     *Key `json:"lower,omitempty"` // Lower bound.
	Upper     *Key `json:"upper,omitempty"` // Upper bound.
	LowerOpen bool `json:"lowerOpen"`       // If true lower bound is open.
	UpperOpen bool `json:"upperOpen"`       // If true upper bound is open.
}

// DataEntry Data entry.
type DataEntry struct {
	Key        runtime.RemoteObject `json:"key"`        // Key object.
	PrimaryKey runtime.RemoteObject `json:"primaryKey"` // Primary key object.
	Value      runtime.RemoteObject `json:"value"`      // Value object.
}

// KeyPath Key path.
type KeyPath struct {
	// Type Key path type.
	//
	// Values: "null", "string", "array".
	Type   string   `json:"type"`
	String *string  `json:"string,omitempty"` // String value.
	Array  []string `json:"array,omitempty"`  // Array value.
}
