# 1. Use the authorized Go picture as a base.
FROM golang:1.24


# 2. Set the current working directory in the container.
WORKDIR /app

# 3. Copy the Go mod/sum scripts firstly for improved cache.
COPY go.mod go.sum ./
RUN go mod download

#4: Copy the remaining source files.
COPY . .

# 5. Create the Go app.
RUN go build -o main .

# 6: Open port 8080 to the outside world.
EXPOSE 8080

# 7: Run the binary
CMD ["./main"]
