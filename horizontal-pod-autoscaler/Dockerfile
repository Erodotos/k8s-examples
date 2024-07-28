################ STAGE - 1 ######################
FROM golang:1.20 as builder
LABEL maintainer="Erodotos Demetriou <erodotosd@gmail.com>"
WORKDIR /app

# Copy go mod and sum files
COPY go.mod ./
COPY main.go .

# Download all dependencies.
RUN go mod tidy
RUN go mod download


# Build the app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

################ STAGE - 2 ######################
FROM scratch

LABEL maintainer="Erodotos Demetriou <erodotosd@gmail.com>"

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Expose port 8080
EXPOSE 8080

# Start the App
CMD ["./main"]
