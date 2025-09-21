package models

import "testing"

func TestAddClient(t *testing.T) {
	server := &Server{}
	cli := server.AddClient("user1", "pass1")
	if cli == nil {
		t.Error("AddClient should return a new client")
	}
	if cli.id != "user1" {
		t.Errorf("Expected id 'user1', got '%v'", cli.id)
	}
	if cli.password != "pass1" {
		t.Errorf("Expected password 'pass1', got '%v'", cli.password)
	}
}

func TestVerifyPassword(t *testing.T) {
	server := &Server{}
	server.AddClient("user2", "pass2")
	if !server.VerifyPassword("pass2", "user2") {
		t.Error("VerifyPassword should return true for correct password")
	}
	if server.VerifyPassword("wrong", "user2") {
		t.Error("VerifyPassword should return false for wrong password")
	}
}
