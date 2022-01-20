package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDSNGenerator(t *testing.T) {
	_, err := GenerateDSN(DBOptions{Driver: "fake"})
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if err.Error() != "unsupported driver: fake" {
		t.Fatal("expected error message, got:", err.Error())
	}
}

func TestDSNGeneratorMySQL(t *testing.T) {
	mysql := DBOptions{
		Driver:   "mysql",
		Host:     "localhost",
		Port:     "3306",
		User:     "root",
		Password: "secret12345",
		Database: "test",
		Options:  "charset=utf8&parseTime=True",
	}

	expected := "root:secret12345@tcp(localhost:3306)/test?charset=utf8&parseTime=True"
	dsn, err := GenerateDSN(mysql)
	if err != nil {
		t.Fatal("expected no error, got:", err)
	}
	if dsn != expected {
		t.Fatal("expected dsn to be:", expected, "got:", dsn)
	}
}

func TestDSNGeneratorPostgres(t *testing.T) {
	postgres := DBOptions{
		Driver:   "postgres",
		Host:     "localhost",
		Port:     "5432",
		User:     "postgres",
		Password: "secret12345",
		Database: "test",
		Options:  "sslmode=disable",
	}

	expected := "host=localhost port=5432 user=postgres password=secret12345 dbname=test sslmode=disable"
	dsn, err := GenerateDSN(postgres)
	if err != nil {
		t.Fatal("expected no error, got:", err)
	}
	assert.Equal(t, expected, dsn)
}

func TestIsValidDriver(t *testing.T) {
	assert.False(t, isValidDriver("fake"))
	assert.True(t, isValidDriver("mysql"))
	assert.True(t, isValidDriver("postgres"))
}
