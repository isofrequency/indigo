// Code generated by cmd/lexgen (see Makefile's lexgen); DO NOT EDIT.

package atproto

// schema: com.atproto.admin.disableAccountInvites

import (
	"context"

	"github.com/bluesky-social/indigo/lex/util"
	"github.com/bluesky-social/indigo/xrpc"
)

// AdminDisableAccountInvites_Input is the input argument to a com.atproto.admin.disableAccountInvites call.
type AdminDisableAccountInvites_Input struct {
	Account util.FormatDID `json:"account" cborgen:"account"`
}

// AdminDisableAccountInvites calls the XRPC method "com.atproto.admin.disableAccountInvites".
func AdminDisableAccountInvites(ctx context.Context, c *xrpc.Client, input *AdminDisableAccountInvites_Input) error {
	if err := c.Do(ctx, xrpc.Procedure, "application/json", "com.atproto.admin.disableAccountInvites", nil, input, nil); err != nil {
		return err
	}

	return nil
}
