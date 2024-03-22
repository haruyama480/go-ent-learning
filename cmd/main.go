package main

import (
	"context"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/haruyama480/go-ent-learning/ent"
)

func main() {
	// https://zenn.dev/masamiki/articles/83a8db3f132fcb1c48f0
	entOptions := []ent.Option{}

	// 発行されるSQLをロギングするなら
	entOptions = append(entOptions, ent.Debug())

	// サンプルなのでここにハードコーディングしてます。
	mc := mysql.Config{
		User:                 "root",
		Passwd:               "root",
		Net:                  "tcp",
		Addr:                 "127.0.0.1" + ":" + "13309", // localhost or dockerMySQL would be failed
		DBName:               "testdb",
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	datasource := mc.FormatDSN()
	fmt.Println(datasource)
	client, err := ent.Open("mysql", datasource, entOptions...)
	if err != nil {
		log.Fatalf("Error open mysql ent client: %v\n", err)
	}

	defer client.Close()

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
