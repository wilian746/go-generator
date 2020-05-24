package database

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetConnection(t *testing.T) {
	t.Run("should return panic when dialect is empty", func(t *testing.T) {
		assert.Panics(t, func() { GetConnection("", "") }, "The code contains panic")
	})
	t.Run("should return panic when dialect is invalid", func(t *testing.T) {
		assert.Panics(t, func() { GetConnection("generic", "") }, "The code contains panic")
	})
	t.Run("should return panic when dialect is not equal", func(t *testing.T) {
		assert.Panics(t, func() { GetConnection("POSTGRES", "") }, "The code contains panic")
	})
	t.Run("should return panic when uri is empty", func(t *testing.T) {
		assert.Panics(t, func() { GetConnection("postgres", "") }, "The code contains panic")
	})
	t.Run("should return panic when connect in database return error", func(t *testing.T) {
		assert.Panics(t, func() { GetConnection("postgres", "generic") }, "The code contains panic")
	})
	t.Run("should not return panic when connect in database", func(t *testing.T) {
		assert.NotPanics(t, func() { GetConnection("sqlite3", ":memory:") }, "The code contains panic")
	})
}
