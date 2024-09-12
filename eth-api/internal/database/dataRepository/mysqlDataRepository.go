package dataRepository

import (
	"eth-api/internal/database"
	"eth-api/internal/models"
	"gorm.io/gorm"
)

type BlocksRepository struct {
	db *gorm.DB
}

func NewBlocksRepository() *BlocksRepository {
	return &BlocksRepository{db: database.DB}
}

// 获取区块数据（从MySQL）
func (repo *BlocksRepository) GetBlockFromMySQL(blockNum int) *models.Block {
	var blockData models.Block
	result := repo.db.Table("blocks").Where("block_number = ?", blockNum).First(&blockData)
	if result.Error != nil {
		return nil
	}
	return &blockData
}
