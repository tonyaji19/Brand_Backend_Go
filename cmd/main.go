package main

import (
	"encoding/json"
	"fmt"
	"golang-api-service/config"
	database "golang-api-service/database/migration"
	"golang-api-service/internal/brand"
	"golang-api-service/internal/transaction"
	"golang-api-service/internal/voucher"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// CreateRedemptionHandler handles the creation of a redemption transaction
func CreateRedemptionHandler(repo transaction.TransactionRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var redemptionRequest transaction.RedemptionRequest
		if err := json.NewDecoder(r.Body).Decode(&redemptionRequest); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		tx := transaction.Transaction{
			CustomerName: redemptionRequest.CustomerName,
			TotalPoints:  0, 
		}

		var items []transaction.TransactionItem
		totalPoints := 0
		for _, item := range redemptionRequest.VoucherItems {
			totalPoints += item.Quantity * item.Points
			items = append(items, transaction.TransactionItem{
				VoucherID:   item.VoucherID,
				Quantity:    item.Quantity,
				TotalPoints: item.Quantity * item.Points,
			})
		}

		tx.Items = items
		tx.TotalPoints = totalPoints

		if err := repo.CreateTransaction(&tx); err != nil {
			http.Error(w, "Failed to create redemption transaction", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(tx)
	}
}
// GetRedemptionByTransactionIDHandler handles fetching a redemption transaction by transaction ID
func GetRedemptionByTransactionIDHandler(repo transaction.TransactionRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		transactionIDStr := vars["transactionId"]

		transactionID, err := strconv.Atoi(transactionIDStr)
		if err != nil {
			http.Error(w, "Invalid transactionId", http.StatusBadRequest)
			return
		}

		transaction, err := repo.GetTransactionByID(transactionID)
		if err != nil {
			http.Error(w, "Transaction not found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(transaction)
	}
}

func GetVouchersByBrandHandler(repo voucher.VoucherRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		brandIDStr := r.URL.Query().Get("id")
		if brandIDStr == "" {
			http.Error(w, "brand_id is required", http.StatusBadRequest)
			return
		}

		brandID, err := strconv.Atoi(brandIDStr)
		if err != nil {
			http.Error(w, "Invalid brand_id", http.StatusBadRequest)
			return
		}

		vouchers, err := repo.GetVouchersByBrandID(brandID)
		if err != nil {
			http.Error(w, "Failed to fetch vouchers", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(vouchers)
	}
}

// GetAllTransactionsHandler handles fetching all transactions
func GetAllTransactionsHandler(repo transaction.TransactionRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		transactions, err := repo.GetAllTransactions()
		if err != nil {
			http.Error(w, "Failed to fetch transactions", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(transactions)
	}
}

// GetTransactionByIDHandler handles fetching a single transaction by ID
func GetTransactionByIDHandler(repo transaction.TransactionRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, "Invalid transaction ID", http.StatusBadRequest)
			return
		}

		transaction, err := repo.GetTransactionByID(id)
		if err != nil {
			http.Error(w, "Transaction not found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(transaction)
	}
}

// CreateTransactionHandler handles creating a new transaction
func CreateTransactionHandler(repo transaction.TransactionRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newTransaction transaction.Transaction
		if err := json.NewDecoder(r.Body).Decode(&newTransaction); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		if err := repo.CreateTransaction(&newTransaction); err != nil {
			http.Error(w, "Failed to create transaction", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newTransaction)
	}
}


// Voucher Handlers
func GetAllVouchersHandler(repo voucher.VoucherRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vouchers, err := repo.GetAllVouchers()
		if err != nil {
			http.Error(w, "Failed to fetch vouchers", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(vouchers)
	}
}

func GetVoucherByIDHandler(repo voucher.VoucherRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, "Invalid voucher ID", http.StatusBadRequest)
			return
		}

		voucher, err := repo.GetVoucherByID(id)
		if err != nil {
			http.Error(w, "Voucher not found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(voucher)
	}
}

func CreateVoucherHandler(repo voucher.VoucherRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newVoucher voucher.Voucher
		if err := json.NewDecoder(r.Body).Decode(&newVoucher); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		if err := repo.CreateVoucher(&newVoucher); err != nil {
			http.Error(w, "Failed to create voucher", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newVoucher)
	}
}


//brand
func GetAllBrandsHandler(repo brand.BrandRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		brands, err := repo.GetAllBrands()
		if err != nil {
			http.Error(w, "Failed to fetch brands", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(brands)
	}
}

func GetBrandByIDHandler(repo brand.BrandRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, "Invalid brand ID", http.StatusBadRequest)
			return
		}

		brand, err := repo.GetBrandByID(id)
		if err != nil {
			http.Error(w, "Brand not found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(brand)
	}
}

func CreateBrandHandler(repo brand.BrandRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newBrand brand.Brand
		if err := json.NewDecoder(r.Body).Decode(&newBrand); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		if err := repo.CreateBrand(&newBrand); err != nil {
			http.Error(w, "Failed to create brand", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newBrand)
	}
}

func main() {
	db, err := config.GetDBConnection()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	database.RunMigrations()

	brandRepo := brand.NewBrandRepository(db)
	voucherRepo := voucher.NewVoucherRepository(db)
	transactionRepo := transaction.NewTransactionRepository(db)

	router := mux.NewRouter()

	router.HandleFunc("/brands", GetAllBrandsHandler(brandRepo)).Methods("GET")
	router.HandleFunc("/brands/{id:[0-9]+}", GetBrandByIDHandler(brandRepo)).Methods("GET")
	router.HandleFunc("/brands", CreateBrandHandler(brandRepo)).Methods("POST")

	router.HandleFunc("/vouchers", GetAllVouchersHandler(voucherRepo)).Methods("GET")
	router.HandleFunc("/vouchers/{id:[0-9]+}", GetVoucherByIDHandler(voucherRepo)).Methods("GET")
	router.HandleFunc("/vouchers", CreateVoucherHandler(voucherRepo)).Methods("POST")
	router.HandleFunc("/voucher/brand", GetVouchersByBrandHandler(voucherRepo)).Methods("GET") 

	router.HandleFunc("/transactions", GetAllTransactionsHandler(transactionRepo)).Methods("GET")
	router.HandleFunc("/transactions/{id:[0-9]+}", GetTransactionByIDHandler(transactionRepo)).Methods("GET")
	router.HandleFunc("/transactions", CreateTransactionHandler(transactionRepo)).Methods("POST")
	router.HandleFunc("/transactions/redemption", CreateRedemptionHandler(transactionRepo)).Methods("POST")
	router.HandleFunc("/transactions/redemption/{transactionId:[0-9]+}", GetRedemptionByTransactionIDHandler(transactionRepo)).Methods("GET")
	

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))

	queryAndPrintVouchers(db)
}

func queryAndPrintVouchers(db *gorm.DB) {
	var vouchers []voucher.Voucher
	if err := db.Preload("Brand").Find(&vouchers).Error; err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, v := range vouchers {
		fmt.Printf("Voucher: %s, Brand: %s\n", v.Code, v.Brand.Name)
	}
}
