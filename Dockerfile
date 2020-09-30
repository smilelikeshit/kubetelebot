FROM golang:1.12.17-alpine3.11 AS build
RUN apk --no-cache add gcc g++ make ca-certificates git
WORKDIR /go/src/github.com/smilelikeshit/kubetelebot

COPY . .
RUN go install ./...

# 5 build the GO program
RUN CGO_ENABLED=0 GOOS=linux

FROM alpine:3.7
RUN apk update \
    && apk upgrade \
    && apk add --no-cache \
    ca-certificates \
    && update-ca-certificates 2>/dev/null || true
WORKDIR /usr/bin
COPY --from=build /go/bin .
CMD ["kubetelebot"]