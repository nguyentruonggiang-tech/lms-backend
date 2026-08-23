package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type Notifications struct{ ent.Schema }

func (Notifications) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id"),
		field.String("title"),
		field.String("content"),
		field.Bool("is_read").Default(false),
		field.Time("created_at").Default(time.Now).Immutable(),
	}
}
