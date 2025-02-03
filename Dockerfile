FROM golang:1.19


WORKDIR /root/backEnd
COPY . /root

RUN go build -o collegeadminstration ./main.go

EXPOSE 5050

CMD [ "/root/collegeadminstration" ]