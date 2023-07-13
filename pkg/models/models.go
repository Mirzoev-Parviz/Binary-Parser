package models

type Data struct {
	DeviceId      uint8
	SourceIP      uint32
	DestinationIP uint32
	NexthopIP     uint32
	Iface         uint16
	Oface         uint16
	Packets       uint32
	Bytes         uint32
	StartTime     uint32
	EndTime       uint32
	Sport         uint16
	Dport         uint16
	TcpFlags      uint8
	Proto         uint8
	Tos           uint8
	SrcAS         uint32
	DstAS         uint32
	SrcMask       uint8
	DstMask       uint8
	SlinkID       uint32
	AccountID     uint32
	BillingIP     uint32
	Tclass        uint32
	DateTime      uint32
	NfSourceIP    uint32
}

type FiltredData struct {
	AccountID   uint32
	Source      string
	Destination string
	Packets     uint16
	Bytes       uint32
	Sport       uint16
	Dport       uint16
	Proto       uint8
	Tclass      uint32
	DataTime    uint32
	NFSource    string
}

type IP struct {
	Source      string
	Destination string
	NFSource    string
}
