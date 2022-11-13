FROM golang:1.19


COPY . /root

WORKDIR /root

RUN go build -o collegeadminstration /root/main.go

EXPOSE 5050

CMD .