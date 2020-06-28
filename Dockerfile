# Build the application
FROM golang:1.14-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o mosquito

# Run the application in an empty alpine environment
FROM alpine:latest
WORKDIR /root
COPY --from=build /app/mosquito .
CMD ["./mosquito"]