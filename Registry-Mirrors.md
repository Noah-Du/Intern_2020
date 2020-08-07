## dockerhub镜像加速器

docker可以通过修改配置文件daemon.json，在里面添加{"registry-mirrors"}来配置私库镜像下载加速器，从而提高中国用户从dockerhub中pull镜像的速度。国内有不少机构提供了免费的加速器以方便大家使用，这里列出一些常用的加速器服务：

或者可以使用 --registry-mirror 配置 Docker 守护进程

可以配置 Docker 守护进程默认使用 Docker 官方镜像加速。这样可以默认通过官方镜像加速拉取镜像，而无需在每次拉取时指定 registry.docker-cn.com。

可以在 Docker 守护进程启动时传入 --registry-mirror 参数：

$ docker --registry-mirror=https://registry.docker-cn.com daemon


–registry-mirror

    镜像仓库会自动同步Docker Hub上的镜像到本地，在国内建立一个缓存，提交下载、上传速率

--------------

鉴于vctl在不指定镜像库的情况下也会从dockerhub中pull镜像，考虑是否可以为vctl添加命令来提高中国用户pull镜像的速度

Possible Solution:

- 和docker的操作一样，通过修改vctl配置文件添加registry-mirrors键值来实现

- 去实现一个新的小程序，比如vctlmirror，直接能通过os.exec调用vctl，把配好的proxy的image url发过去 (From Mentor)

----------------

如果我们在 docker 官方仓库拉取的镜像是以下形式：

docker pull xxx:yyy

那么使用 Azure 中国镜像，应该是这样拉取：

docker pull dockerhub.azk8s.cn/library/xxx:yyy

如果我们在 docker 官方仓库拉取的镜像是以下形式：

docker pull xxx/yyy:zz

那么使用 Azure 中国镜像，应该是这样拉取：

docker pull dockerhub.azk8s.cn/xxx/yyy:zz

下面以拉取 mysql:5.7 和 360cloud/wayne 为例，如下：

docker pull dockerhub.azk8s.cn/library/mysql:5.7

docker pull dockerhub.azk8s.cn/360cloud/wayne


拉完了重新 tag 为你要的镜像。
---------------

目前可用registry-mirror地址：

https://docker.mirrors.ustc.edu.cn/     [中科大](https://mirrors.ustc.edu.cn/help/dockerhub.html)

https://docker.mirrors.ustc.edu.cn/     网易云

实践发现不可用地址：

https://reg-mirror.qiniu.com/


