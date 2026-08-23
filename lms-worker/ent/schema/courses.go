package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Courses struct{ ent.Schema }

func (Courses) Fields() []ent.Field {
	return []ent.Field{
		field.Int("category_id"),
		field.String("title"),
		field.Text("description").Optional(),
		field.Float("price").Default(0),
		field.Enum("level").Values("beginner", "intermediate", "advanced").Default("beginner"),
		field.Enum("status").Values("draft", "published", "archived").Default("draft"),
		field.Time("created_at").Default(time.Now).Immutable(),
	}
}

func (Courses) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("Categories", Categories.Type).Ref("Courses").Field("category_id").Unique().Required(),
	}
}
