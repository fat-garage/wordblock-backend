package dao

import (
	"github.com/fat-garage/wordblock-backend/models"
	"gorm.io/gorm"
)

type BlockDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) *BlockDao {
	return &BlockDao{
		db: db,
	}
}

// Query .
type Query struct {
	Limit   uint
	Offset  uint
	BlockID uint
	DID     string
	CID     string
}

// Where
func (c Query) where() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if c.BlockID > 0 {
			db = db.Where("id = ?", c.BlockID)
		}
		if len(c.DID) > 0 {
			db = db.Where("did = ?", c.DID)
		}
		if len(c.CID) > 0 {
			db = db.Where("address = ?", c.CID)
		}
		return db
	}
}

// GetDIDBlockList .
func (repo *BlockDao) GetDIDBlockList(query Query) ([]models.Block, error) {
	var results []models.Block
	if err := repo.db.Model(&models.Block{}).Scopes(query.where()).Find(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}

func (repo *BlockDao) AddBlock(did, fileName, cid string) error {
	entity := models.Block{
		Did:      did,
		Cid:      cid,
		FileName: fileName,
	}
	return repo.db.FirstOrCreate(&entity, entity).Error
}

func (repo *BlockDao) count(query Query) (int64, error) {
	var count int64
	if err := repo.db.Model(&models.Block{}).Scopes(query.where()).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
