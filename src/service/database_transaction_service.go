package service

import "gorm.io/gorm"

type IDatabaseTransaction interface {
	StartTransaction() *gorm.DB
	CommitTransaction(tx *gorm.DB) error
	RollBackTransaction(tx *gorm.DB) error
}

type DatabaseTransaction struct {
	db *gorm.DB
}

func NewDatabaseTransaction(db *gorm.DB) IDatabaseTransaction {
	return &DatabaseTransaction{db: db}
}
func (b *DatabaseTransaction) StartTransaction() *gorm.DB {
	tx := b.db.Begin()
	return tx
}
func (b *DatabaseTransaction) CommitTransaction(tx *gorm.DB) error {
	return tx.Commit().Error
}
func (b *DatabaseTransaction) RollBackTransaction(tx *gorm.DB) error {
	return tx.Rollback().Error
}
