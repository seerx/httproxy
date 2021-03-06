#!/usr/bin/env bash

export GOOS=linux
export CGO_ENABLE=0

# 获取根脚本所在目录
sh_path=$(cd `dirname $0`; pwd)
cd $sh_path

echo 'Current Directory'
echo `pwd`

echo 'Building application ...'

go build -o ./proxy-amd64 \
    ../cmd/main.go

export GOOS=darwin

echo "Stoping service ..."
sshpass -p $PWD_OF_ALIYUN ssh root@xval.cn "systemctl stop myproxy"

echo "Uploading ..."
sshpass -p $PWD_OF_ALIYUN scp ./proxy-amd64 root@xval.cn:/usr/local/proxy/

echo "Starting service ..."
sshpass -p $PWD_OF_ALIYUN ssh root@xval.cn "systemctl start myproxy"

rm -f ./proxy-amd64

echo "DONE!"