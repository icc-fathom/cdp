// Code generated by cdpgen. DO NOT EDIT.

package log

import (
	"github.com/icc-fathom/cdp/protocol/network"
	"github.com/icc-fathom/cdp/protocol/runtime"
)

// Entry Log entry.
type Entry struct {
	// Source Log entry source.
	//
	// Values: "xml", "javascript", "network", "storage", "appcache", "rendering", "security", "deprecation", "worker", "violation", "intervention", "recommendation", "other".
	Source string `json:"source"`
	// Level Log entry severity.
	//
	// Values: "verbose", "info", "warning", "error".
	Level string `json:"level"`
	Text  string `json:"text"` // Logged text.
	// Category
	//
	// Values: "cors".
	Category         *string                `json:"category,omitempty"`
	Timestamp        runtime.Timestamp      `json:"timestamp"`                  // Timestamp when this entry was added.
	URL              *string                `json:"url,omitempty"`              // URL of the resource if known.
	LineNumber       *int                   `json:"lineNumber,omitempty"`       // Line number in the resource.
	StackTrace       *runtime.StackTrace    `json:"stackTrace,omitempty"`       // JavaScript stack trace.
	NetworkRequestID *network.RequestID     `json:"networkRequestId,omitempty"` // Identifier of the network request associated with this entry.
	WorkerID         *string                `json:"workerId,omitempty"`         // Identifier of the worker associated with this entry.
	Args             []runtime.RemoteObject `json:"args,omitempty"`             // Call arguments.
}

// ViolationSetting Violation configuration setting.
type ViolationSetting struct {
	// Name Violation type.
	//
	// Values: "longTask", "longLayout", "blockedEvent", "blockedParser", "discouragedAPIUse", "handler", "recurringHandler".
	Name      string  `json:"name"`
	Threshold float64 `json:"threshold"` // Time threshold to trigger upon.
}
