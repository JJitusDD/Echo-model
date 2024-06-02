# stage 1
FROM docker-hub.g-pay.vn/baseimages/golang:1.20.1-alpine as build

WORKDIR /build

ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

COPY . .

RUN go build -mod vendor -tags musl -o  server ./cmd/server/main.go
RUN go build -mod vendor -tags musl -o  worker ./cmd/worker/worker.go
RUN go build -mod vendor -tags musl -o  kafka-consumer ./cmd/consumer/consumer.go
RUN go build -mod vendor -tags musl -o  schedule-job ./cmd/schedule/job-model.go


# stage 2
FROM docker-hub.g-pay.vn/baseimages/alpine:3.12

WORKDIR /app

COPY bin/run bin/run
COPY --from=build /build/.reconcileData ./.reconcileData
COPY --from=build /build/server .
COPY --from=build /build/worker .
COPY --from=build /build/kafka-consumer .
COPY --from=build /build/schedule-get-napas-file .
COPY --from=build /build/job-update-reconcile-wallet .
COPY --from=build /build/job-export-file-for-f88 .


RUN  chmod +x bin/run

CMD ["bin/run", "server"]
