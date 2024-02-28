FROM golang:1.21-alpine AS BuilStage

ENV CGO_ENABLED 0

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

COPY build/p2p_parser/ ./../p2p_parser/
COPY build/gojq_extention/ ./../gojq_extention/

RUN go mod download

COPY src/ ./src/

RUN go build -C src/main -o /app/p2p_parser_jq

FROM alpine:latest

WORKDIR /app

COPY --from=BuilStage /app/p2p_parser_jq ./p2p_parser_jq

ENTRYPOINT [ "./p2p_parser_jq" ]
