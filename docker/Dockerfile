FROM golang:1.17-buster as builder

WORKDIR /go/src/app

COPY . .
RUN go mod download

RUN go build -v -o api cmd/api/main.go

#RUN migrate -source file://database/migrations -database "mysql://root1:root1@tcp(go_api_database:3306)/go_api_database" up 2
#RUN docker run -v database:/migrations --network host migrate/migrate -path=database/migrations/ -database "mysql://root1:root1@tcp(go_api_database:3306)/go_api_database" up 2

FROM builder as prod

WORKDIR /bin

COPY --from=builder /go/src/app/api .
COPY --from=builder /go/src/app/.env .
COPY --from=builder /go/src/app/database/. database/.

CMD api

FROM builder as dev

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

WORKDIR /go/src/app

CMD ["air", "-c", ".air.toml"]