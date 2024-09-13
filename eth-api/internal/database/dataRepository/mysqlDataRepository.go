package dataRepository

import (
	"errors"
	"eth-api/internal/database"
	"eth-api/internal/models"
	"fmt"
	"gorm.io/gorm"
	"log"
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

// 保存区块数据到MySQL
func (repo *BlocksRepository) SaveBlockToMySQL(blockData *models.Block) error {
	result := repo.db.Table("blocks").Save(blockData)
	if result.Error != nil {
		log.Printf("Failed to save block to MySQL: %v", result.Error)
		return result.Error
	}
	return nil
}

// 保存交易数据到MySQL
func (repo *BlocksRepository) SaveTxToMySQL(txDataList []*models.Transaction) error {
	for _, tx := range txDataList {
		txHash := tx.TxHash
		//查对应hash是都已经存储
		_, err := repo.GetTxFromMySQL(txHash)
		if err != nil {
			result := repo.db.Table("transactions").Save(tx)
			if result.Error != nil {
				return fmt.Errorf("failed to save transaction %s: %v", tx.TxHash, result.Error)
			}
		}
	}
	return nil
}

// 获取交易数据（从MySQL）
func (repo *BlocksRepository) GetTxFromMySQL(txHash string) (*models.Transaction, error) {
	var txData models.Transaction
	result := repo.db.Table("transactions").Where("tx_hash = ?", txHash).First(&txData)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			log.Printf("Transaction not found for hash: %s", txHash)
			return nil, fmt.Errorf("Transaction not found for hash: %s", txHash)
		} else {
			log.Fatalf("Failed to query transaction: %v", result.Error)
		}
	}

	if result.RowsAffected == 0 {
		log.Printf("No transaction found for hash: %s", txHash)
		return nil, fmt.Errorf("No transaction found for hash: %s", txHash)
	}

	return &txData, nil
}

// 获取交易存根数据（从MySQL）
func (repo *BlocksRepository) GetReceiptFromMySQL(txHash string) *models.TransactionReceipt {
	var receiptData models.TransactionReceipt
	result := repo.db.Table("receipts").Where("tx_hash = ?", txHash).First(&receiptData)
	if result.Error != nil {
		return nil
	}
	return &receiptData
}

// 保存交易存根数据到MySQL
func (repo *BlocksRepository) SaveReceiptToMySQL(receiptData *models.TransactionReceipt) error {
	result := repo.db.Table("transaction_receipts").Create(receiptData)
	if result.Error != nil {
		log.Printf("Failed to save receipt to MySQL: %v", result.Error)
		return result.Error
	}
	return nil
}
