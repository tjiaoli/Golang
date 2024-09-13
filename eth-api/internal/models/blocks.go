package models

import "time"

// Block 区块数据表，存储区块基本信息
type Block struct {
	ID               int64     `gorm:"primaryKey;autoIncrement" json:"id" comment:"自增主键，区块唯一标识"`
	BlockHash        string    `gorm:"size:66;not null;unique" json:"block_hash" comment:"区块哈希值，66个字符，开头为0x"`
	BlockNumber      int64     `gorm:"not null;index:idx_block_number" json:"block_number" comment:"区块号"`
	ParentHash       string    `gorm:"size:66" json:"parent_hash" comment:"父区块哈希值，66个字符，开头为0x"`
	Timestamp        int64     `gorm:"not null" json:"timestamp" comment:"区块的时间戳（Unix 时间格式）"`
	Miner            string    `gorm:"size:42" json:"miner" comment:"挖矿者地址，42个字符，开头为0x"`
	GasUsed          int64     `json:"gas_used" comment:"区块中消耗的总Gas"`
	GasLimit         int64     `json:"gas_limit" comment:"区块的Gas上限"`
	Size             int64     `json:"size" comment:"区块的大小"`
	TransactionCount int       `json:"transaction_count" comment:"区块中的交易数量"`
	CreatedAt        time.Time `gorm:"autoCreateTime" json:"created_at" comment:"数据插入时间"`
	UpdatedAt        time.Time `gorm:"autoUpdateTime" json:"updated_at" comment:"数据更新时间"`
}
