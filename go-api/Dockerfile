FROM golang:latest

WORKDIR /app

COPY . .

RUN go build -o main .

ENV DB_USERNAME=root
ENV DB_PASSWORD=mel1
ENV DB_HOST=cnt-database
ENV DB_PORT=3306
ENV DB_DATABASE=proyectohotel

EXPOSE 8080

CMD ["./main"]
