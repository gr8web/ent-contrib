package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Node struct {
	ent.Schema
}

func (Node) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Unique().
			Annotations(
				entproto.Field(2),
			),
		field.String("user").
			Unique().
			Annotations(
				entproto.Field(3),
			),
	}
}

func (Node) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("parent", Node.Type).
			Unique().
			From("subinstances").
			Annotations(entproto.Field(4)),
	}
}

func (Node) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(),
	}
}
