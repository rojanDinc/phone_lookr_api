FROM golang:1.14

WORKDIR /app

COPY . .

EXPOSE 8080

ENV PORT 8080

RUN go install -v

RUN go build -o app .

CMD [ "./app" ]