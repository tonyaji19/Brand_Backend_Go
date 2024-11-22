# Brand_Backend
```json
GO + Sql Server to - Create Brand POST /brand Create brand endpoint - Create Voucher POST /voucher Create voucher endpoint - Get Single Voucher GET /voucher?id={voucher_id} - Get All Vouchers by Brand GET /voucher/brand?id={brand_id}
```
# How to run

## 1. Create Database

## 2. Migrations
```json
goose -dir ./database/migration {sqlserver} "{sqlserver}://{USERDB}:{PASSWORDDB}@localhost:{PORT}?database=VoucherAPI" up
```

## 3. Run App
```json
go run cmd/main.go
```


# API Endpoints Documentation

## 1. Brands

### GET /brands

- **Response**:

```json
[
  { "ID": 1, "Name": "Brand A" },
  { "ID": 2, "Name": "Brand B" }
]
```

### GET /brands/{id}

- **Response**:

```json
{
  "ID": 1,
  "Name": "Brand A"
}
```

### POST /brands

- **Request Body**:

```json
{
  "Name": "Brand C"
}
```

- **Response**:

```json
{
  "ID": 3,
  "Name": "Brand C"
}
```

---

## 2. Vouchers

### GET /vouchers

- **Response**:

```json
[
  { "ID": 1, "Code": "VOUCHER1", "BrandID": 1 },
  { "ID": 2, "Code": "VOUCHER2", "BrandID": 2 }
]
```

### GET /vouchers/{id}

- **Response**:

```json
{
  "ID": 1,
  "Code": "VOUCHER1",
  "BrandID": 1
}
```

### POST /vouchers

- **Request Body**:

```json
{
  "Code": "VOUCHER3",
  "BrandID": 3
}
```

- **Response**:

```json
{
  "ID": 3,
  "Code": "VOUCHER3",
  "BrandID": 3
}
```

### GET /voucher/brand

- **Response**:

```json
[
  { "ID": 1, "Code": "VOUCHER1", "BrandID": 1 },
  { "ID": 2, "Code": "VOUCHER2", "BrandID": 1 }
]
```

---

## 3. Transactions

### GET /transactions

- **Response**:

```json
[
  { "ID": 1, "CustomerName": "John Doe", "TotalPoints": 100 },
  { "ID": 2, "CustomerName": "Jane Smith", "TotalPoints": 200 }
]
```

### GET /transactions/{id}

- **Response**:

```json
{
  "ID": 1,
  "CustomerName": "John Doe",
  "TotalPoints": 100
}
```

### POST /transactions

- **Request Body**:

```json
{
  "CustomerName": "Alice",
  "TotalPoints": 150
}
```

- **Response**:

```json
{
  "ID": 3,
  "CustomerName": "Alice",
  "TotalPoints": 150
}
```

---

## 4. Redemption Transactions

### POST /transactions/redemption

- **Request Body**:

```json
{
  "CustomerName": "Bob",
  "TotalPoints": 50,
  "Items": [{ "VoucherID": 1, "Quantity": 2, "TotalPoints": 40 }]
}
```

- **Response**:

```json
{
  "ID": 4,
  "CustomerName": "Bob",
  "TotalPoints": 50,
  "Items": [{ "VoucherID": 1, "Quantity": 2, "TotalPoints": 40 }]
}
```

### GET /transactions/redemption/{transactionId}

- **Response**:

```json
{
  "ID": 4,
  "CustomerName": "Bob",
  "TotalPoints": 50,
  "Items": [{ "VoucherID": 1, "Quantity": 2, "TotalPoints": 40 }]
}
```

---
