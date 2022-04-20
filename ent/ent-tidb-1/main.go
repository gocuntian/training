package main

import (
	"context"
	"ent-tidb-1/ent"
	"ent-tidb-1/ent/node"
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
	head, err := client.Node.Create().SetValue(1).Save(ctx)
	if err != nil {
		return fmt.Errorf("creating the head: %w", err)
	}
	curr := head
	//linked-list: 1<->2<->3<->4<->5.
	for i := 0; i < 4; i++ {
		curr, err = client.Node.Create().SetValue(curr.Value + 1).SetPrev(curr).Save(ctx)
		if err != nil {
			return err
		}
	}

	for curr = head; curr != nil; curr = curr.QueryNext().FirstX(ctx) {
		fmt.Printf(" %d", curr.Value)
	}

	tail, err := client.Node.Query().Where(node.Not(node.HasNext())).Only(ctx)
	if err != nil {
		return fmt.Errorf("getting the tail of the list: %v", tail)
	}

	tail, err = tail.Update().SetNext(head).Save(ctx)
	if err != nil {
		return err
	}

	prev, err := head.QueryPrev().Only(ctx)
	if err != nil {
		return fmt.Errorf("getting head's prev: %w", err)
	}
	fmt.Printf("\n%v", prev.Value == tail.Value)
	return nil
}

//root, err := client.Node.Create().SetValue(2).Save(ctx)
//if err != nil {
//return fmt.Errorf("creating the root: %w", err)
//}
//// Add additional nodes to the tree:
////
////       2
////     /   \
////    1     4
////        /   \
////       3     5
////
//n1 := client.Node.Create().SetValue(1).SetParent(root).SaveX(ctx)
//n4 := client.Node.Create().SetValue(4).SetParent(root).SaveX(ctx)
//n3 := client.Node.Create().SetValue(3).SetParent(n4).SaveX(ctx)
//n5 := client.Node.Create().SetValue(5).SetParent(n4).SaveX(ctx)
//fmt.Println("Tree leafs", []int{n1.Value, n3.Value, n5.Value})
//return nil
