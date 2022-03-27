FROM golang:1.17-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
COPY . ./

RUN go test ./...
RUN go build -o /lmwn-covid-19

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /lmwn-covid-19 /lmwn-covid-19

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/lmwn-covid-19"]