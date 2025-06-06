// Code generated by cdpgen. DO NOT EDIT.

package webauthn

import (
	"github.com/icc-fathom/cdp/rpcc"
)

// CredentialAddedClient is a client for CredentialAdded events. Triggered
// when a credential is added to an authenticator.
type CredentialAddedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*CredentialAddedReply, error)
	rpcc.Stream
}

// CredentialAddedReply is the reply for CredentialAdded events.
type CredentialAddedReply struct {
	AuthenticatorID AuthenticatorID `json:"authenticatorId"` // No description.
	Credential      Credential      `json:"credential"`      // No description.
}

// CredentialAssertedClient is a client for CredentialAsserted events.
// Triggered when a credential is used in a webauthn assertion.
type CredentialAssertedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*CredentialAssertedReply, error)
	rpcc.Stream
}

// CredentialAssertedReply is the reply for CredentialAsserted events.
type CredentialAssertedReply struct {
	AuthenticatorID AuthenticatorID `json:"authenticatorId"` // No description.
	Credential      Credential      `json:"credential"`      // No description.
}
