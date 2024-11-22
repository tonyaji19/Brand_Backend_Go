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



func GetVouchersByBrandHandler(repo voucher.VoucherRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Ambil parameter query id
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

		// Ambil data voucher berdasarkan brand_id
		vouchers, err := repo.GetVouchersByBrandID(brandID)
		if err != nil {
			http.Error(w, "Failed to fetch vouchers", http.StatusInternalServerError)
			return
		}

		// Kirim response
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
	// Database Connection
	db, err := config.GetDBConnection()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Run Database Migrations
	database.RunMigrations()

	// Repositories
	brandRepo := brand.NewBrandRepository(db)
	voucherRepo := voucher.NewVoucherRepository(db)
	transactionRepo := transaction.NewTransactionRepository(db)

	// Router
	router := mux.NewRouter()

	// Define routes
	router.HandleFunc("/brands", GetAllBrandsHandler(brandRepo)).Methods("GET")
	router.HandleFunc("/brands/{id:[0-9]+}", GetBrandByIDHandler(brandRepo)).Methods("GET")
	router.HandleFunc("/brands", CreateBrandHandler(brandRepo)).Methods("POST")

	// Voucher Routes
	router.HandleFunc("/vouchers", GetAllVouchersHandler(voucherRepo)).Methods("GET")
	router.HandleFunc("/vouchers/{id:[0-9]+}", GetVoucherByIDHandler(voucherRepo)).Methods("GET")
	router.HandleFunc("/vouchers", CreateVoucherHandler(voucherRepo)).Methods("POST")
	router.HandleFunc("/voucher/brand", GetVouchersByBrandHandler(voucherRepo)).Methods("GET") // Tambahkan route ini



	router.HandleFunc("/transactions", GetAllTransactionsHandler(transactionRepo)).Methods("GET")
	router.HandleFunc("/transactions/{id:[0-9]+}", GetTransactionByIDHandler(transactionRepo)).Methods("GET")
	router.HandleFunc("/transactions", CreateTransactionHandler(transactionRepo)).Methods("POST")
	//router.HandleFunc("/transaction/redemption", transaction.CreateRedemptionHandler(transactionRepo)).Methods("POST")


	// Start Server
	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))

	// Query vouchers and print
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
