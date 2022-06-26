package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
    "entgo.io/ent/schema/edge"
)

// Image holds the schema definition for the Image entity.
type Image struct {
	ent.Schema
}

// Fields of the Image.
func (Image) Fields() []ent.Field {
	return []ent.Field{
		field.String("image").NotEmpty(),
		field.String("longitude").NotEmpty(),
		field.String("latitude").NotEmpty(),
	}
}

// Edges of the Image.
func (Image) Edges() []ent.Edge {
	return []ent.Edge{
        edge.From("owner", User.Type).
            Ref("image").
            Unique(),
    }
}
