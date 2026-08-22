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
		field.String("slug").Unique(),
		field.Text("description").Optional(),
		field.String("thumbnail").Optional().Nillable(),
		field.Float("price").Default(0),
		field.Enum("level").Values("beginner", "intermediate", "advanced").Default("beginner"),
		field.Enum("status").Values("draft", "published", "archived").Default("draft"),
		field.Int("total_lessons").Default(0),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Time("deleted_at").Optional().Nillable(),
	}
}

func (Courses) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("Categories", Categories.Type).Ref("Courses").Field("category_id").Unique().Required(),
		edge.To("Sections", Sections.Type),
		edge.To("Lessons", Lessons.Type),
		edge.To("Quizzes", Quizzes.Type),
		edge.To("Enrollments", Enrollments.Type),
		edge.To("LessonProgresses", LessonProgresses.Type),
		edge.To("Certificates", Certificates.Type),
	}
}
