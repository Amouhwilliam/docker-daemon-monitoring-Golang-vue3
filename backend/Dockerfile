FROM golang:1.22-alpine3.19 as build
WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN CGO_ENABLED=0 go build -o server .


FROM alpine:3.19
COPY --from=build /build/server /server

ENTRYPOINT ["/server"]