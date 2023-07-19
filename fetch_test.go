package main

import (
	"fmt"
	"testing"
)

func TestFetchFeed(t *testing.T) {
	fmt.Println("Test Case 1 - Valid Fetch URL")
	_, err := FetchFeed("https://blog.boot.dev/index.html")
	if err == nil {
		t.Errorf("Expected error but found nil")
	}
	fmt.Println("Test Case 2 - Invalid Fetch URL")
	_, err = FetchFeed("https://blog.boot.dev/index.xml")
	if err != nil {
		t.Errorf("Expected nil error but found error")
	}
}
