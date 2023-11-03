package orm

import (
	"context"
	"fmt"
	"time"

	"github.com/scroll-tech/go-ethereum/log"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"

	"github.com/scroll-tech/chain-monitor/internal/types"
)

// TransactionMatch contains the tx of l1 & l2
type TransactionMatch struct {
	db *gorm.DB `gorm:"column:-"`

	ID          int64  `json:"id" gorm:"column:id"`
	MessageHash string `json:"messageHash" gorm:"message_hash"`
	TokenType   int    `json:"token_type" gorm:"token_type"`

	// tx info of l1
	L1EventType   int             `json:"l1_event_type" gorm:"l1_event_type"`
	L1BlockNumber uint64          `json:"l1_block_number" gorm:"l1_block_number"`
	L1TxHash      string          `json:"l1_tx_hash" gorm:"l1_tx_hash"`
	L1TokenId     string          `json:"l1_token_id" gorm:"l1_token_id"`
	L1Value       decimal.Decimal `json:"l1_value" gorm:"l1_value"`

	// tx info of l2
	L2EventType   int             `json:"l2_event_type" gorm:"l2_event_type"`
	L2BlockNumber uint64          `json:"l2_block_number" gorm:"l2_block_number"`
	L2TxHash      string          `json:"l2_tx_hash" gorm:"l2_tx_hash"`
	L2TokenId     string          `json:"l2_token_id" gorm:"l2_token_id"`
	L2Value       decimal.Decimal `json:"l2_value" gorm:"l2_value"`

	// status
	CheckStatus        int    `json:"check_status" gorm:"check_status"`
	WithdrawRootStatus int    `json:"withdraw_root_status" gorm:"withdraw_root_status"`
	L1GatewayStatus    int    `json:"l1_gateway_status" gorm:"l1_gateway_status"`
	L2GatewayStatus    int    `json:"l2_gateway_status" gorm:"l2_gateway_status"`
	L1CrossChainStatus int    `json:"l1_cross_chain_status" gorm:"l1_cross_chain_status"`
	L2CrossChainStatus int    `json:"l2_cross_chain_status" gorm:"l2_cross_chain_status"`
	MessageProof       string `json:"msg_proof" gorm:"message_proof"`

	CreatedAt time.Time      `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"column:deleted_at"`
}

// NewTransactionMatch creates a new TransactionMatch database instance.
func NewTransactionMatch(db *gorm.DB) *TransactionMatch {
	return &TransactionMatch{db: db}
}

// TableName returns the table name for the Batch model.
func (*TransactionMatch) TableName() string {
	return "transaction_match"
}

// GetLatestTransactionMatch get the latest uncheck transaction match record
func (t *TransactionMatch) GetLatestTransactionMatch(ctx context.Context, limit int) ([]TransactionMatch, error) {
	var transactions []TransactionMatch
	db := t.db.WithContext(ctx)
	db = db.Where("check_status = ?", types.CheckStatusUnchecked)
	db = db.Order("id asc")
	db = db.Limit(limit)
	if err := db.Find(&transactions).Error; err != nil {
		log.Warn("TransactionMatch.GetLatestTransactionMatch failed", "error", err)
		return nil, fmt.Errorf("TransactionMatch.GetLatestTransactionMatch failed err:%w", err)
	}
	return transactions, nil
}

func (t *TransactionMatch) InsertOrUpdate(ctx context.Context, transactions []TransactionMatch) (int, error) {
	// insert or update
}

func (t *TransactionMatch) UpdateGatewayStatus(ctx context.Context, id []int64, layerType types.LayerType, status types.GatewayStatusType) error {
	db := t.db.WithContext(ctx)
	db = db.Model(&TransactionMatch{})
	db = db.Where("id = (?)", id)

	var err error
	switch layerType {
	case types.Layer1:
		err = db.Update("l1_gateway_status", status).Error
	case types.Layer2:
		err = db.Update("l2_gateway_status", status).Error
	}

	if err != nil {
		log.Warn("TransactionMatch.UpdateGatewayStatus failed", "error", err)
		return fmt.Errorf("TransactionMatch.UpdateGatewayStatus failed err:%w", err)
	}
	return nil
}

func (t *TransactionMatch) UpdateCrossChainStatus(ctx context.Context, id []int64, layerType types.LayerType, status types.CrossChainStatusType) error {
	db := t.db.WithContext(ctx)
	db = db.Model(&TransactionMatch{})
	db = db.Where("id in (?)", id)

	var err error
	switch layerType {
	case types.Layer1:
		err = db.Update("l1_cross_chain_status", status).Error
	case types.Layer2:
		err = db.Update("l2_cross_chain_status", status).Error
	}

	if err != nil {
		log.Warn("TransactionMatch.UpdateCrossChainStatus failed", "error", err)
		return fmt.Errorf("TransactionMatch.UpdateCrossChainStatus failed err:%w", err)
	}
	return nil
}
