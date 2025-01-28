package connectdb_test

import (
	connectdb "main/data-access"
	"testing"
)

func TestConnectDB(t *testing.T) {
	db, err := connectdb.ConnectDB()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if db == nil {
		t.Fatalf("Expected a database connection, got nil")
	}
}