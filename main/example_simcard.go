package main

import (
	"fmt"

	"github.com/Napat/sscard/sscard"
	"github.com/ebfe/scard"
)

// exampleSimCard ...
func exampleSimCard() {
	// Establish a PC/SC context
	context, err := scard.EstablishContext()
	if err != nil {
		fmt.Println("Error EstablishContext:", err)
		return
	}

	// Release the PC/SC context (when needed)
	defer context.Release()

	// List available readers
	readers, err := context.ListReaders()
	if err != nil {
		fmt.Println("Error ListReaders:", err)
		return
	}

	// Use the first reader
	reader := readers[0]
	fmt.Println("Using reader:", reader)

	// Connect to the card
	card, err := context.Connect(reader, scard.ShareShared, scard.ProtocolAny)
	if err != nil {
		fmt.Println("Error Connect:", err)
		return
	}

	// Disconnect (when needed)
	defer card.Disconnect(scard.LeaveCard)

	// Send select APDU

	// APDUReadICCID
	iccid, err := sscard.APDUGetRsp(card, sscard.APDUReadICCID)
	if err != nil {
		fmt.Println("Error Transmit:", err)
		return
	}
	fmt.Printf("resp APDUReadICCID: % 02x\r\n", iccid)
	iccidDecode := sscard.DecodeICCID(iccid)
	// fmt.Printf("ICCID Decode : % 02x Valid : %v\r\n\n", iccidDecode, sscard.LuhnByteValid(iccidDecode[:10], 20))
	fmt.Printf("ICCID Decode : % 02x\r\n\n", iccidDecode)

	// APDUReadICCIDclass00
	iccid00, err := sscard.APDUGetRsp(card, sscard.APDUReadICCIDclass00)
	if err != nil {
		fmt.Println("Error Transmit:", err)
		return
	}
	fmt.Printf("resp APDUReadICCIDclass00: % 02x\r\n", iccid00)
	iccid00Decode := sscard.DecodeICCID(iccid00)
	// fmt.Printf("ICCID Decode : % 02x Valid : %v\r\n\n", iccidDecode, sscard.LuhnByteValid(iccidDecode[:10], 20))
	fmt.Printf("ICCIDclass00 Decode : % 02x\r\n\n", iccid00Decode)

	imsi, err := sscard.APDUGetRsp(card, sscard.APDUReadIMSI)
	if err != nil {
		fmt.Println("Error Transmit:", err)
		return
	}
	fmt.Printf("resp APDUReadIMSI: % 02x\r\n", imsi)
	imsiDecode := sscard.DecodeIMSI(imsi)
	// fmt.Printf("IMSI Decode : % 02x Valid : %v\r\n\n", imsiDecode, sscard.LuhnByteValid(imsiDecode[:9], 18))
	fmt.Printf("IMSI Decode : % 02x\r\n\n", imsiDecode)

	imsi00, err := sscard.APDUGetRsp(card, sscard.APDUReadIMSIclass00)
	if err != nil {
		fmt.Println("Error Transmit:", err)
		return
	}
	fmt.Printf("resp APDUReadIMSIclass00: % 02x\r\n", imsi00)
	imsi00Decode := sscard.DecodeIMSI(imsi00)
	// fmt.Printf("IMSI Decode : % 02x Valid : %v\r\n\n", imsiDecode, sscard.LuhnByteValid(imsiDecode[:9], 18))
	fmt.Printf("IMSIclass00 Decode : % 02x\r\n\n", imsi00Decode)
	return
	// APDUReadSMS
	sms, err := sscard.APDUGetRsp(card, sscard.APDUReadSMS)
	if err != nil {
		fmt.Println("Error Transmit:", err)
		return
	}
	fmt.Printf("resp sscard.APDUReadSMS: % 02x\r\n\n", sms)

	// APDUReadLOCI
	data, err := sscard.APDUGetRsp(card, sscard.APDUReadLOCI)
	if err != nil {
		fmt.Println("Error Transmit:", err)
		return
	}
	fmt.Printf("resp sscard.APDUReadLOCI: % 02x\r\n\n", data)

}
