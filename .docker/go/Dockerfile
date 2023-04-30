ARG USER_NAME='app'
ARG USER_UID='1000'
ARG USER_GID='1000'
ARG USER_GROUP='app_users'
ARG REPO='{{.repoOrg}}/{{.codeName}}'

FROM golang:1.20-alpine3.17 AS dev

ARG USER_NAME
ARG USER_UID
ARG USER_GID
ARG USER_GROUP
ARG REPO

ENV GOPATH="/home/${USER_NAME}"
ENV CGO_ENABLED=0
ENV WORK_DIR="/home/${USER_NAME}/src/${REPO}"

WORKDIR /tmp

COPY --chmod=0775 .docker/go/start.sh /usr/local/bin

# VS Code Requirements: openssh, musl, libgcc, libstdc++
RUN env && apk --no-progress --purge --no-cache upgrade \
 && apk --no-progress --purge --no-cache add --upgrade \
    curl \
    git \
    gnupg \
    libgcc \
    libstdc++ \
    openssh \
    openssl \
 && apk --no-progress --purge --no-cache upgrade \
 && rm -vrf /var/cache/apk/*

# Add a non-root group and user, helpful if you dev on Linux.
RUN addgroup --system --gid ${USER_GID} ${USER_GROUP} \
 && adduser --system \
    --disabled-password \
    --ingroup ${USER_GROUP} \
    --uid ${USER_UID} \
    ${USER_NAME}

USER ${USER_NAME}

# Install VS Code Server.
COPY --chmod=0775 .devcontainer/download-vs-code-server.sh .
RUN ./download-vs-code-server.sh

RUN mkdir -p ~/bin \
 && curl -L -o ~/bin/git-chglog https://github.com/git-chglog/git-chglog/releases/download/0.9.1/git-chglog_linux_amd64 \
 && chmod +x ~/bin/git-chglog

# Install Go helpful dev tools.
RUN go install golang.org/x/tools/gopls@latest \
 && go install github.com/go-delve/delve/cmd/dlv@latest
# && go github.com/ramya-rao-a/go-outline@latest \
# && go github.com/cweill/gotests/...@latest

# VSCode Requirements for pre-installing extensions
RUN mkdir -p /home/${USER_NAME}/.vscode-server/extensions

WORKDIR $WORK_DIR

ENTRYPOINT [ "start.sh" ]

CMD [ "" ]

FROM dev AS build

ENV GOPATH /home/${USER_NAME}

COPY ./ /home/${USER_NAME}/src

RUN go build

FROM alpine3.17 AS release

COPY  --chown=${USER_NAME}:${USER_GROUP} --from=build /home/${USER_NAME}/bin /usr/local/bin

ENTRYPOINT [ "diablo-assistant" ]
