package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Certificates struct{ ent.Schema }

func (Certificates) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id"),
		field.Int("course_id"),
		field.String("code").Unique(),
		field.Time("issued_at").Default(time.Now),
		field.Time("created_at").Default(time.Now).Immutable(),
	}
}

func (Certificates) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("Users", Users.Type).Ref("Certificates").Field("user_id").Unique().Required(),
		edge.From("Courses", Courses.Type).Ref("Certificates").Field("course_id").Unique().Required(),
	}
}
