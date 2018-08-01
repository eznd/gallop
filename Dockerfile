FROM openjdk
USER root

ENV GOPATH /go
ADD . /go/src/eznd/gallop

RUN apt-get update
RUN apt-get install -y golang

WORKDIR /tmp
RUN wget https://bintray.com/qameta/generic/download_file?file_path=io%2Fqameta%2Fallure%2Fallure%2F2.7.0%2Fallure-2.7.0.tgz -O allure.tar.gz
RUN gzip -d allure.tar.gz
RUN tar -xvf allure.tar
RUN mv allure-2.7.0/ /usr/local/
RUN ln -s /usr/local/allure-2.7.0/bin/allure /usr/bin/allure

WORKDIR $GOPATH/src/eznd/gallop
RUN go get -u github.com/golang/dep/cmd/dep
RUN /go/bin/dep ensure
RUN go build -o gallop

CMD ["./gallop"]
