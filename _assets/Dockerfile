FROM golang:1.13 as builder

WORKDIR /github.com/videocoin/go-bridge

COPY . .

RUN make build-vendor

FROM bitnami/minideb:jessie

COPY --from=builder /github.com/videocoin/go-bridge/build/bridge /usr/local/bin/

COPY --from=builder /github.com/videocoin/go-bridge/solidity /solidity

RUN install_packages curl ca-certificates git g++ make

RUN curl -sL https://deb.nodesource.com/setup_10.x | bash - && install_packages nodejs

RUN npm_config_user=root npm i -g truffle@5.1.21

RUN cd solidity && npm_config_user=root npm i @truffle/hdwallet-provider

RUN npm install -g openzeppelin-solidity

RUN npm i -g solc@0.5.16

CMD ["bridge"]