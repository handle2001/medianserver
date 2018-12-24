FROM iron/go:dev
WORKDIR /app
ENV GOPATH=/go/
RUN go get -v github.com/handle2001/median
RUN go get -v github.com/handle2001/medianserver
RUN cd $GOPATH/src/github.com/handle2001/median; go build; go install
RUN cd $GOPATH/src/github.com/handle2001/medianserver; go build -o medianserver; cp -v medianserver /app/ 
ENTRYPOINT ["./medianserver"]
