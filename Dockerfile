FROM node:24 AS nodebuild
COPY . /work
WORKDIR /work
RUN npm install -g pnpm
RUN make web/dist

FROM golang:1.25 AS gobuild
COPY --from=nodebuild /work /go/src/delegator
WORKDIR /go/src/delegator
RUN make

FROM alpine:latest
RUN apk add tzdata
COPY --from=gobuild /go/src/delegator/delegator /usr/bin/delegator
EXPOSE 3000
CMD ["/usr/bin/delegator"]
