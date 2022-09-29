FROM golang:latest

LABEL maintainer ="Htin Lynn <htinlin01@gmail.com>"

WORKDIR /usr/app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build

EXPOSE 8080

CMD ["./final-year-project-mcc"]



