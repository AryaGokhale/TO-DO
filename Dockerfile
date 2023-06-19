#Go base image 
FROM golang:alpine as builder

#Maintainer info
LABEL maintainer="Arya Gokhale <gokhale.arya@gmail.com>"

#Install git for fetching dependences
RUN apk update && apk add --no-cache git

#Set working directory 
WORKDIR /app

COPY go.mod go.sum ./

COPY . .


RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Start a new stage from scratch
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage. Observe we also copied the .env file
COPY --from=builder /app/main .
COPY --from=builder /app/.env .       

# Expose port 8080 to the outside world
EXPOSE 8080

#Command to run the executable
CMD ["./main"]