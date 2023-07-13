package main

import (
	"fmt"
	"net_parser/config"
	"net_parser/pkg/parser"
	"time"
)

func main() {

	db := config.ConnectDB()
	defer config.DisconnectDB(db)

	config.InitTables()

	start := time.Now()
	filename := "example.utm"
	fmt.Println("Идёт загрузка данных...")
	err := parser.ParseBinaryData(db, filename)
	if (err != nil) && (err.Error() != "EOF") {
		fmt.Println(err.Error())
	}

	timer := time.Since(start)
	fmt.Println(timer)
}
