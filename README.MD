#### 演示站点:https://pinkmoe.com

#### 1. 连接linux服务器终端（本教程暂时适用于centos7机器，其他机器没有测试过）

使用sftp这类软件连接到服务器终端后更新软件源

`sudo yum update -y`

##### 2. 安装docker

`curl -fsSL https://get.docker.com | bash -s docker --mirror Aliyun`

#### 3. 设置开机自启

`systemctl enable docker`

#### 4. 启动docker

`systemctl start docker`

#### 5. 安装docker-compose

`curl -L "https://github.com/docker/compose/releases/download/v2.7.0/docker-compose-$(uname -s)-$(uname -m)" > /usr/bin/docker-compose`

`chmod +x /usr/bin/docker-compose`

#### 6. 下载本程序到您服务器的root目录中

1. 直接在您网站运行`yum install git -y && git clone https://github.com/Coder-ZhaoLu/pinkmoe && cd pinkmoe/ && ls`（此方法克隆到您网站之后需要使用sftp等软件进行下面步骤的文件修改）

2. 直接下载文件到您本地，然后进行下面步骤文件修改后上传到您网站服务器

#### 7. 修改网站前端配置文件

编辑首页文件：/pinkmoe_index/nginx/pinkmoe_ssl.conf和/pinkmoe_index/nginx/pinkmoe.conf

编辑后台文件：/pinkmoe_admin/nginx/admin_pinkmoe_ssl.conf和/pinkmoe_admin/nginx/admin_pinkmoe_ssl.conf

修改server_name后面的域名为你自己的域名，每个域名以英文空格隔开

#### 8. 修改网站后端配置文件

复制/pinkmoe_server/config.yaml.example为config.yaml到/pinkmoe_server/目录中，按需编辑config.yaml中的配置为您网站的后台配置

#### 9. 替换ssl文件为你自己的文件

替换/pinkmoe_index/nginx/ssl与/pinkmoe_admin/nginx/ssl中的文件为你自己网站的ssl文件（注意：替换的文件名称一定要与ssl文件夹中的名称一致）

#### 10. 启动docker服务进行安装

`docker-compose up -d`

如果中途没有报错即表示安装成功，访问域名即可

#### 常见问题:

1. 报内存不足错误, 运行`docker system prune` 删除那些已停止的容器、dangling 镜像、未被容器引用的 network 和构建过程中的 cache

##### （如有报错请加QQ群进行反馈：829704922）

管理员账户:izhaicy@163.com

管理员密码:Admin123

#### 网站图片预览：![image](https://github.com/Coder-ZhaoLu/pinkmoe/raw/pinkmoe-img/img1.png)![image](https://github.com/Coder-ZhaoLu/pinkmoe/raw/pinkmoe-img/img2.png)![image](https://github.com/Coder-ZhaoLu/pinkmoe/raw/pinkmoe-img/img3.png)![image](https://github.com/Coder-ZhaoLu/pinkmoe/raw/pinkmoe-img/img4.png)![image](https://github.com/Coder-ZhaoLu/pinkmoe/raw/pinkmoe-img/img5.png)![image](https://github.com/Coder-ZhaoLu/pinkmoe/raw/pinkmoe-img/img6.png)![image](https://github.com/Coder-ZhaoLu/pinkmoe/raw/pinkmoe-img/img7.png)![image](https://github.com/Coder-ZhaoLu/pinkmoe/raw/pinkmoe-img/img8.png)![image](https://github.com/Coder-ZhaoLu/pinkmoe/raw/pinkmoe-img/img9.png)![image](https://github.com/Coder-ZhaoLu/pinkmoe/raw/pinkmoe-img/img10.png)![image](https://github.com/Coder-ZhaoLu/pinkmoe/raw/pinkmoe-img/img11.png)![image](https://github.com/Coder-ZhaoLu/pinkmoe/raw/pinkmoe-img/img12.png)![image](https://github.com/Coder-ZhaoLu/pinkmoe/raw/pinkmoe-img/img13.png)![image](https://github.com/Coder-ZhaoLu/pinkmoe/raw/pinkmoe-img/img14.png)![image](https://github.com/Coder-ZhaoLu/pinkmoe/raw/pinkmoe-img/img15.png)![image](https://github.com/Coder-ZhaoLu/pinkmoe/raw/pinkmoe-img/img16.png)![image](https://github.com/Coder-ZhaoLu/pinkmoe/raw/pinkmoe-img/img17.png)![image](https://github.com/Coder-ZhaoLu/pinkmoe/raw/pinkmoe-img/img18.png)![image](https://github.com/Coder-ZhaoLu/pinkmoe/raw/pinkmoe-img/img19.png)![image](https://github.com/Coder-ZhaoLu/pinkmoe/raw/pinkmoe-img/img20.png)![image](https://github.com/Coder-ZhaoLu/pinkmoe/raw/pinkmoe-img/img21.png)![image](https://github.com/Coder-ZhaoLu/pinkmoe/raw/pinkmoe-img/img22.png)夜间模式（部分）：![image](https://github.com/Coder-ZhaoLu/pinkmoe/raw/pinkmoe-img/img23.png)![image](https://github.com/Coder-ZhaoLu/pinkmoe/raw/pinkmoe-img/img24.png)![image](https://github.com/Coder-ZhaoLu/pinkmoe/raw/pinkmoe-img/img25.png)![image](https://github.com/Coder-ZhaoLu/pinkmoe/raw/pinkmoe-img/img26.png)![image](https://github.com/Coder-ZhaoLu/pinkmoe/raw/pinkmoe-img/img27.png)![image](https://github.com/Coder-ZhaoLu/pinkmoe/raw/pinkmoe-img/img28.png)   后台演示图（部分）：![image](https://github.com/Coder-ZhaoLu/pinkmoe/raw/pinkmoe-img/img29.png)![image](https://github.com/Coder-ZhaoLu/pinkmoe/raw/pinkmoe-img/img30.png)![image](https://github.com/Coder-ZhaoLu/pinkmoe/raw/pinkmoe-img/img31.png)![image](https://github.com/Coder-ZhaoLu/pinkmoe/raw/pinkmoe-img/img32.png)![image](https://github.com/Coder-ZhaoLu/pinkmoe/raw/pinkmoe-img/img33.png)![image](https://github.com/Coder-ZhaoLu/pinkmoe/raw/pinkmoe-img/img34.png)
