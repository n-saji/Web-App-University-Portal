FROM golang:1.19


WORKDIR /root

RUN ls

COPY ./backEnd /root/

RUN ls

RUN go mod download

RUN go build -o collegeadminstration ./main.go

EXPOSE 5050

CMD ["./collegeadminstration"]