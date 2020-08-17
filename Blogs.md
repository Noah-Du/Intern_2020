## Configure Registry Mirrors VMware Fusion Project Nautilus

### Why You Need Registry Mirrors

With Project Nautilus, we can pull images from remote container repositaries Docker Hub. 
Using command `vctl pull`, you pull the images you want from `index.docker.io/library/`.
But if you are in some countries with Firewall, you pulling speed could be limited.
Then you would need a local registry mirror to get the pulling command faster.

### How to Configure Registry Mirrors

If you have tried Docker to configure a registry mirror, you could know that you can edit the setting from Docker Desktop or edit /etc/docker/daemon.json and add the registry-mirrors key and value, to make the change persistent.

    {
      "registry-mirrors": ["https://<my-docker-mirror-host>"]
    }
    
But we do not have the configuration file like dockerd to set a registry mirror with Project Nautilus. Then how can we do it?
There are two ways to use it easily:

#### 1. Add Resgistry Mirrors Url In Pulling Command

When you use "vctl pull" command, you can add the registry url before the images' names. Here is the example:

Using this way, we can pull the images from local mirrors. But this method cannot set the registry mirrors permanently. You need to add the registry mirrors' url every time you pull.

#### 2. Using the Mirrors.go Plug-in

You can download the filr mirrors.go from this github and use it to set the registry mirrors permanently. After you download it, you can compile it with command:

    > go build mirrors.go
    > ls
      mirrors		mirrors.go
      
Then you would get a executable file called `mirrors`. And then you can run it with images you want to pull or change the registry mirrors.
One thing to notice that you need to add `-i` before the images you would like to pull and `-a` before the registry mirror's url you would like to change.
Default setting of registry mirrors is 163yun with url https://hub-mirror.c.163.com.

![images](https://github.com/Noah-Du/Intern_2020/blob/master/source/mirrors%20example.png)

But you need to remember that your registry mirrors url must end with repositories directory like `/library/`.

### Registry Mirrors Url for Chinese Users

Here are some registry mirrors you can use with the addresses.

| Registry Mirrors | Addresses                               | Exclusive or Not                |
| ---------------- | --------------------------------------- | ------------------------------- |
| [USTC](https://mirrors.ustc.edu.cn/help/dockerhub.html)             | https://docker.mirrors.ustc.edu.cn      |                                 |
| [Alibaba Cloud](https://cr.console.aliyun.com/)    | https://<your_code>.mirror.aliyuncs.com | Need to sign in and get the url |
| [163yun](https://c.163yun.com/hub)           | https://hub-mirror.c.163.com            |                                 |
