# Copyright (C) 2021 Nicolas Lamirault <nicolas.lamirault@gmail.com>
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

FROM golang:1.23 as build-env

WORKDIR /go/src/app
ADD . /go/src/app

RUN go get -d -v ./...

RUN go build -o /go/bin/bbox_exporter

#####

FROM gcr.io/distroless/base

LABEL maintainer="" \
    org.opencontainers.image.authors="The Bbox Prometheus Exporter Authors" \
    org.opencontainers.image.title="gcr.io/nlamirault/bbox_exporter" \
	org.opencontainers.image.description="A Prometheus exporter for the Bbox, a Set-Top-Box (TV box) provided by French Internet Service Provider Bouygues Telecom" \
	org.opencontainers.image.documentations="https://github.com/nlamirault/bbox_exporter" \
    org.opencontainers.image.url="https://github.com/nlamirault/bbox_exporter" \
	org.opencontainers.image.source="git@github.com:nlamirault/bbox_exporter" \
    org.opencontainers.image.licenses="Apache 2.0" \
    org.opencontainers.image.vendor=""

COPY --from=build-env /go/bin/bbox_exporter /
# set the uid as an integer for compatibility with runAsNonRoot in Kubernetes
USER 65534:65534
CMD ["/bbox_exporter"]
