# Stage 1: Build frontend (Tailwind CSS and templ files)
FROM node:18 AS frontend-builder

# Set the working directory
WORKDIR /app

# Install package manager (pnpm)
RUN npm install -g pnpm

# Copy package.json and pnpm-lock.yaml and install dependencies
COPY package.json pnpm-lock.yaml ./
RUN pnpm install

# Copy Tailwind and templ configuration and source files
COPY ./views ./views
COPY ./components ./components
COPY ./tailwind.config.js ./
COPY ./static ./static

# Build Tailwind CSS
RUN pnpm run build:css



# Stage 2: Build Go application
FROM golang:1.23 AS go-builder

# Set the working directory
WORKDIR /app

# Set Go proxy environment variable
ENV GOPROXY=https://proxy.golang.org


# Copy go.mod and go.sum and download Go dependencies
COPY go.mod go.sum ./
RUN go mod tidy
RUN go mod download -x


# Copy the rest of the application source code
COPY . .

RUN go install github.com/a-h/templ/cmd/templ@latest

ENV PATH=$PATH:/go/bin

# Generate templ files
RUN templ generate -path ./views
RUN templ generate -path ./components

# Copy the built CSS and templ files from the frontend-builder stage
COPY --from=frontend-builder /app ./ 

# Build the Go application
RUN go build -o jobin cmd/server/main.go

# Stage 3: Run the Go application
FROM debian:bookworm-slim

# Set the working directory
WORKDIR /app

# Copy the built Go binary from the go-builder stage
COPY --from=go-builder /app/jobin /app/jobin

COPY --from=go-builder /app/static /app/static

# Expose the application's port
EXPOSE 3000

# Run the Go application
CMD ["/app/jobin"]
