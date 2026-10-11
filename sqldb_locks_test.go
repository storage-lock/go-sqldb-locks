package sqldb_locks

import (
	"context"
	"database/sql"
	"testing"

	"github.com/storage-lock/go-sqldb-storage"
	"github.com/storage-lock/go-storage-lock"
	"github.com/stretchr/testify/assert"
)

func TestNewLockBySqlDb_NilDB(t *testing.T) {
	_, err := NewLockBySqlDb(context.Background(), nil, "lock-1")
	assert.Error(t, err)
}

func TestNewLockBySqlDb_Success(t *testing.T) {
	db, err := sql.Open("sqlite3", "file:test_success.db?mode=memory")
	assert.NoError(t, err)
	defer db.Close()

	lock, err := NewLockBySqlDb(context.Background(), db, "test-lock-success")
	assert.NoError(t, err)
	assert.NotNil(t, lock)
}

func TestNewLockBySqlDbWithOptions_NilDB(t *testing.T) {
	options := storage_lock.NewStorageLockOptions()
	options.SetLockId("test-lock")
	_, err := NewLockBySqlDbWithOptions(context.Background(), nil, options)
	assert.Error(t, err)
}

func TestNewLockBySqlDbWithOptions(t *testing.T) {
	db, err := sql.Open("sqlite3", "file:test.db?mode=memory")
	assert.NoError(t, err)
	defer db.Close()

	options := storage_lock.NewStorageLockOptions()
	options.SetLockId("test-lock")
	lock, err := NewLockBySqlDbWithOptions(context.Background(), db, options)
	assert.NoError(t, err)
	assert.NotNil(t, lock)
}

func TestGetLockFactoryBySqlDb(t *testing.T) {
	db, err := sql.Open("sqlite3", "file:test.db?mode=memory")
	assert.NoError(t, err)
	defer db.Close()

	factory, err := GetLockFactoryBySqlDb(context.Background(), db)
	assert.NoError(t, err)
	assert.NotNil(t, factory)
}

func TestNewStorageBySqlDb_UnsupportedDriver(t *testing.T) {
	// Test the underlying storage function directly for unsupported driver
	_, err := sqldb_storage.NewStorageBySqlDb(nil, nil)
	assert.Error(t, err)
}

