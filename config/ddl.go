package config

import "log"

const (
	NetFlow = `
	CREATE TABLE IF NOT EXISTS netflow (
		account_id BIGINT NOT NULL,
		source TEXT NOT NULL,
		destination TEXT NOT NULL,
		packets BIGINT NOT NULL,
		bytes BIGINT NOT NULL,
		sport BIGINT NOT NULL,
		dport BIGINT NOT NULL,
		proto VARCHAR(255) NOT NULL,
		tclass TEXT NOT NULL,
		data_time TEXT NOT NULL,
		nf_source TEXT NOT NULL
	)
	`
)

func InitTables() {
	_, err := DB.Exec(NetFlow)
	if err != nil {
		log.Fatalf("%s", err.Error())
	}
}
