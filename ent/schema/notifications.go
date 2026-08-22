package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Notifications struct{ ent.Schema }

func (Notifications) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id"),
		field.String("title"),
		field.Text("content"),
		field.Bool("is_read").Default(false),
		field.Time("created_at").Default(time.Now).Immutable(),
	}
}

func (Notifications) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("Users", Users.Type).Ref("Notifications").Field("user_id").Unique().Required(),
	}
}
