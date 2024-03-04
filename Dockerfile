FROM golang:bookworm

RUN mkdir -p /app/src

WORKDIR /app/src





COPY . .

EXPOSE 8080

CMD ["golang", "./main.go"]