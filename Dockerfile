FROM golang:1.19-alpine3.16

# RUN cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && echo 'Asia/Shanghai' >/etc/timezo


WORKDIR /code
ADD .  /code
ADD config  /root/.kube/


# RUN go env -w GO111MODULE=on \
#     && go env -w GOPROXY=https://goproxy.cn,https://goproxy.io,direct \
#     && go mod tidy \
#     && go build -tags netgo 

CMD ["/code/testapi"]

