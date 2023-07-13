package utils

import (
	"log"
	"net"
	"net_parser/pkg/models"
	"os"
)

func GetIP(data models.Data) (IP models.IP) {
	IP.Source = net.IPv4(byte(data.SourceIP>>24), byte(data.SourceIP>>16), byte(data.SourceIP>>8), byte(data.SourceIP)).String()
	IP.Destination = net.IPv4(byte(data.DestinationIP>>24), byte(data.DestinationIP>>16), byte(data.DestinationIP>>8), byte(data.DestinationIP)).String()
	IP.NFSource = net.IPv4(byte(data.NexthopIP>>24), byte(data.NexthopIP>>16), byte(data.NfSourceIP>>8), byte(data.NfSourceIP)).String()

	return IP
}

func InitFile() *os.File {
	outputFile, err := os.Create("data.txt")
	if err != nil {
		log.Fatal(err)
	}

	return outputFile
}

func SaveToFile(outputFile *os.File, data string) {
	_, err := outputFile.WriteString(data)
	if err != nil {
		log.Fatal(err)
	}
}
