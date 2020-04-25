FROM golang:1.14-alpine AS builder

ARG BUILD_DATE
ARG VCS_REF
ARG VERSION

ENV GO111MODULE=on

RUN mkdir -p /src/app
WORKDIR /src/app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
	-ldflags="-X 'github.com/brpaz/do-snapshot-pruner/cmd.buildVersion=${VERSION}' -X 'github.com/brpaz/do-snapshot-pruner/cmd.buildDate=${BUILD_DATE}' -X 'github.com/brpaz/do-snapshot-pruner/cmd.buildCommit=${VCS_REF}'" \
	-o build/main main.go

FROM alpine:3.11 AS final

ARG BUILD_DATE
ARG VCS_REF
ARG VERSION

COPY --from=builder /src/app/build/main /bin/do-snapshot-pruner

ENTRYPOINT ["/bin/do-snapshot-pruner"]

LABEL maintainer="Bruno Paz <oss@brunopaz.net>" \
	org.opencontainers.image.title="DigitalOcean Snapshot Pruner" \
	org.opencontainers.image.description="Deletes old snapshots from DigitalOcean" \
	org.opencontainers.image.url="https://github.com/brpaz/do-snapshot-pruner" \
	org.opencontainers.image.source="git@github.com:brpaz/do-snapshot-pruner" \
	org.opencontainers.image.name="do-snpashot-pruner" \
	org.opencontainers.image.created=${BUILD_DATE} \
	org.opencontainers.image.revision=${VCS_REF} \
	org.opencontainers.image.version=${VERSION}
