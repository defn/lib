VERSION 0.6

build:
    FROM registry.fly.io/defn:dev-tower
    COPY hello.go .
    RUN ~/bin/e go build hello.go
    SAVE ARTIFACT hello

image:
    FROM scratch
    COPY +build/hello /
    ENTRYPOINT ["/hello"]
    SAVE ARTIFACT /hello AS LOCAL hello
    SAVE IMAGE hello
