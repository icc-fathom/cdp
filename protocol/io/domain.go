// Code generated by cdpgen. DO NOT EDIT.

// Package io implements the IO domain. Input/Output operations for streams
// produced by DevTools.
package io

import (
	"context"

	"github.com/icc-fathom/cdp/protocol/internal"
	"github.com/icc-fathom/cdp/rpcc"
)

// domainClient is a client for the IO domain. Input/Output operations for
// streams produced by DevTools.
type domainClient struct{ conn *rpcc.Conn }

// NewClient returns a client for the IO domain with the connection set to conn.
func NewClient(conn *rpcc.Conn) *domainClient {
	return &domainClient{conn: conn}
}

// Close invokes the IO method. Close the stream, discard any temporary
// backing storage.
func (d *domainClient) Close(ctx context.Context, args *CloseArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "IO.close", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "IO.close", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "IO", Op: "Close", Err: err}
	}
	return
}

// Read invokes the IO method. Read a chunk of the stream
func (d *domainClient) Read(ctx context.Context, args *ReadArgs) (reply *ReadReply, err error) {
	reply = new(ReadReply)
	if args != nil {
		err = rpcc.Invoke(ctx, "IO.read", args, reply, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "IO.read", nil, reply, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "IO", Op: "Read", Err: err}
	}
	return
}

// ResolveBlob invokes the IO method. Return UUID of Blob object specified by
// a remote object id.
func (d *domainClient) ResolveBlob(ctx context.Context, args *ResolveBlobArgs) (reply *ResolveBlobReply, err error) {
	reply = new(ResolveBlobReply)
	if args != nil {
		err = rpcc.Invoke(ctx, "IO.resolveBlob", args, reply, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "IO.resolveBlob", nil, reply, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "IO", Op: "ResolveBlob", Err: err}
	}
	return
}
