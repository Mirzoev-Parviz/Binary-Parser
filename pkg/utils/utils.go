package utils

import (
	"database/sql"
	"log"
	"net"
	"net_parser/pkg/models"
	"sync"
)

func SaveToDatabase(db *sql.DB, dataChan <-chan models.Data, wg *sync.WaitGroup) {
	defer wg.Done()

	stmt, err := db.Prepare("INSERT INTO netflow VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	batchSize := 1000
	batch := make([]models.Data, 0, batchSize)

	for data := range dataChan {
		batch = append(batch, data)

		if len(batch) >= batchSize {
			tx, err := db.Begin()
			if err != nil {
				log.Fatal(err)
			}

			for _, data := range batch {
				ip := GetIP(data)
				_, err := tx.Stmt(stmt).Exec(data.AccountID, ip.Source, ip.Destination, data.Packets,
					data.Bytes, data.Sport, data.Dport, data.Proto, data.Tclass, data.DateTime, ip.NFSource)
				if err != nil {
					log.Fatal(err)
				}
			}

			if err := tx.Commit(); err != nil {
				log.Fatal(err)
			}

			batch = batch[:0]
		}
	}

	if len(batch) > 0 {
		tx, err := db.Begin()
		if err != nil {
			log.Fatal(err)
		}

		for _, data := range batch {
			ip := GetIP(data)
			_, err := tx.Stmt(stmt).Exec(data.AccountID, ip.Source, ip.Destination, data.Packets,
				data.Bytes, data.Sport, data.Dport, data.Proto, data.Tclass, data.DateTime, ip.NFSource)
			if err != nil {
				log.Fatal(err)
			}
		}

		if err := tx.Commit(); err != nil {
			log.Fatal(err)
		}
	}

	db.Close()
}

func GetIP(data models.Data) (IP models.IP) {
	IP.Source = net.IPv4(byte(data.SourceIP>>24), byte(data.SourceIP>>16), byte(data.SourceIP>>8), byte(data.SourceIP)).String()
	IP.Destination = net.IPv4(byte(data.DestinationIP>>24), byte(data.DestinationIP>>16), byte(data.DestinationIP>>8), byte(data.DestinationIP)).String()
	IP.NFSource = net.IPv4(byte(data.NexthopIP>>24), byte(data.NexthopIP>>16), byte(data.NfSourceIP>>8), byte(data.NfSourceIP)).String()

	return IP
}
