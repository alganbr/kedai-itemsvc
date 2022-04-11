FROM golang:1.17-alpine AS builder

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0

# Move to working directory /build
WORKDIR /build

# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the code into the container
COPY . .

# Install dependencies
RUN apk add --no-cache make

# Generate files
RUN make install generate

# Build the application
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -a -o itemsvc -tags musl

FROM alpine

# Move to /dist directory as the place for resulting binary folder
WORKDIR /app

# Copy binary from build to app folder
COPY --from=builder /build/docs ./docs
COPY --from=builder /build/configs/env.yml ./configs/env.yml
COPY --from=builder /build/internal/migrations ./migrations
COPY --from=builder /build/itemsvc ./

# Export necessary port
EXPOSE 8080

# Command to run when starting the container
ENTRYPOINT ["./itemsvc"]