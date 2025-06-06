// Code generated by cdpgen. DO NOT EDIT.

package debugger

import (
	"encoding/json"

	"github.com/icc-fathom/cdp/protocol/runtime"
	"github.com/icc-fathom/cdp/rpcc"
)

// BreakpointResolvedClient is a client for BreakpointResolved events. Fired
// when breakpoint is resolved to an actual script and location.
type BreakpointResolvedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*BreakpointResolvedReply, error)
	rpcc.Stream
}

// BreakpointResolvedReply is the reply for BreakpointResolved events.
type BreakpointResolvedReply struct {
	BreakpointID BreakpointID `json:"breakpointId"` // Breakpoint unique identifier.
	Location     Location     `json:"location"`     // Actual breakpoint location.
}

// PausedClient is a client for Paused events. Fired when the virtual machine
// stopped on breakpoint or exception or any other stop criteria.
type PausedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*PausedReply, error)
	rpcc.Stream
}

// PausedReply is the reply for Paused events.
type PausedReply struct {
	CallFrames []CallFrame `json:"callFrames"` // Call stack the virtual machine stopped on.
	// Reason Pause reason.
	//
	// Values: "ambiguous", "assert", "CSPViolation", "debugCommand", "DOM", "EventListener", "exception", "instrumentation", "OOM", "other", "promiseRejection", "XHR", "step".
	Reason          string              `json:"reason"`
	Data            json.RawMessage     `json:"data,omitempty"`            // Object containing break-specific auxiliary properties.
	HitBreakpoints  []string            `json:"hitBreakpoints,omitempty"`  // Hit breakpoints IDs
	AsyncStackTrace *runtime.StackTrace `json:"asyncStackTrace,omitempty"` // Async stack trace, if any.
	// AsyncStackTraceID Async stack trace, if any.
	//
	// Note: This property is experimental.
	AsyncStackTraceID *runtime.StackTraceID `json:"asyncStackTraceId,omitempty"`
	// AsyncCallStackTraceID is deprecated.
	//
	// Deprecated: Never present, will be removed.
	//
	// Note: This property is experimental.
	AsyncCallStackTraceID *runtime.StackTraceID `json:"asyncCallStackTraceId,omitempty"`
}

// ResumedClient is a client for Resumed events. Fired when the virtual
// machine resumed execution.
type ResumedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*ResumedReply, error)
	rpcc.Stream
}

// ResumedReply is the reply for Resumed events.
type ResumedReply struct {
}

// ScriptFailedToParseClient is a client for ScriptFailedToParse events. Fired
// when virtual machine fails to parse the script.
type ScriptFailedToParseClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*ScriptFailedToParseReply, error)
	rpcc.Stream
}

// ScriptFailedToParseReply is the reply for ScriptFailedToParse events.
type ScriptFailedToParseReply struct {
	ScriptID                runtime.ScriptID           `json:"scriptId"`                          // Identifier of the script parsed.
	URL                     string                     `json:"url"`                               // URL or name of the script parsed (if any).
	StartLine               int                        `json:"startLine"`                         // Line offset of the script within the resource with given URL (for script tags).
	StartColumn             int                        `json:"startColumn"`                       // Column offset of the script within the resource with given URL.
	EndLine                 int                        `json:"endLine"`                           // Last line of the script.
	EndColumn               int                        `json:"endColumn"`                         // Length of the last line of the script.
	ExecutionContextID      runtime.ExecutionContextID `json:"executionContextId"`                // Specifies script creation context.
	Hash                    string                     `json:"hash"`                              // Content hash of the script, SHA-256.
	ExecutionContextAuxData json.RawMessage            `json:"executionContextAuxData,omitempty"` // Embedder-specific auxiliary data likely matching {isDefault: boolean, type: 'default'|'isolated'|'worker', frameId: string}
	SourceMapURL            *string                    `json:"sourceMapURL,omitempty"`            // URL of source map associated with script (if any).
	HasSourceURL            *bool                      `json:"hasSourceURL,omitempty"`            // True, if this script has sourceURL.
	IsModule                *bool                      `json:"isModule,omitempty"`                // True, if this script is ES6 module.
	Length                  *int                       `json:"length,omitempty"`                  // This script length.
	// StackTrace JavaScript top stack frame of where the script parsed
	// event was triggered if available.
	//
	// Note: This property is experimental.
	StackTrace *runtime.StackTrace `json:"stackTrace,omitempty"`
	// CodeOffset If the scriptLanguage is WebAssembly, the code section
	// offset in the module.
	//
	// Note: This property is experimental.
	CodeOffset *int `json:"codeOffset,omitempty"`
	// ScriptLanguage The language of the script.
	//
	// Note: This property is experimental.
	ScriptLanguage ScriptLanguage `json:"scriptLanguage,omitempty"`
	// EmbedderName The name the embedder supplied for this script.
	//
	// Note: This property is experimental.
	EmbedderName *string `json:"embedderName,omitempty"`
}

