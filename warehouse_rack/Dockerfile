FROM golang:1.23.1

# Set destination for COPY
WORKDIR /app

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/reference/dockerfile/#copy
COPY . ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o src/build/warehouse_rack src/main.go