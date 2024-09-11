package models

import "time"

// TransactionReceipt 交易收据数据表，存储每个交易的收据信息
type TransactionReceipt struct {
	ID                int64     `gorm:"primaryKey;autoIncrement" json:"id" comment:"自增主键，交易收据唯一标识"`
	TxHash            string    `gorm:"size:66;not null;unique" json:"tx_hash" comment:"交易哈希值，66个字符，开头为0x"`
	BlockHash         string    `gorm:"size:66;index:idx_block_hash" json:"block_hash" comment:"所属区块的哈希值"`
	BlockNumber       int64     `json:"block_number" comment:"所属区块的区块号"`
	CumulativeGasUsed int64     `json:"cumulative_gas_used" comment:"区块中从起始位置到此交易消耗的累计Gas量"`
	GasUsed           int64     `json:"gas_used" comment:"此交易实际消耗的Gas量"`
	ContractAddress   string    `gorm:"size:42" json:"contract_address" comment:"如果是合约部署交易，记录合约地址"`
	Status            uint8     `json:"status" comment:"交易执行状态，1表示成功，0表示失败"`
	Logs              string    `gorm:"type:text" json:"logs" comment:"交易的事件日志"`
	CreatedAt         time.Time `gorm:"autoCreateTime" json:"created_at" comment:"数据插入时间"`
	UpdatedAt         time.Time `gorm:"autoUpdateTime" json:"updated_at" comment:"数据更新时间"`
}
