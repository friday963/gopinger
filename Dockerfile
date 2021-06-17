FROM golang:1.16 AS builder

ADD . /src

RUN cd /src && go build -o goapp 

FROM debian:buster-slim
WORKDIR /app 
COPY --from=builder /src/goapp /app/
ENTRYPOINT [ "./goapp" ] 
# Create a layered image to cut down on the size of the build.