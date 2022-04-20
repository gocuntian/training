package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Card holds the schema definition for the Card entity.
type Card struct {
	ent.Schema
}

// Fields of the Card.
func (Card) Fields() []ent.Field {
	return []ent.Field{
		field.Time("expired"),
		field.String("number"),
	}
}

// Edges of the Card.
func (Card) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).Ref("card").Unique().Required(),
	}
}

//
//CREATE TABLE `cards` (
//`id` bigint(20) NOT NULL AUTO_INCREMENT,
//`expired` timestamp NULL DEFAULT NULL,
//`number` varchar(255) NOT NULL,
//`user_card` bigint(20) NOT NULL,
//PRIMARY KEY (`id`) /*T![clustered_index] CLUSTERED */,
//UNIQUE KEY `user_card` (`user_card`),
//CONSTRAINT `cards_users_card` FOREIGN KEY (`user_card`) REFERENCES `users` (`id`) ON DELETE NO ACTION
//) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin AUTO_INCREMENT=2000001
