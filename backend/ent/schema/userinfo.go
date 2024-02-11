package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
	)

// UserInfo holds the schema definition for the UserInfo entity.
type UserInfo struct {
	ent.Schema
}

// Fields of the UserInfo.
func (UserInfo) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id").
		    Positive(),
		field.String("title").
			NotEmpty(),
		field.String("body").
			NotEmpty(),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		}
}

// Edges of the UserInfo.
func (UserInfo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("userinfo").
			Field("user_id").
			Unique().
			Required(),
		}
}
