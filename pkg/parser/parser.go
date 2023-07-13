package parser

import (
	"database/sql"
	"encoding/binary"
	"fmt"
	"log"
	"net_parser/pkg/models"
	"net_parser/pkg/utils"
	"os"
	"sync"
)

func parseRecord(record []byte) models.Data {
	var nfRecord models.Data

	nfRecord.DeviceId = uint8(record[0])
	nfRecord.SourceIP = binary.LittleEndian.Uint32(record[1:5])
	nfRecord.DestinationIP = binary.LittleEndian.Uint32(record[5:9])
	nfRecord.NexthopIP = binary.LittleEndian.Uint32(record[9:13])
	nfRecord.Iface = binary.LittleEndian.Uint16(record[13:15])
	nfRecord.Oface = binary.LittleEndian.Uint16(record[15:17])
	nfRecord.Packets = binary.LittleEndian.Uint32(record[17:21])
	nfRecord.Bytes = binary.LittleEndian.Uint32(record[21:25])
	nfRecord.StartTime = binary.LittleEndian.Uint32(record[25:29])
	nfRecord.EndTime = binary.LittleEndian.Uint32(record[29:33])
	nfRecord.Sport = binary.LittleEndian.Uint16(record[33:35])
	nfRecord.Dport = binary.LittleEndian.Uint16(record[35:37])
	nfRecord.TcpFlags = uint8(record[37])
	nfRecord.Proto = uint8(record[38])
	nfRecord.Tos = uint8(record[39])
	nfRecord.SrcAS = binary.LittleEndian.Uint32(record[40:44])
	nfRecord.DstAS = binary.LittleEndian.Uint32(record[44:48])
	nfRecord.SrcMask = uint8(record[48])
	nfRecord.DstMask = uint8(record[49])
	nfRecord.SlinkID = binary.LittleEndian.Uint32(record[50:54])
	nfRecord.AccountID = binary.LittleEndian.Uint32(record[54:58])
	nfRecord.BillingIP = binary.LittleEndian.Uint32(record[58:62])
	nfRecord.Tclass = binary.LittleEndian.Uint32(record[62:66])
	nfRecord.DateTime = binary.LittleEndian.Uint32(record[66:70])
	nfRecord.NfSourceIP = binary.LittleEndian.Uint32(record[70:74])

	return nfRecord
}

func ParseBinaryData(db *sql.DB, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.Seek(175, 0)
	if err != nil {
		log.Fatal(err)
	}

	recordBuf := make([]byte, 74)

	var record models.Data

	dataChan := make(chan models.Data)

	var wg sync.WaitGroup

	for i := 0; i < 8; i++ {
		wg.Add(1)
		go utils.SaveToDatabase(db, dataChan, &wg)
	}

	for {
		_, err := file.Read(recordBuf)
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			log.Fatal(err)
		}

		record = parseRecord(recordBuf)

		dataChan <- record

	}

	close(dataChan)
	wg.Wait()

	fmt.Println("Загрузка данных завершена")

	return nil
}
