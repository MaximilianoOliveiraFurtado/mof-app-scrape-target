FROM golang:alpine AS builder

WORKDIR /app

# dependencies cache 
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o mof-app-scrape-target .


FROM alpine:latest  

WORKDIR /root/

COPY --from=builder /app/mof-app-scrape-target .
CMD ["./mof-app-scrape-target"]