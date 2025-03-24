# Backend Engineer  Take-Home Assessment 
A basic backend service for processing receipt data and calculating reward points based on what's inside the receipt.

## Technical Stack Details

--> Language used : Golang
--> API Framework is net/http
--> Used Docker for Containerization
--> Used Postman for testing

### Prerequisites details:

--> Install Golang on your device at https://go.dev/ 
--> Install Docker on your Device at https://www.docker.com/products/docker-desktop
--> Install Postman on your Device at https://www.postman.com/downloads/

### We can Run it Locally using the below commands in your command prompt

```bash
go mod tidy
go run main.go
```

The service will be available at http://localhost:8080

## (or)

### We can Run with the Docker tool

--> Launch the docker tool and keep it running 
--> From te same directory where Dockerfile is saved follow the below commands:

```bash
docker build -t receipt-processor .
docker run -p 8080:8080 receipt-processor
```

##  API Endpoints as follows In Postman:

### 1. `POST /receipts/process`

It take a receipt in JSON format and returns a unique ID for that receipt.

#### 📦 Example Request

```
{
  "retailer": "Target",
  "purchaseDate": "2022-01-01",
  "purchaseTime": "13:01",
  "items": [
    { "shortDescription": "Mountain Dew 12PK", "price": "6.49" },
    { "shortDescription": "Emils Cheese Pizza", "price": "12.25" },
    { "shortDescription": "Knorr Creamy Chicken", "price": "1.26" },
    { "shortDescription": "Doritos Nacho Cheese", "price": "3.35" },
    { "shortDescription": "Klarbrunn 12-PK 12 FL OZ", "price": "12.00" }
  ],
  "total": "35.35"
}
```

#### The Response will look something like this

```json
{
  "id": "94bec470-5456-4746-9e14-69a038f45bbb" 
}
```

### GET /receipts/{id}/points

You can use the receipt ID to determine how many points were gained from that specific receipt.

### The Response will look something like this:

```json
{
  "points": 28
}
```