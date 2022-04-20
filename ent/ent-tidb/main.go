package main

import (
	"context"
	"ent-tidb/ent"
	"entgo.io/ent/dialect/sql/schema"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
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
	user1, err := client.User.Create().SetAge(30).SetName("Mashraki").Save(ctx)
	if err != nil {
		return fmt.Errorf("creating user: %w", err)
	}
	log.Println("user1:", user1)
	expired, err := time.Parse(time.RFC3339, "2019-12-08T15:04:05Z")
	if err != nil {
		return err
	}
	card1, err := client.Card.Create().SetOwner(user1).SetNumber("1024").SetExpired(expired).Save(ctx)
	if err != nil {
		return fmt.Errorf("creating card: %w", err)
	}
	log.Println("card1:", card1)
	card, err := user1.QueryCard().Only(ctx)
	if err != nil {
		return fmt.Errorf("querying card: %w", err)
	}
	log.Println("card:", card)
	owner, err := card1.QueryOwner().Only(ctx)
	if err != nil {
		return fmt.Errorf("querying owner: %w", err)
	}
	log.Println("owner:", owner)
	return nil
}
