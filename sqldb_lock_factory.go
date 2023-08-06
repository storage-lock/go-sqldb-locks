package gorp_locks

import (
	"database/sql"
	sqldb_storage "github.com/storage-lock/go-sqldb-storage"
	"github.com/storage-lock/go-storage"
	storage_lock_factory "github.com/storage-lock/go-storage-lock-factory"
)

type SqlDbLockFactory struct {
	db *sql.DB
	*storage_lock_factory.StorageLockFactory[*sql.DB]
}

func NewSqlDbLockFactory(db *sql.DB) (*SqlDbLockFactory, error) {
	connectionManager := storage.NewFixedSqlDBConnectionManager(db)

	storage, err := sqldb_storage.NewStorageBySqlDb(db, connectionManager)
	if err != nil {
		return nil, err
	}

	factory := storage_lock_factory.NewStorageLockFactory[*sql.DB](storage, connectionManager)

	return &SqlDbLockFactory{
		db:                 db,
		StorageLockFactory: factory,
	}, nil
}
