package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Default(time.Now).
			Annotations(entsql.DefaultExprs(map[string]string{
				dialect.Postgres: "NOW()",
				dialect.MySQL:    "NOW()",
				dialect.SQLite:   "CURRENT_TIMESTAMP",
			})),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
