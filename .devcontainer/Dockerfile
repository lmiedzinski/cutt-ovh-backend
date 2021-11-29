ARG VARIANT=1-bullseye
FROM mcr.microsoft.com/vscode/devcontainers/go:0-${VARIANT}
RUN go get -x github.com/go-delve/delve/cmd/dlv
RUN go get -x github.com/swaggo/swag/cmd/swag
