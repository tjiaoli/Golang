package models

import (
	"math/big"
	"time"
)

type Transaction struct {
	ID          int64     `gorm:"primaryKey;autoIncrement" json:"id" comment:"自增主键，交易唯一标识"`
	TxHash      string    `gorm:"size:66;not null;unique" json:"tx_hash" comment:"交易哈希值，66个字符，开头为0x"`
	BlockHash   string    `gorm:"size:66;index:idx_block_hash" json:"block_hash" comment:"所属区块的哈希值"`
	BlockNumber int64     `json:"block_number" comment:"所属区块的区块号"`
	FromAddress string    `gorm:"size:42" json:"from_address" comment:"发送方地址，42个字符，开头为0x"`
	ToAddress   string    `gorm:"size:42" json:"to_address" comment:"接收方地址，42个字符，开头为0x"`
	Value       *big.Int  `gorm:"type:decimal(38,0)" json:"value" comment:"交易发送的以太坊数量"`
	GasPrice    int64     `json:"gas_price" comment:"交易的Gas价格"`
	GasUsed     int64     `json:"gas_used" comment:"交易消耗的Gas量"`
	Nonce       int64     `json:"nonce" comment:"发送方账户的交易序号"`
	InputData   string    `gorm:"type:text" json:"input_data" comment:"交易的输入数据（ABI调用数据或合约代码）"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at" comment:"数据插入时间"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at" comment:"数据更新时间"`
}
