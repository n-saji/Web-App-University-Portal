FROM golang:1.19


WORKDIR /root
COPY . /root

RUN go build -o collegeadminstration /root/main.go

EXPOSE 5050

CMD [ "/root/collegeadminstration" ]