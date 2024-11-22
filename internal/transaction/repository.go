package transaction

import (
	"gorm.io/gorm"
)

type TransactionRepository interface {
	CreateTransaction(transaction *Transaction) error
	GetTransactionByID(transactionID int) (*Transaction, error)
	GetAllTransactions() ([]Transaction, error) 
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{db: db}
}

func (r *transactionRepository) CreateTransaction(transaction *Transaction) error {
	if err := r.db.Create(transaction).Error; err != nil {
		return err
	}
	return nil
}

// GetAllTransactions retrieves all transactions with their associated items
func (r *transactionRepository) GetAllTransactions() ([]Transaction, error) {
	var transactions []Transaction
	err := r.db.Preload("Items").Find(&transactions).Error
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

// GetTransactionByID retrieves a transaction by its ID
func (r *transactionRepository) GetTransactionByID(transactionID int) (*Transaction, error) {
	var transaction Transaction
	err := r.db.Preload("Items").First(&transaction, transactionID).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil 
		}
		return nil, err 
	}
	return &transaction, nil
}