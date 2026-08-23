package schema

import (
	"time"

	softdelete "lms-api/ent/soft-delete"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Lessons struct{ ent.Schema }

func (Lessons) Fields() []ent.Field {
	return []ent.Field{
		field.Int("section_id"),
		field.Int("course_id"),
		field.String("title"),
		field.Text("content").Optional(),
		field.String("video_url").Optional().Nillable(),
		field.Int("duration_minutes").Default(0),
		field.Int("sort_order").Default(0),
		field.Bool("is_preview").Default(false),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (Lessons) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("Sections", Sections.Type).Ref("Lessons").Field("section_id").Unique().Required(),
		edge.From("Courses", Courses.Type).Ref("Lessons").Field("course_id").Unique().Required(),
		edge.To("LessonProgresses", LessonProgresses.Type),
		edge.To("Quizzes", Quizzes.Type),
	}
}

func (Lessons) Mixin() []ent.Mixin {
	return []ent.Mixin{
		softdelete.SoftDeleteMixin{},
	}
}
