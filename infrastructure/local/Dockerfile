FROM golang:1.22 as build
WORKDIR /build

env GOTOOLCHAIN=auto

#RUN apk add --no-cache gcc musl-dev make
#RUN #go install github.com/cosmtrek/air@latest

#COPY go.mod go.sum ./
#RUN go mod download
#COPY infrastructure/local/.env .
#COPY ../../air.toml .
#COPY .bin .

COPY . .

RUN make build

FROM scratch as service

COPY --from=build /build/bin/macrod /app/macrod

CMD ["/app/macrod"]