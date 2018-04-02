#!/bin/sh
OS_NAME=`uname`

if [[ "${OS_NAME}" == "Darwin" ]]; then
    SCRIPT_FOLDER=$(cd "$(dirname "$0")"; pwd)
else
    SCRIPT_FOLDER=$(dirname $(readlink -f $0))
fi

SCRIPT_ARGS="--job-name priceCrawler  --config-path ${SCRIPT_FOLDER}/../config/config.json --log-config-path ${SCRIPT_FOLDER}/../config/logger.json "
if [ -f ${SCRIPT_FOLDER}/../bin/batch.o ]; then
    GIN_MODE=release ${SCRIPT_FOLDER}/../bin/batch.o ${SCRIPT_ARGS} &
    exit $?
fi
