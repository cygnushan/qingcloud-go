# Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
# Use of this source code is governed by a Apache
# license that can be found in the LICENSE file.

# docker run --rm -it -v `pwd`:/root -w /root chai2010/qingcloud-go qcli

FROM golang:1.9.2-alpine3.6 as builder

WORKDIR /go/src/github.com/chai2010/qingcloud-go/
COPY . .
RUN go install ./cmd/...

FROM alpine:3.6

COPY --from=builder /go/bin/qcli /usr/local/bin/qcli

ENTRYPOINT []
CMD ["/usr/local/bin/qcli"]
