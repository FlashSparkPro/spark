package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// PreimageShare is the schema for the preimage shares table.
type PreimageShare struct {
	ent.Schema
}

// Mixin returns the mixin for the preimage shares table.
func (PreimageShare) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Indexes returns the indexes for the preimage shares table.
func (PreimageShare) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("payment_hash"),
	}
}

// Fields returns the fields for the preimage shares table.
func (PreimageShare) Fields() []ent.Field {
	return []ent.Field{
		field.Bytes("payment_hash").NotEmpty().Immutable().Unique(),
		field.Bytes("preimage_share").NotEmpty().Immutable(),
		field.Int32("threshold").Immutable(),
		field.Bytes("owner_identity_pubkey").NotEmpty().Immutable(),
		field.String("invoice_string").NotEmpty().Immutable(),
	}
}

// Edges returns the edges for the preimage shares table.
func (PreimageShare) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("preimage_request", PreimageRequest.Type).
			Ref("preimage_shares").
			Unique(),
	}
}
