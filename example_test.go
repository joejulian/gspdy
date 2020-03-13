// Copyright 2015 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gspdy_test

import (
	"log"
	"net/http"
	"testing"

	"github.com/joejulian/gspdy"
)

var (
	c   *gspdy.Conn
	req *http.Request
)

// The gspdy.IsUnexpectedCloseError function is useful for identifying
// application and protocol errors.
//
// This server application works with a client application running in the
// browser. The client application does not explicitly close the SPDY socket. The
// only expected close message from the client has the code
// gspdy.CloseGoingAway. All other other close messages are likely the
// result of an application or protocol error and are logged to aid debugging.
func ExampleIsUnexpectedCloseError() {

	for {
		messageType, p, err := c.ReadMessage()
		if err != nil {
			if gspdy.IsUnexpectedCloseError(err, gspdy.CloseGoingAway) {
				log.Printf("error: %v, user-agent: %v", err, req.Header.Get("User-Agent"))
			}
			return
		}
		processMesage(messageType, p)
	}
}

func processMesage(mt int, p []byte) {}

// TestX prevents godoc from showing this entire file in the example. Remove
// this function when a second example is added.
func TestX(t *testing.T) {}
