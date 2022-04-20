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
		edge.To("children", Node.Type).From("parent").Unique(),
	}
}

//CREATE TABLE `nodes` (
//`id` bigint(20) NOT NULL AUTO_INCREMENT,
//`value` bigint(20) NOT NULL,
//`node_children` bigint(20) DEFAULT NULL,
//PRIMARY KEY (`id`) /*T![clustered_index] CLUSTERED */,
//CONSTRAINT `nodes_nodes_children` FOREIGN KEY (`node_children`) REFERENCES `nodes` (`id`) ON DELETE SET NULL
//) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin AUTO_INCREMENT=2000001
