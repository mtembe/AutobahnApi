FROM golang:bookworm

RUN mkdir -p /app/src 

WORKDIR /app/src

ENTRYPOINT ["go run main.go -roads A1,A3,A7"]





COPY . .

EXPOSE 8080

CMD ["Autobahn"]
