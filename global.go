package gorp_locks

import (
	"database/sql"
	storage_lock "github.com/storage-lock/go-storage-lock"
	"sync"
)

var GlobalSqlDbLockFactory *SqlDbLockFactory
var globalSqlDbLockFactoryOnce sync.Once
var globalSqlDbLockFactoryErr error

func InitGlobalSqlDbLockFactory(db *sql.DB) error {
	factory, err := NewSqlDbLockFactory(db)
	if err != nil {
		return err
	}
	GlobalSqlDbLockFactory = factory
	return nil
}

func NewGorpLock(db *sql.DB, lockId string) (*storage_lock.StorageLock, error) {
	globalSqlDbLockFactoryOnce.Do(func() {
		globalSqlDbLockFactoryErr = InitGlobalSqlDbLockFactory(db)
	})
	if globalSqlDbLockFactoryErr != nil {
		return nil, globalSqlDbLockFactoryErr
	}
	return GlobalSqlDbLockFactory.CreateLock(lockId)
}
