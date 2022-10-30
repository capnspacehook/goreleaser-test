FROM golang:1.19.2-alpine AS builder

COPY . /build
WORKDIR /build

# add git so VCS info will be stamped in binary
RUN apk add --no-cache git-2.36.3-r0

# build as PIE to take advantage of exploit mitigations
ARG CGO_ENABLED=0
ARG VERSION
RUN go build -buildmode pie -ldflags "-s -w -X main.version=${VERSION}" -trimpath -o go-project-template

# pie-loader is built and scanned daily, we want the most recent version
# hadolint ignore=DL3006
FROM ghcr.io/capnspacehook/pie-loader
COPY --from=builder /build/go-project-template /go-project-template

USER 1000:1000

ENTRYPOINT [ "/go-project-template" ]
CMD [ "-version" ]
