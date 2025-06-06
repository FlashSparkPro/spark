package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// TreeStatus is the status of a tree node.
type TreeStatus string

const (
	// TreeStatusPending is the status of a tree that the base L1 transaction is not confirmed yet.
	TreeStatusPending TreeStatus = "PENDING"
	// TreeStatusAvailable is the status of a tree that the base L1 transaction is confirmed.
	TreeStatusAvailable TreeStatus = "AVAILABLE"
	// TreeStatusExited is the status of a tree that has exited.
	TreeStatusExited TreeStatus = "EXITED"
)

// Values returns the values of the tree node status.
func (TreeStatus) Values() []string {
	return []string{
		string(TreeStatusPending),
		string(TreeStatusAvailable),
		string(TreeStatusExited),
	}
}

// Tree is the schema for the trees table.
type Tree struct {
	ent.Schema
}

// Mixin is the mixin for the trees table.
func (Tree) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Fields are the fields for the trees table.
func (Tree) Fields() []ent.Field {
	return []ent.Field{
		field.Bytes("owner_identity_pubkey").NotEmpty(),
		field.Enum("status").GoType(TreeStatus("")),
		field.Enum("network").GoType(Network("")),
		field.Bytes("base_txid").NotEmpty(),
		field.Int16("vout").NonNegative(),
	}
}

// Edges are the edges for the trees table.
func (Tree) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("root", TreeNode.Type).
			Unique(),
		edge.From("nodes", TreeNode.Type).Ref("tree"),
	}
}

// Indexes are the indexes for the trees table.
func (Tree) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("status"),
		index.Fields("network"),
		index.Fields("base_txid", "vout").Unique(),
	}
}
