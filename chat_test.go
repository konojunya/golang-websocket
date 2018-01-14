package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestValidCase(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(createRouter().Handler()))
	defer ts.Close()

	client1, err := createClient(ts)
	if err != nil {
		t.Fatal(err)
	}
	defer client1.Close()

	client2, err := createClient(ts)
	if err != nil {
		t.Fatal(err)
	}
	defer client2.Close()

	err = writeMessage(client1, `join {"name":"room1}`)
	if err != nil {
		t.Fatal(err)
	}

	err = writeMessage(client2, `join {"name":"room1"}`)
	if err != nil {
		t.Fatal(err)
	}

	err = writeMessage(client1, `say {"name":"room1", "msg":"hello room1"}`)
	if err != nil {
		t.Fatal(err)
	}

	res, err := readMessage(client1)
	if err != nil {
		t.Error(err)
	}
	if res != `say {"msg":"hello room1"}` {
		t.Error("response is not valid: " + res)
	}

	res, err = readMessage(client2)
	if err != nil {
		t.Error(err)
	}
	if res != `say {"msg":"hello room1"}` {
		t.Error("response is not valid: " + res)
	}
}
