ARG BUILDER_IMAGE=golang:1.19.1-alpine3.16
ARG BASE_IMAGE=alpine:3.16

FROM --platform=$BUILDPLATFORM $BUILDER_IMAGE AS build
ENV CGO_ENABLED=0
WORKDIR /go/src/

RUN mkdir common && mkdir create-atoms

COPY common/go.mod common/
COPY create-atoms/go.mod create-atoms/go.sum create-atoms/
RUN cd create-atoms && go mod download && cd ..

COPY . .
ARG TARGETOS TARGETARCH
RUN cd create-atoms/ && GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o /usr/local/bin/function

#############################################

FROM $BASE_IMAGE
COPY --from=build /usr/local/bin/function /usr/local/bin/function
COPY create-atoms/elements.yaml /etc/elements.yaml
ENTRYPOINT ["function"]
