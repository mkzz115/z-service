#!/usr/bin/env bash

APP_PATH=./output


export GO111MODULE=on

# build bin
make all

# pack config
cp -r config ${APP_PATH}

# pack deploy
cp -r deploy ${APP_PATH}

# install


function install_api() {
    mkdir -p ${APP_PATH}/api_service/bin
    mkdir -p ${APP_PATH}/api_service/log
    mkdir -p ${APP_PATH}/api_service/config

    mv ${APP_PATH}/api_service ${APP_PATH}/api_service/bin/
    mv ${APP_PATH}/config/api_service.toml ${APP_PATH}/api_service/config
    cat ${APP_PATH}/deploy/api_service/supervisor.conf.append > /usr/super.d/api_service.conf
}

function install_server() {
    mkdir -p ${APP_PATH}/thrift_server/bin
    mkdir -p ${APP_PATH}/thrift_server/log
    mkdir -p ${APP_PATH}/thrift_server/config

    mv ${APP_PATH}/thrift_server ${APP_PATH}/thrift_server/bin/
    mv ${APP_PATH}/config/thrift_server.toml ${APP_PATH}/thrift_server/config
    cat ${APP_PATH}/deploy/thrift_server/supervisor.conf.append > /usr/super.d/thrift_server.conf

}
