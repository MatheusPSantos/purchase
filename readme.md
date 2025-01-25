# Wex Tech Test

This repository contains a technical test developed in Go, using the Gin framework and the CORM ORM.

## Download and Setup Instructions

1. Clone this repository to your local machine:

```bash
git clone <REPOSITORY_URL>
cd wex-tech-test
```

1. Make sure you have Go installed on your machine. To check, run:

```
go version
```

## How to Run

### Method 1: Manual Execution
1. Build the application:
```bash
go build -o app . ```
2. Start the database and PgAdmin services using Docker:
```bash
docker-compose up -d
```

3. Make sure the environment variables are set correctly:

• BANK: wex
• USER: wex
• PASSWORD: wex

4. Run the application:
```bash
./app
```

The application will be available at http://localhost:8888.

### Method 2: Simplified Execution with Script
1. Check if the start.sh script has execution permission:
```bash
chmod +x start.sh
```
2. Run the script:
```bash
./start.sh
```

The application will be started automatically, including the database and PgAdmin.

### Available Endpoints
> 1. Create a new transaction (POST /purchase)

Example request:
```bash
curl --request POST \
--url http://localhost:8888/purchase \
--header 'Content-Type: application/json' \
--data '{
"description": "stone psr transaction",
"transaction_date": "2024-09-13",
"amount": 1
}'
```

> 2. Retrieve a transaction by ID with exchange rate (GET /purchase/{id})

Retrieves the details of a transaction, informing the desired currency for calculating the exchange rate.

Request example:
```bash
curl --request GET \
--url 'http://localhost:8888/purchase/1?currency=Brazil-Real'
```

> 3. List all transactions (GET /purchase)

Lists all transactions registered in the system.

Request example:
```bash
curl --request GET \
--url http://localhost:8888/purchase
```

### Accessing PgAdmin
1. Access PgAdmin at http://localhost:54321.
2. Use the following credentials:
• Email: teste@wex.com
• Password: testewex
3. To connect to the database, you may need the PostgreSQL container IP. Use the command below to get the IP:

```bash
docker inspect -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' postgres
```

Database Details:
- Database: wex
- User: wex
- Password: wex

## Notes

This repository contains an answer to a technical quiz. Feel free to explore and adapt as needed.