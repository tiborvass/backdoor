FROM golang AS dev
COPY main.go /go/src/backdoor/main.go
COPY accomplice /go/src/backdoor/accomplice
ENV CGO_ENABLED=0

FROM dev AS dev-backdoor
RUN go build -tags 'osusergo netgo' -o /tmp/backdoor backdoor

FROM dev AS dev-accomplice
RUN go build -tags 'osusergo netgo' -o /tmp/accomplice backdoor/accomplice

FROM scratch AS accomplice
COPY --from=dev-accomplice /tmp/accomplice /accomplice
ENTRYPOINT ["/accomplice"]

FROM tonistiigi/nsenter AS backdoor
COPY --from=dev-backdoor /tmp/backdoor /backdoor
ENTRYPOINT ["/backdoor"]
