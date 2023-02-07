#!/usr/bin/env bash
# shellcheck disable=SC2046,SC2206
set -o errexit

export powered_by="https://github.com/jeyrce/gin-app"
export flag=".init_ok"
export build_files=(
    Makefile
    go.mod
)

# step: 从输入读取module配置
tmpl="$(echo $powered_by | awk -F// '{print $NF}')"
module=""
if [ -f "$flag" ]; then
    echo "项目已经执行过初始化: "
    cat "$flag"
    exit 0
else 
    read -rp "请输入项目module地址($tmpl): " module
fi
if [[ ! "$module" =~ ^[a-zA-Z0-9_][-a-zA-Z0-9_]{1,62}\.[a-zA-Z0-9_][-a-zA-Z0-9_]{1,62}\/.+\/.+$ ]]; then
    echo "不合法的module!"
    exit 1
fi

# step: app.go 重命名
app_name="$(echo "${module//-/_}" | awk -F/ '{print $NF}' |tr '[A-Z]' '[a-z]')"
mv app.go "$app_name".go

# step: 各go文件导入路径变更 module
go_files=$(find $(pwd) -type f -name "*.go")
files=(${build_files[*]} ${go_files[*]})
for file in ${files[*]}; do
    if [ -f "$file" ]; then
        if [ "$(uname)" == "Linux" ]; then
            sed -i "s#${tmpl}#${module}#g" "$file"
        elif [ "$(uname)" == "Darwin" ]; then
            sed -i "" "s#${tmpl}#${module}#g" "$file"
        fi 
    else 
        echo "$file 文件不存在跳过"
    fi
done 

# step: 设置授权
author=$(echo "${module}" | awk -F/ '{print $2}')
year=$(date "+%Y")
if [ "$(uname)" == "Linux" ]; then
    sed -i "s#2022~#${year}~#g" LICENSE
    sed -i "s#Jeyrce.Lu#${author}#g" LICENSE
elif [ "$(uname)" == "Darwin" ]; then
    sed -i "" "s#2022~#${year}~#g" LICENSE
    sed -i "" "s#Jeyrce.Lu#${author}#g" LICENSE
fi

# step: 初始化依赖
go mod download
go install github.com/swaggo/swag/cmd/swag@latest

# step: 设置执行过初始化的标识
echo "Init by $powered_by" > "$flag"
date >> "$flag"
git add .
git commit -m "init ok"

echo "项目初始化完成"

exit 0
