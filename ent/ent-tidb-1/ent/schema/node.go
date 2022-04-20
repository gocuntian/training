package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Node holds the schema definition for the Node entity.
type Node struct {
	ent.Schema
}

// Fields of the Node.
func (Node) Fields() []ent.Field {
	return []ent.Field{
		field.Int("value"),
	}
}

// Edges of the Node.
func (Node) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("next", Node.Type).Unique().From("prev").Unique(),
	}
}

//CREATE TABLE `nodes` (
//`id` bigint(20) NOT NULL AUTO_INCREMENT,
//`value` bigint(20) NOT NULL,
//`node_next` bigint(20) DEFAULT NULL,
//PRIMARY KEY (`id`) /*T![clustered_index] CLUSTERED */,
//UNIQUE KEY `node_next` (`node_next`),
//CONSTRAINT `nodes_nodes_next` FOREIGN KEY (`node_next`) REFERENCES `nodes` (`id`) ON DELETE SET NULL
//) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin AUTO_INCREMENT=2000001
