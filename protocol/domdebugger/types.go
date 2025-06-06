// Code generated by cdpgen. DO NOT EDIT.

package domdebugger

import (
	"github.com/icc-fathom/cdp/protocol/dom"
	"github.com/icc-fathom/cdp/protocol/runtime"
)

// DOMBreakpointType DOM breakpoint type.
type DOMBreakpointType string

// DOMBreakpointType as enums.
const (
	DOMBreakpointTypeNotSet            DOMBreakpointType = ""
	DOMBreakpointTypeSubtreeModified   DOMBreakpointType = "subtree-modified"
	DOMBreakpointTypeAttributeModified DOMBreakpointType = "attribute-modified"
	DOMBreakpointTypeNodeRemoved       DOMBreakpointType = "node-removed"
)

func (e DOMBreakpointType) Valid() bool {
	switch e {
	case "subtree-modified", "attribute-modified", "node-removed":
		return true
	default:
		return false
	}
}

func (e DOMBreakpointType) String() string {
	return string(e)
}

// CSPViolationType CSP Violation type.
//
// Note: This type is experimental.
type CSPViolationType string

// CSPViolationType as enums.
const (
	CSPViolationTypeNotSet                     CSPViolationType = ""
	CSPViolationTypeTrustedtypeSinkViolation   CSPViolationType = "trustedtype-sink-violation"
	CSPViolationTypeTrustedtypePolicyViolation CSPViolationType = "trustedtype-policy-violation"
)

func (e CSPViolationType) Valid() bool {
	switch e {
	case "trustedtype-sink-violation", "trustedtype-policy-violation":
		return true
	default:
		return false
	}
}

func (e CSPViolationType) String() string {
	return string(e)
}

// EventListener Object event listener.
type EventListener struct {
	Type            string                `json:"type"`                      // `EventListener`'s type.
	UseCapture      bool                  `json:"useCapture"`                // `EventListener`'s useCapture.
	Passive         bool                  `json:"passive"`                   // `EventListener`'s passive flag.
	Once            bool                  `json:"once"`                      // `EventListener`'s once flag.
	ScriptID        runtime.ScriptID      `json:"scriptId"`                  // Script id of the handler code.
	LineNumber      int                   `json:"lineNumber"`                // Line number in the script (0-based).
	ColumnNumber    int                   `json:"columnNumber"`              // Column number in the script (0-based).
	Handler         *runtime.RemoteObject `json:"handler,omitempty"`         // Event handler function value.
	OriginalHandler *runtime.RemoteObject `json:"originalHandler,omitempty"` // Event original handler function value.
	BackendNodeID   *dom.BackendNodeID    `json:"backendNodeId,omitempty"`   // Node the listener is added to (if any).
}
