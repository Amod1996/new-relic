FROM golang:1.20-alpine

WORKDIR /app
COPY server/ ./

COPY go.mod ./
COPY go.sum ./
RUN go mod download


RUN go build  -o three_word_occurance_calculator

EXPOSE 8080

CMD ["./three_word_occurance_calculator","file1.txt"]