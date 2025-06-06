// Code generated by cdpgen. DO NOT EDIT.

package audits

import (
	"github.com/icc-fathom/cdp/rpcc"
)

// IssueAddedClient is a client for IssueAdded events.
type IssueAddedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*IssueAddedReply, error)
	rpcc.Stream
}

// IssueAddedReply is the reply for IssueAdded events.
type IssueAddedReply struct {
	Issue InspectorIssue `json:"issue"` // No description.
}
