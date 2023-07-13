package repository

import (
	"database/sql"
	"log"
	"net_parser/config"
	"net_parser/pkg/models"
	"net_parser/pkg/utils"
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
				ip := utils.GetIP(data)
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
			ip := utils.GetIP(data)
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

func GetFiltredProducts(fData models.FiltredData) (Data []models.FiltredData, err error) {
	rows, err := config.DB.Query("SELECT * FROM netflow WHERE account_id = $1 AND tclass = $2 AND source = $3 AND  destination = $4",
		fData.AccountID, fData.Tclass, fData.Source, fData.Destination)
	if err != nil {
		return []models.FiltredData{}, err
	}
	defer rows.Close()

	for rows.Next() {
		t := models.FiltredData{}
		err = rows.Scan(&t.AccountID, &t.Source, &t.Destination, &t.Packets, &t.Bytes,
			&t.Sport, &t.Dport, &t.Proto, &t.Tclass, &t.DataTime, &t.NFSource)
		if err != nil {
			return []models.FiltredData{}, err
		}

		Data = append(Data, t)
	}

	return Data, nil
}
