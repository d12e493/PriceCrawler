#!/bin/sh
OS_NAME=`uname`

PACKAGE_NAME="product-query"

if [[ "${OS_NAME}" == "Darwin" ]]; then
    SCRIPT_FOLDER=$(cd "$(dirname "$0")"; pwd)
    ARGS_GOOS="darwin"
    ARGS_GOARCH="amd64"
else
    SCRIPT_FOLDER=$(dirname $(readlink -f $0))
    ARGS_GOOS="linux"
    ARGS_GOARCH="amd64"
fi

echo "Script Folder: ${SCRIPT_FOLDER}"
if [[ "${SCRIPT_FOLDER}" == "" ]]; then 
    echo "Cannot get the folder path"
    exit 2
else
    mkdir -p ${SCRIPT_FOLDER}/src && \
    ln -fs ${SCRIPT_FOLDER} ${SCRIPT_FOLDER}/src/${PACKAGE_NAME} && \
    rm -f ${SCRIPT_FOLDER}/bin/* && \
    mkdir -p ${SCRIPT_FOLDER}/bin && \
    GOPATH=${SCRIPT_FOLDER} CGO_ENABLED=0 GOOS=${ARGS_GOOS} GOARCH=${ARGS_GOARCH} go build -o ${SCRIPT_FOLDER}/bin/api.o ${SCRIPT_FOLDER}/src/${PACKAGE_NAME}/main/api.go && \
    GOPATH=${SCRIPT_FOLDER} CGO_ENABLED=0 GOOS=${ARGS_GOOS} GOARCH=${ARGS_GOARCH} go build -o ${SCRIPT_FOLDER}/bin/batch.o ${SCRIPT_FOLDER}/src/${PACKAGE_NAME}/main/batch.go && \
    rm -rf ${SCRIPT_FOLDER}/src
fi

