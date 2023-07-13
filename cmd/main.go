package main

import (
	"fmt"
	"net_parser/config"
	"net_parser/pkg/models"
	"net_parser/pkg/parser"
	"net_parser/pkg/repository"
	"net_parser/pkg/utils"
	"time"
)

func main() {

	db := config.ConnectDB()
	defer config.DisconnectDB(db)

	config.InitTables()
	file := utils.InitFile()

	start := time.Now()
	var input models.FiltredData
	filename := "ip60.utm"
	fmt.Println("Идёт загрузка данных...")
	err := parser.ParseBinaryData(db, filename)
	if (err != nil) && (err.Error() != "EOF") {
		fmt.Println(err.Error())
	}

	timer := time.Since(start)
	fmt.Println(timer)

	for {
		fmt.Println("Введите нужные параметры")
		fmt.Print("account_id: ")
		fmt.Scan(&input.AccountID)
		fmt.Print("source:")
		fmt.Scan(&input.Source)
		fmt.Print("destination: ")
		fmt.Scan(&input.Destination)
		fmt.Print("tclass: ")
		fmt.Scan(&input.Tclass)

		data, err := repository.GetFiltredProducts(input)
		if err != nil {
			fmt.Println(err)
		}

		for _, d := range data {
			temp := fmt.Sprintf("DateTime: %v; AccountID: %v; Source: %s; Destination: %s; Tclass: %v; Packets: %v; Sport: %v; Dport: %v;  Bytes: %v; nfSource: %s\n",
				d.DataTime, d.AccountID, d.Source, d.Destination, d.Tclass, d.Packets, d.Sport, d.Dport, d.Bytes, d.NFSource)
			utils.SaveToFile(file, temp)
		}

	}

}
