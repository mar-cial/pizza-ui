FROM golang:1.23 AS builder

WORKDIR /ui-builder

COPY go.mod go.sum ./

ARG port

ENV UI_PORT=${port}

RUN go mod download

COPY . .

RUN go install github.com/a-h/templ/cmd/templ@latest

RUN templ generate 

RUN go build -o ui-server cmd/main.go

FROM gcr.io/distroless/base-debian12

COPY --from=builder /ui-builder/ui-server /ui-server

EXPOSE ${port}

ENTRYPOINT [ "/ui-server" ]