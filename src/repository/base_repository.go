package repository

import "gorm.io/gorm"

type base struct {
	db *gorm.DB
}

func (b *base) StartTransaction() *gorm.DB {
	tx := b.db.Begin()
	return tx
}
func (b *base) CommitTransaction(tx *gorm.DB) error {
	return tx.Commit().Error
}
func (b *base) RollBackTransaction(tx *gorm.DB) error {
	return tx.Rollback().Error
}
