// Code generated by cmd/lexgen (see Makefile's lexgen); DO NOT EDIT.

package bsky

// schema: app.bsky.graph.block

import (
	"github.com/bluesky-social/indigo/lex/util"
)

func init() {
	util.RegisterType("app.bsky.graph.block", &GraphBlock{})
} //
// RECORDTYPE: GraphBlock
type GraphBlock struct {
	LexiconTypeID string         `json:"$type,const=app.bsky.graph.block" cborgen:"$type,const=app.bsky.graph.block"`
	CreatedAt     string         `json:"createdAt" cborgen:"createdAt"`
	Subject       util.FormatDID `json:"subject" cborgen:"subject"`
}
