FROM golang:1.19 as builder
LABEL maintainer="Erodotos Demetriou <erodotosd@gmail.com>"
WORKDIR /app

# Copy go mod and sum files
COPY go.mod ./

# Download all dependencies.
RUN go mod tidy
RUN go mod download

COPY main.go .

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





# Toggle CGO based on your app requirement. CGO_ENABLED=1 for enabling CGO
RUN CGO_ENABLED=0 go build -ldflags '-s -w -extldflags "-static"' -o /app/appbin *.go
# Use below if using vendor
# RUN CGO_ENABLED=0 go build -mod=vendor -ldflags '-s -w -extldflags "-static"' -o /app/appbin *.go

FROM scratch
LABEL MAINTAINER Author <author@example.com>

COPY --from=builder /app /app

WORKDIR /app
# Since running as a non-root user, port bindings < 1024 are not possible
# 8000 for HTTP; 8443 for HTTPS;
EXPOSE 8000
EXPOSE 8443

CMD ["./appbin"]
