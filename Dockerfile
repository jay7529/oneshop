FROM golang
WORKDIR /oneshop
COPY . .
RUN go mod download
RUN go build -o main .
EXPOSE 8080:8080
CMD ["./main"]