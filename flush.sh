#!/bin/bash

# 获取当前目录
current_dir=$(pwd)

# 遍历当前目录下的所有子目录
for service_dir in "$current_dir"/*; do
    if [ -d "$service_dir" ]; then
        # 进入子目录
        cd "$service_dir"

        echo "$service_dir"

        # 删除 output 目录
        rm -rf output

        # # 执行 build.sh 和 bootstrap.sh
        # if [ -f "build.sh" ]; then
        #     sh build.sh
        # fi

        # if [ -f "output/bootstrap.sh" ]; then
        #     sh output/bootstrap.sh
        # fi

        # 返回上级目录
        cd ..
    fi
done