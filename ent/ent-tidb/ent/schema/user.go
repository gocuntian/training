package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("age"),
		field.String("name"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("card", Card.Type).Unique(),
	}
}

//CREATE TABLE `users` (
//`id` bigint(20) NOT NULL AUTO_INCREMENT,
//`age` bigint(20) NOT NULL,
//`name` varchar(255) NOT NULL,
//PRIMARY KEY (`id`) /*T![clustered_index] CLUSTERED */
//) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin AUTO_INCREMENT=2000001
