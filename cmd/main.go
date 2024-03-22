package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/haruyama480/go-ent-learning/ent"
)

func main() {
	dsn := "root:root@tcp(mysql:13306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v\n", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatalf("failed to ping db: %v\n", err)
	}
	fmt.Println("db connected!!")

	// sql.DB -> entクライアントを作成
	drv := entsql.OpenDB(dialect.MySQL, db)
	client := ent.NewClient(ent.Driver(drv))
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	fmt.Println("success!!")
}
