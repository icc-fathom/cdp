// Code generated by cdpgen. DO NOT EDIT.

// Package audits implements the Audits domain. Audits domain allows
// investigation of page violations and possible improvements.
package audits

import (
	"context"

	"github.com/icc-fathom/cdp/protocol/internal"
	"github.com/icc-fathom/cdp/rpcc"
)

// domainClient is a client for the Audits domain. Audits domain allows
// investigation of page violations and possible improvements.
type domainClient struct{ conn *rpcc.Conn }

// NewClient returns a client for the Audits domain with the connection set to conn.
func NewClient(conn *rpcc.Conn) *domainClient {
	return &domainClient{conn: conn}
}

// GetEncodedResponse invokes the Audits method. Returns the response body and
// size if it were re-encoded with the specified settings. Only applies to
// images.
func (d *domainClient) GetEncodedResponse(ctx context.Context, args *GetEncodedResponseArgs) (reply *GetEncodedResponseReply, err error) {
	reply = new(GetEncodedResponseReply)
	if args != nil {
		err = rpcc.Invoke(ctx, "Audits.getEncodedResponse", args, reply, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Audits.getEncodedResponse", nil, reply, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Audits", Op: "GetEncodedResponse", Err: err}
	}
	return
}

// Disable invokes the Audits method. Disables issues domain, prevents further
// issues from being reported to the client.
func (d *domainClient) Disable(ctx context.Context) (err error) {
	err = rpcc.Invoke(ctx, "Audits.disable", nil, nil, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "Audits", Op: "Disable", Err: err}
	}
	return
}

// Enable invokes the Audits method. Enables issues domain, sends the issues
// collected so far to the client by means of the `issueAdded` event.
func (d *domainClient) Enable(ctx context.Context) (err error) {
	err = rpcc.Invoke(ctx, "Audits.enable", nil, nil, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "Audits", Op: "Enable", Err: err}
	}
	return
}

// CheckContrast invokes the Audits method. Runs the contrast check for the
// target page. Found issues are reported using Audits.issueAdded event.
func (d *domainClient) CheckContrast(ctx context.Context, args *CheckContrastArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Audits.checkContrast", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Audits.checkContrast", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Audits", Op: "CheckContrast", Err: err}
	}
	return
}

// CheckFormsIssues invokes the Audits method. Runs the form issues check for
// the target page. Found issues are reported using Audits.issueAdded event.
func (d *domainClient) CheckFormsIssues(ctx context.Context) (reply *CheckFormsIssuesReply, err error) {
	reply = new(CheckFormsIssuesReply)
	err = rpcc.Invoke(ctx, "Audits.checkFormsIssues", nil, reply, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "Audits", Op: "CheckFormsIssues", Err: err}
	}
	return
}

func (d *domainClient) IssueAdded(ctx context.Context) (IssueAddedClient, error) {
	s, err := rpcc.NewStream(ctx, "Audits.issueAdded", d.conn)
	if err != nil {
		return nil, err
	}
	return &issueAddedClient{Stream: s}, nil
}

type issueAddedClient struct{ rpcc.Stream }

// GetStream returns the original Stream for use with cdp.Sync.
func (c *issueAddedClient) GetStream() rpcc.Stream { return c.Stream }

func (c *issueAddedClient) Recv() (*IssueAddedReply, error) {
	event := new(IssueAddedReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "Audits", Op: "IssueAdded Recv", Err: err}
	}
	return event, nil
}
