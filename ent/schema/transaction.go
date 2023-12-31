package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Transaction holds the schema definition for the Transaction entity.
type Transaction struct {
	ent.Schema
}

// Fields of the Transaction.
func (Transaction) Fields() []ent.Field {
	return []ent.Field{
		field.Int("amount").Positive(),
		field.String("note").Optional(),
	}
}

// Edges of the Transaction.
func (Transaction) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("belongs_to", Group.Type).
			Ref("transactions").Unique(),
		edge.From("source", User.Type).
			Ref("lent").Unique(),
		edge.To("destination", User.Type).Unique(),
	}
}
