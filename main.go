package main

import (
	"fmt"
	"time"
)

type client struct {
	id int
}

type clientWithDoc struct {
	c client
	d document
}

func (c client) TransmitToDigitalSafetyDepositBox(d *digitalSafetyDepositBox, doc document) {
	d.RecordDocument(doc)
}

type document struct {
	content string
}

type stampedDocument struct {
	d document
	t time.Time
}

type digitalSafetyDepositBox struct {
	documentsRecorded []stampedDocument
}

func (d *digitalSafetyDepositBox) RecordDocument(doc document) {
	stampedDoc := stampedDocument{doc, time.Now()}
	d.documentsRecorded = append(d.documentsRecorded, stampedDoc)
}

func main() {
	_digSafDepoBox := digitalSafetyDepositBox{}
	clients := []client{}
	clientsWithDocs := []clientWithDoc{}
	for i := 0; i < 10; i++ {
		clients = append(clients, client{i})
	}

	for i := 0; i < 10; i++ {
		newDoc := document{fmt.Sprint(i)}
		_clientWithDoc := clientWithDoc{clients[i], newDoc}
		clientsWithDocs = append(clientsWithDocs, _clientWithDoc)
	}

	for _, c := range clientsWithDocs {
		c.c.TransmitToDigitalSafetyDepositBox(&_digSafDepoBox, c.d)
	}

	for _, d := range _digSafDepoBox.documentsRecorded {
		fmt.Println(d.d)
		fmt.Println(d.t)
	}

}
