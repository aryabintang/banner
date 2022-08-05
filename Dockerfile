<<<<<<< HEAD
FROM golang:latest

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64


WORKDIR /build


COPY go.mod .
COPY go.sum .
COPY . .
RUN go mod download


COPY . .


RUN go build -o golang_cms


WORKDIR /dist


RUN cp /build/main .
COPY .env /dist

EXPOSE 27017


CMD ./golang_cms
=======
FROM golang:latest

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64


WORKDIR /build


COPY go.mod .
COPY go.sum .
RUN go mod download


COPY . .


RUN go build -o /golang_cms.


WORKDIR /dist


RUN cp /build/main .
COPY .env /dist

EXPOSE 8787


CMD ["golang_cms"]
>>>>>>> 4027e718a4ebd2b380127196ae3ea524fddafa6f
