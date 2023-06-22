FROM golang:1.20 AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY Makefile main.go ./
COPY cmd/ ./cmd
COPY internal/ ./internal

RUN make build

FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build /app/build/house-facts /house-facts

EXPOSE 3000

USER nonroot:nonroot

ENTRYPOINT ["/house-facts"]
