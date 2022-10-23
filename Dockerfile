FROM golang:1.19.2-alpine AS builder

COPY . /build
WORKDIR /build

# build as PIE to take advantage of exploit mitigations
ARG CGO_ENABLED=0
ARG VERSION
RUN go build -buildmode pie -ldflags "-s -w -X main.version=${VERSION}" -trimpath -o go-project-template

FROM ghcr.io/capnspacehook/pie-loader
COPY --from=builder /build/go-project-template /go-project-template

# apparently giving capabilities to containers doesn't work when the
# container isn't running as root inside the container, see
# https://github.com/moby/moby/issues/8460

ENTRYPOINT [ "/go-project-template" ]
