# Stage 1: Build stage
FROM golang:1.21.5 AS builder

WORKDIR /backend/

COPY ./cmd /backend/cmd
COPY ./internal /backend/internal
COPY ./docs /backend/docs
COPY ./dev /backend/dev

COPY ./go.mod /backend/go.mod
COPY ./go.sum /backend/go.sum

RUN go build -o /backend/build ./cmd/effective-mobile-test/

# Stage 2: Final stage
FROM ubuntu:22.04

WORKDIR /backend

COPY --from=builder /backend/build /backend/build
COPY --from=builder /backend/docs /backend/docs
COPY --from=builder /backend/dev/.env /backend/dev/.env
COPY --from=builder /backend/internal/app/migrations /backend/internal/app/migrations

CMD [ "/backend/build" ]
