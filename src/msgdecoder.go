//____________________________________________________________________
// File: msgdecoder.go
//____________________________________________________________________
//
// Author: Shaun Ashby <shaun@ashby.ch>
// Created: 2017-11-07 08:00:27+0100
// Revision: $Id$
// Description: Testbed for JSON message decoding
//
// Copyright (C) 2017 Shaun Ashby
//
//
//--------------------------------------------------------------------
package main

import (
	"encoding/json"
	"fmt"
)

var message = `{ "data": { "timestamp": "Tue  7 Nov 2017 08:01:44 CET", "clientid": "1234445" } }`

type MessageBody struct {
	Timestamp string `json: "timestamp"`
	ClientID  string `json: "clientid"`
}

func (m *MessageBody) Time() (string) {
	return m.Timestamp
}

func (m *MessageBody) Client() (string) {
	return m.ClientID
}

type Message struct {
	Data *MessageBody `json: "data"`
}

func (m *Message) Body() (*MessageBody) {
	return m.Data
}

func main() {
	fmt.Printf("%s\n", message)
	msg := Message{}
	if err := json.Unmarshal([]byte(message), &msg); err != nil {
		panic(err)
	}

	fmt.Printf("Decoded: TS -> \"%s\", Client Ident -> \"%s\"\n",msg.Body().Time(), msg.Body().Client())

}
