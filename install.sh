#! /bin/bash
###
 # @Author: coderzhaolu && izhaicy@163.com
 # @Date: 2022-08-27 18:40:50
 # @LastEditors: coderzhaolu && izhaicy@163.com
 # @LastEditTime: 2022-08-29 19:04:42
 # @FilePath: /system-main/Users/zhaolu/Desktop/project/go_vue_gin/pinkmoe/install.sh
 # @Description: https://github.com/Coder-ZhaoLu/pinkmoe   (如需用于商业用途或者二开，请联系作者捐助任意金额即可)
 # QQ:2419857357;支付宝:13135986153
 # Copyright (c) 2022 by coderzhaolu, All Rights Reserved. 
### 
#!/bin/bash
#file:docker_install.sh

function docker_install()
{
	echo -e "\033[32m 检查Docker......\033[0m"
	docker -v
    if [ $? -eq  0 ]; then
        echo -e "\033[32m 检查到Docker已安装!\033[0m"
    else
    	echo -e "\033[32m 安装docker环境...\033[0m"
        curl -fsSL https://get.docker.com | bash -s docker --mirror Aliyun && systemctl enable docker && systemctl start docker && curl -L "https://github.com/docker/compose/releases/download/v2.7.0/docker-compose-$(uname -s)-$(uname -m)" > /usr/bin/docker-compose && chmod +x /usr/bin/docker-compose
        echo -e "\033[32m 安装docker环境...安装完成!\033[0m"
    fi
    # 创建公用网络==bridge模式
    #docker network create share_network
}

function write_conf()
{
    echo -e "\033[34m 请输入网站前台域名(多个域名用英文空格隔开,如:pinkmoe.com www.pinkmoe.com;直接回车默认为pinkmoe.com www.pinkmoe.com):\033[0m"
    cp ./pinkmoe_server/config.yaml.example ./pinkmoe_server/config.yaml
    read headDomain
    if [ -z $headDomain ]
    then
        headDomain="pinkmoe.com www.pinkmoe.com"
    fi
    echo $headDomain
    sed -i "s/pinkmoe.com www.pinkmoe.com/$headDomain/g" ./pinkmoe_index/nginx/pinkmoe_ssl.conf
    sed -i "s/pinkmoe.com www.pinkmoe.com/$headDomain/g" ./pinkmoe_index/nginx/pinkmoe.conf
    echo -e "\033[34m 请输入网站主域名(如:pinkmoe.com;直接回车默认为pinkmoe.com):\033[0m"
    read mainDomain
    if [ -z $mainDomain ]
    then
        mainDomain="pinkmoe.com"
    fi
    echo $mainDomain
    sed -i "s/pinkmoe.com/$mainDomain/g" ./pinkmoe_server/config.yaml
    echo "export const BASEURL = 'https://${mainDomain}:9527/api/v1'" > ./pinkmoe_index/constant.ts
    echo -e "\033[34m 请输入网站后台域名(多个域名用英文空格隔开,如:admin.pinkmoe.com;直接回车默认为admin.pinkmoe.com):\033[0m"
    read backgroundDomain
    if [ -z $backgroundDomain ]
    then
        backgroundDomain="admin.pinkmoe.com"
    fi
    echo $backgroundDomain
    sed -i "s/admin.pinkmoe.com/$backgroundDomain/g" ./pinkmoe_server/config.yaml
    sed -i "s/admin.pinkmoe.com/$backgroundDomain/g" ./pinkmoe_admin/nginx/admin_pinkmoe_ssl.conf
    sed -i "s/admin.pinkmoe.com/$backgroundDomain/g" ./pinkmoe_admin/nginx/admin_pinkmoe.conf
}

function write_ssl()
{
    read -d$ -p "请输入前台ssl域名的key文件字符串(直接文本打开key文件后复制到这里)，输入$符号退出以完成输入:" pinkmoeKey
    echo "$pinkmoeKey"
    echo "$pinkmoeKey" > ./pinkmoe_index/nginx/ssl/pinkmoe.key
    echo "$pinkmoeKey" > ./pinkmoe_server/ssl/pinkmoe.key
    read -d$ -p "请输入前台ssl域名的crt文件字符串(直接文本打开crt文件后复制到这里)，输入$符号退出以完成输入:" pinkmoeCrt
    echo "$pinkmoeCrt"
    echo "$pinkmoeCrt" > ./pinkmoe_index/nginx/ssl/pinkmoe.crt
    echo "$pinkmoeCrt" > ./pinkmoe_server/ssl/pinkmoe.crt

    read -d$ -p "请输入后台ssl域名的key文件字符串(直接文本打开key文件后复制到这里)，输入$符号退出以完成输入:" dminPinkmoeKey
    echo "$dminPinkmoeKey"
    echo "$dminPinkmoeKey" > ./pinkmoe_admin/nginx/ssl/admin.pinkmoe.key
    read -d$ -p "请输入后台ssl域名的crt文件字符串(直接文本打开crt文件后复制到这里)，输入$符号退出以完成输入:" adminPinkmoeCrt
    echo "$adminPinkmoeCrt"
    echo "$adminPinkmoeCrt" > ./pinkmoe_admin/nginx/ssl/admin.pinkmoe.crt
}

function install()
{
    echo -e "\033[32m 开始更新本机软件包:\033[0m"
    cd /
    sudo yum update -y
    docker_install
    echo -e "\033[32m 开始下载pinkmoe软件包:\033[0m"
    yum install git -y && git clone https://github.com/Coder-ZhaoLu/pinkmoe && cd pinkmoe/ 
    # 配置网站信息
    echo -e "\033[32m 开始配置网站基本信息:\033[0m"
    write_conf
    # 配置ssl文件
    write_ssl
}
echo -e "\033[32m #########################################################################\033[0m"
echo -e "\033[32m ##                         欢迎使用PinkMoe程序                         ##\033[0m"
echo -e "\033[32m ##                                                                     ##\033[0m"
echo -e "\033[32m ##                         1. 开始安装程序                             ##\033[0m"
echo -e "\033[32m ##                         2. 卸载程序                                 ##\033[0m"
echo -e "\033[32m ##                         3. 清除docker残余镜像                       ##\033[0m"
echo -e "\033[32m ##                         4. 退出命令                                 ##\033[0m"
echo -e "\033[32m ##                                                                     ##\033[0m"
echo -e "\033[32m #########################################################################\033[0m"
while true
do
    read -p "请输入安装序号:" num
    if [ $num ==  "1" ]
    then
        install
        # 开始安装
        echo -e "\033[32m 启动docker开始部署服务...\033[0m"
        docker-compose up -d
        if [ $? -eq  0 ]; then
            echo -e "\033[32m 恭喜你安装成功!!!\033[0m"
        else
            echo -e "\033[31m 安装失败!\033[0m"
        fi
        break
    elif [ $num ==  "2" ]
    then
       cd / && rm -rf /pinkmoe && docker stop pinkmoe_mysql pinkmoe_redis pinkmoe_server pinkmoe_index_admin && docker rm pinkmoe_mysql pinkmoe_redis pinkmoe_server pinkmoe_index_admin && docker rmi pinkmoe_server pinkmoe_index_admin && docker rmi daff57b7d2d1 1319b1eaa0b7 && docker system prune
        break
    elif [ $num ==  "3" ]
    then
        docker system prune
        break
    elif [ $num ==  "4" ]
    then
        break
    else
        echo -e "\033[31m 请输入正确的序号\033[0m"
    fi
done