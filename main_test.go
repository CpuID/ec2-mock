package main

import (
	"net/http/httptest"
	"os"
	"testing"
)

// Not a fan of globals, but it's the only sane way to pass an httptest instance into each of the tests...
var (
	testServer *httptest.Server
)

func TestMain(m *testing.M) {
	// Setup the test API server
	server := &Server{}
	// Mock parameters
	server.Tags = InstanceTags{
		"i-asdfasdf": []InstanceTag{
			InstanceTag{
				Key:   "BLAH",
				Value: "asdf",
			},
			InstanceTag{
				Key:   "aaaa",
				Value: "bbbb",
			},
		},
		"i-aaaabbbb": []InstanceTag{
			InstanceTag{
				Key:   "aaaa",
				Value: "cccc",
			},
			InstanceTag{
				Key:   "ZZZZ",
				Value: "yyyy",
			},
		},
	}
	testServer = httptest.NewServer(server.NewMux())
	defer testServer.Close()

	// Run the tests
	os.Exit(m.Run())
}
