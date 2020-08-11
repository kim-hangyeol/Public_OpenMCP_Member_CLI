#!/bin/bash

export GO111MODULE=on
go mod vendor

EXTERNAL_IP="10.0.3.12"

go build -o omctl && \
cp omctl /usr/local/bin && \
mkdir -p /var/lib/omctl && \
cp config.yaml /var/lib/omctl/config.yaml &&
sed -i 's/<YOUR_EXTERNAL_IP>/'${EXTERNAL_IP}'/g' /var/lib/omctl/config.yaml
