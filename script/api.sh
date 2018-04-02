#!/bin/sh
OS_NAME=`uname`
PROGRAM="api"

PROCESS=`ps aux|grep /bin/api.o|grep -v "grep"|awk '{print $2}'`

if [[ "${OS_NAME}" == "Darwin" ]]; then
    SCRIPT_FOLDER=$(cd "$(dirname "$0")"; pwd)
else
    SCRIPT_FOLDER=$(dirname $(readlink -f $0))
fi
start(){
    SCRIPT_ARGS="--config-path ${SCRIPT_FOLDER}/../config/config.json --log-config-path ${SCRIPT_FOLDER}/../config/logger.json "
    if [ -f ${SCRIPT_FOLDER}/../bin/api.o ]; then

        echo "api start"

        GIN_MODE=release ${SCRIPT_FOLDER}/../bin/api.o ${SCRIPT_ARGS} &
        exit $?
    fi
}

stop() {
    if [ -z $PROCESS ]; then
        echo "$PROGRAM is not running. Exit."
        failure
    else
        kill $PROCESS
        sleep 2
        echo "api stop"
    fi
    echo 
    return 0
}

status () {
    if [ -z $PROCESS ]; then
            echo "$PROGRAM is not running. "
    else
            echo "$PROGRAM is running.(PID:$PROCESS)"
    fi
}

case "$1" in
  start)
	start
	;;
  stop)
	stop
	;;
  status)
    status 
	;;
  restart)
    stop
    start
    ;;	
  *)
	echo $"Usage: $PROGRAM {start|stop|restart|status}"
	exit 1
esac
exit 0
