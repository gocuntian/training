package main

import (
	"context"
	"ent-tidb-2/ent"
	"ent-tidb-2/ent/node"
	"entgo.io/ent/dialect/sql/schema"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	client, err := ent.Open("mysql", "root@tcp(localhost:4000)/test?parseTime=true")
	if err != nil {
		log.Fatalf("failed opening connection to tidb: %v", err)
	}
	defer client.Close()

	ctx := context.Background()
	if err := client.Schema.Create(ctx, schema.WithAtlas(true)); err != nil {
		log.Fatalf("failed printing schema changes: %v", err)
	}

	if err := Do(ctx, client); err != nil {
		log.Fatal(err)
	}
}

func Do(ctx context.Context, client *ent.Client) error {
	root, err := client.Node.Create().SetValue(2).Save(ctx)
	if err != nil {
		return fmt.Errorf("creating the root: %w", err)
	}
	// Add additional nodes to the tree:
	//
	//       2
	//     /   \
	//    1     4
	//        /   \
	//       3     5
	//
	n1 := client.Node.Create().SetValue(1).SetParent(root).SaveX(ctx)
	n4 := client.Node.Create().SetValue(4).SetParent(root).SaveX(ctx)
	n3 := client.Node.Create().SetValue(3).SetParent(n4).SaveX(ctx)
	n5 := client.Node.Create().SetValue(5).SetParent(n4).SaveX(ctx)
	fmt.Println("Tree leafs", []int{n1.Value, n3.Value, n5.Value})
	//Tree leafs [1 3 5]
	ints := client.Node.Query().Where(node.Not(node.HasChildren())).Order(ent.Asc(node.FieldValue)).GroupBy(node.FieldValue).IntsX(ctx)
	fmt.Println(ints)
	//[1 3 5]

	orphan := client.Node.Query().Where(node.Not(node.HasParent())).OnlyX(ctx)
	fmt.Println(orphan)
	//Node(id=1, value=2)
	return nil
}
