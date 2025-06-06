// Code generated by cdpgen. DO NOT EDIT.

// Package tracing implements the Tracing domain.
package tracing

import (
	"context"

	"github.com/icc-fathom/cdp/protocol/internal"
	"github.com/icc-fathom/cdp/rpcc"
)

// domainClient is a client for the Tracing domain.
type domainClient struct{ conn *rpcc.Conn }

// NewClient returns a client for the Tracing domain with the connection set to conn.
func NewClient(conn *rpcc.Conn) *domainClient {
	return &domainClient{conn: conn}
}

// End invokes the Tracing method. Stop trace events collection.
func (d *domainClient) End(ctx context.Context) (err error) {
	err = rpcc.Invoke(ctx, "Tracing.end", nil, nil, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "Tracing", Op: "End", Err: err}
	}
	return
}

// GetCategories invokes the Tracing method. Gets supported tracing
// categories.
func (d *domainClient) GetCategories(ctx context.Context) (reply *GetCategoriesReply, err error) {
	reply = new(GetCategoriesReply)
	err = rpcc.Invoke(ctx, "Tracing.getCategories", nil, reply, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "Tracing", Op: "GetCategories", Err: err}
	}
	return
}

// RecordClockSyncMarker invokes the Tracing method. Record a clock sync
// marker in the trace.
func (d *domainClient) RecordClockSyncMarker(ctx context.Context, args *RecordClockSyncMarkerArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Tracing.recordClockSyncMarker", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Tracing.recordClockSyncMarker", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Tracing", Op: "RecordClockSyncMarker", Err: err}
	}
	return
}

// RequestMemoryDump invokes the Tracing method. Request a global memory dump.
func (d *domainClient) RequestMemoryDump(ctx context.Context, args *RequestMemoryDumpArgs) (reply *RequestMemoryDumpReply, err error) {
	reply = new(RequestMemoryDumpReply)
	if args != nil {
		err = rpcc.Invoke(ctx, "Tracing.requestMemoryDump", args, reply, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Tracing.requestMemoryDump", nil, reply, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Tracing", Op: "RequestMemoryDump", Err: err}
	}
	return
}

// Start invokes the Tracing method. Start trace events collection.
func (d *domainClient) Start(ctx context.Context, args *StartArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Tracing.start", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Tracing.start", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Tracing", Op: "Start", Err: err}
	}
	return
}

func (d *domainClient) BufferUsage(ctx context.Context) (BufferUsageClient, error) {
	s, err := rpcc.NewStream(ctx, "Tracing.bufferUsage", d.conn)
	if err != nil {
		return nil, err
	}
	return &bufferUsageClient{Stream: s}, nil
}

type bufferUsageClient struct{ rpcc.Stream }

// GetStream returns the original Stream for use with cdp.Sync.
func (c *bufferUsageClient) GetStream() rpcc.Stream { return c.Stream }

func (c *bufferUsageClient) Recv() (*BufferUsageReply, error) {
	event := new(BufferUsageReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "Tracing", Op: "BufferUsage Recv", Err: err}
	}
	return event, nil
}

func (d *domainClient) DataCollected(ctx context.Context) (DataCollectedClient, error) {
	s, err := rpcc.NewStream(ctx, "Tracing.dataCollected", d.conn)
	if err != nil {
		return nil, err
	}
	return &dataCollectedClient{Stream: s}, nil
}

type dataCollectedClient struct{ rpcc.Stream }

// GetStream returns the original Stream for use with cdp.Sync.
func (c *dataCollectedClient) GetStream() rpcc.Stream { return c.Stream }

func (c *dataCollectedClient) Recv() (*DataCollectedReply, error) {
	event := new(DataCollectedReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "Tracing", Op: "DataCollected Recv", Err: err}
	}
	return event, nil
}

func (d *domainClient) TracingComplete(ctx context.Context) (CompleteClient, error) {
	s, err := rpcc.NewStream(ctx, "Tracing.tracingComplete", d.conn)
	if err != nil {
		return nil, err
	}
	return &completeClient{Stream: s}, nil
}

type completeClient struct{ rpcc.Stream }

// GetStream returns the original Stream for use with cdp.Sync.
func (c *completeClient) GetStream() rpcc.Stream { return c.Stream }

func (c *completeClient) Recv() (*CompleteReply, error) {
	event := new(CompleteReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "Tracing", Op: "TracingComplete Recv", Err: err}
	}
	return event, nil
}
