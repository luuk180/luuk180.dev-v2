FROM golang:alpine
WORKDIR /app

COPY ./ ./
COPY ./files ./files

RUN go build -o main .

EXPOSE 3000

CMD ["./main"]