// ScriptParsedClient is a client for ScriptParsed events. Fired when virtual
// machine parses script. This event is also fired for all known and
// uncollected scripts upon enabling debugger.
type ScriptParsedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*ScriptParsedReply, error)
	rpcc.Stream
}

// ScriptParsedReply is the reply for ScriptParsed events.
type ScriptParsedReply struct {
	ScriptID                runtime.ScriptID           `json:"scriptId"`                          // Identifier of the script parsed.
	URL                     string                     `json:"url"`                               // URL or name of the script parsed (if any).
	StartLine               int                        `json:"startLine"`                         // Line offset of the script within the resource with given URL (for script tags).
	StartColumn             int                        `json:"startColumn"`                       // Column offset of the script within the resource with given URL.
	EndLine                 int                        `json:"endLine"`                           // Last line of the script.
	EndColumn               int                        `json:"endColumn"`                         // Length of the last line of the script.
	ExecutionContextID      runtime.ExecutionContextID `json:"executionContextId"`                // Specifies script creation context.
	Hash                    string                     `json:"hash"`                              // Content hash of the script, SHA-256.
	ExecutionContextAuxData json.RawMessage            `json:"executionContextAuxData,omitempty"` // Embedder-specific auxiliary data likely matching {isDefault: boolean, type: 'default'|'isolated'|'worker', frameId: string}
	// IsLiveEdit True, if this script is generated as a result of the
	// live edit operation.
	//
	// Note: This property is experimental.
	IsLiveEdit   *bool   `json:"isLiveEdit,omitempty"`
	SourceMapURL *string `json:"sourceMapURL,omitempty"` // URL of source map associated with script (if any).
	HasSourceURL *bool   `json:"hasSourceURL,omitempty"` // True, if this script has sourceURL.
	IsModule     *bool   `json:"isModule,omitempty"`     // True, if this script is ES6 module.
	Length       *int    `json:"length,omitempty"`       // This script length.
	// StackTrace JavaScript top stack frame of where the script parsed
	// event was triggered if available.
	//
	// Note: This property is experimental.
	StackTrace *runtime.StackTrace `json:"stackTrace,omitempty"`
	// CodeOffset If the scriptLanguage is WebAssembly, the code section
	// offset in the module.
	//
	// Note: This property is experimental.
	CodeOffset *int `json:"codeOffset,omitempty"`
	// ScriptLanguage The language of the script.
	//
	// Note: This property is experimental.
	ScriptLanguage ScriptLanguage `json:"scriptLanguage,omitempty"`
	// DebugSymbols If the scriptLanguage is WebASsembly, the source of
	// debug symbols for the module.
	//
	// Note: This property is experimental.
	DebugSymbols *DebugSymbols `json:"debugSymbols,omitempty"`
	// EmbedderName The name the embedder supplied for this script.
	//
	// Note: This property is experimental.
	EmbedderName *string `json:"embedderName,omitempty"`
}
