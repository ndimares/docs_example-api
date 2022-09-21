# syntax=docker/dockerfile:1

# Build the application

FROM golang:1.18 AS build

WORKDIR /app

ARG GIT_USERNAME
ARG TOKEN

COPY ./ ./

ENV GOPRIVATE=github.com/speakeasy-api
RUN --mount=type=secret,id=git_auth_token ./scripts/git_access.sh

RUN go mod download

RUN ls -lrt

RUN CGO_ENABLED=0 GOOS=linux go build -v -o ./server ./cmd/server/main.go 

# Build the server image

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=build /app/server ./
COPY --from=build /app/config/ ./config/
COPY --from=build /app/migrations/ ./migrations/
COPY --from=build /app/openapi.yml ./

CMD ["./server"]