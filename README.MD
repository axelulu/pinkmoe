#### 演示站点:https://pinkmoe.com

#### 连接linux服务器终端（本教程暂时适用于centos7机器，其他机器没有测试过）

使用sftp这类软件连接到服务器终端后更新软件源

`sudo yum update -y`

##### 安装docker

`curl -fsSL https://get.docker.com | bash -s docker --mirror Aliyun`

#### 设置开机自启

`systemctl enable docker`

#### 启动docker

`systemctl start docker`

#### 安装docker-compose

`curl -L "https://github.com/docker/compose/releases/download/v2.7.0/docker-compose-$(uname -s)-$(uname -m)" > /usr/bin/docker-compose`

`chmod +x /usr/bin/docker-compose`

#### 下载本程序到您服务器的root目录中

1. 直接在您网站运行`yum install git -y && git clone https://github.com/Coder-ZhaoLu/pinkmoe && cd pinkmoe/ && cp ./pinkmoe_server/config.yaml.example ./pinkmoe_server/config.yaml && ls`（此方法克隆到您网站之后需要使用sftp等软件进行下面步骤的文件修改）

2. 直接下载文件到您本地，然后进行下面步骤文件修改后上传到您网站服务器

#### 修改网站前端配置文件

编辑首页文件：/pinkmoe_index/nginx/pinkmoe_ssl.conf和/pinkmoe_index/nginx/pinkmoe.conf

编辑后台文件：/pinkmoe_admin/nginx/admin_pinkmoe_ssl.conf和/pinkmoe_admin/nginx/admin_pinkmoe_ssl.conf

修改server_name后面的域名为你自己的域名，每个域名以英文空格隔开

#### 修改网站后端配置文件

复制/pinkmoe_server/config.yaml.example为config.yaml到/pinkmoe_server/目录中，按需编辑config.yaml中的配置为您网站的后台配置

#### 替换ssl文件为你自己的文件

替换/pinkmoe_index/nginx/ssl与/pinkmoe_admin/nginx/ssl中的文件为你自己网站的ssl文件（注意：替换的文件名称一定要与ssl文件夹中的名称一致）

#### 启动docker服务进行安装

`docker-compose up -d`

如果中途没有报错即表示安装成功，访问域名即可

#### 常见问题:

1. 报内存不足错误, 运行`docker system prune` 删除那些已停止的容器、dangling 镜像、未被容器引用的 network 和构建过程中的 cache

##### （如有报错请加QQ群进行反馈：829704922）

管理员账户:izhaicy@163.com

管理员密码:admin123
