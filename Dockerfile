FROM golang
WORKDIR /oneshop
COPY . .
RUN go mod download
RUN go build -o main .
EXPOSE 8000
CMD ["./main"]