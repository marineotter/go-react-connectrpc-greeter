FROM debian:bookworm

# Install golang
ARG GO_VERSION=1.24.4

RUN apt-get update && \
    apt-get install -y curl gzip && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/* && \
    mkdir -p /temp && \
    cd /temp && \
    curl -O -L https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz && \
    rm -rf /temp/go${GO_VERSION}.linux-amd64.tar.gz
ENV PATH=$PATH:/usr/local/go/bin

# Install nodejs
ARG NVM_VERSION=0.40.3
ARG NODE_VERSION=24.2.0

ENV NVM_DIR=/usr/local/nvm
WORKDIR $NVM_DIR
RUN mkdir -p $NVM_DIR && \
    curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v${NVM_VERSION}/install.sh | bash && \
    . $NVM_DIR/nvm.sh && \
    nvm install ${NODE_VERSION} && \
    nvm alias default ${NODE_VERSION} && \
    nvm use default
ENV PATH=$PATH:$NVM_DIR/versions/node/v${NODE_VERSION}/bin

# Install additional tools
RUN apt-get update && \
    apt-get install -y \
    git \
    procps \
    lsof \
    && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

# Create user and group for development
ARG DEV_USER=devcontainer
ARG DEV_GROUP=devcontainer
ARG DEV_UID=1000
ARG DEV_GID=1000

RUN groupadd -g $DEV_GID $DEV_GROUP && \
    useradd -m -s /bin/bash -u $DEV_UID -g $DEV_GID $DEV_USER

USER $DEV_USER

# Install devtools
RUN go install github.com/go-delve/delve/cmd/dlv@latest && \
    go install github.com/bufbuild/buf/cmd/buf@latest && \
    go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest && \
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
    go install connectrpc.com/connect/cmd/protoc-gen-connect-go@latest
ENV PATH=$PATH:/home/$DEV_USER/go/bin
RUN npm install -g \
    @bufbuild/buf \
    @bufbuild/protoc-gen-connect-query \
    @bufbuild/protoc-gen-es
