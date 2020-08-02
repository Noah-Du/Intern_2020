This file is to record knowledge of docker I got through wesites.

# 1. What is Virtualization?

##  Concept
In computing, virtualization refers to the act of creating a virtual (rather than actual) version of something, including virtual computer hardware platforms, storage devices, and computer network resources.

Virtualization is a resource management technology that abstracts and transforms various physical resources of a computer, such as servers, networks, memory, and storage. Virtualization breaks the inseparable barriers between physical structures, allowing users to use these resources in a better way than the original configuration. The new virtual part of these resources is not restricted by the way existing resources are erected,and geographical or physical configuration. Generally referred to as virtualized resources include computer power and data storage.

In the actual production environment, virtualization technology is mainly used to solve the reorganization and reuse of high-performance physical hardware overcapacity and old hardware undercapacity so as to maximize the use of physical hardware. So the purpose of virtualization is to make full use of resources.

![image](https://github.com/Noah-Du/Intern_2020/blob/master/source/Evolution%20of%20Vitualization.png)


# 2. Docker

## 2.1 Container

### Virtual Machine vs. Container
Virtual machines and containers differ in several ways, but the primary difference is that containers provide a way to virtualize an OS so that multiple workloads can run on a single OS instance. With VMs, the hardware is being virtualized to run multiple OS instances. Containers’ speed, agility, and portability make them yet another tool to help streamline software development.

## 2.2 Docker's advantage compared with classical virtualization
1. The fast startup of docker belongs to the second level. The virtual machine usually takes a few minutes to start
2. Docker requires fewer resources. Docker is virtualized at the operating system level. The docker container interacts with the kernel. There is almost no performance loss. The performance is better than virtualization through the Hypervisor layer and the kernel layer.
3. Docker is lighter, and the docker architecture can share a core and shared application library, occupying very little memory. In the same hardware environment, Docker runs far more images than virtual machines, and the utilization of the system is very high.
4. Compared with virtual machines, docker has weaker isolation. docker belongs to the isolation between processes, and virtual machines can achieve system-level isolation
5. Security: The security of docker is also weaker. Docker's tenant root is equivalent to the host root. Once the user in the container is elevated from ordinary user privileges to root privileges, it will directly have the root privileges of the host and can perform unlimited operations. The root permissions of the virtual machine tenant and the root virtual machine permissions of the host machine are separated, and the virtual machine utilizes the ring-1 hardware isolation technology such as Intel's VT-d and VT-x. This isolation technology can prevent virtual machines from breaking through and each other Interaction, and the container has not yet had any form of hardware isolation, which makes the container vulnerable to attacks
6. Manageability: docker's centralized management tools are not yet mature. Various virtualization technologies have mature management tools. For example, VMware vCenter provides complete virtual machine management capabilities
7. High availability and recoverability: Docker's high availability support for business is achieved through rapid redeployment. Virtualization has mature guarantee mechanisms that have been tested in production practices such as load balancing, high availability, fault tolerance, migration, and data protection. VMware can promise 99.999% high availability of virtual machines to ensure business continuity
8. Fast creation and deletion: Virtualization creation is in minutes, Docker container creation is in seconds. Docker's rapid iterative nature determines that it can save a lot of time in development, testing, and deployment.
9. Delivery and deployment: Virtual machines can achieve consistency in environment delivery through mirroring, but mirroring distribution cannot be systematic. Docker records the container construction process in the Dockerfile, which can realize rapid distribution and rapid deployment in the cluster

| Feature          | Container   | Virtual Machine |
| ---------------- | ----------- | --------------- |
| Startup          | seconds     | minutes         |
| Hard Disks Usage | MB          | GB              |
| Performance      | close to OS | weaker than OS  |
| Systems support  | Thousands   | Tens            |

## 2.3 Images

### 2.3.1 What is images?

Mirror is a lightweight, executable independent software package. The image is used to package the software operating environment and the software developed based on the operating environment. It contains all the content needed to run a certain software, including code, libraries, environment variables and configuration files needed at runtime.

### 2.3.2 Why one image is so large?

UnionFS(联合文件系统):The Union file system is a hierarchical, lightweight and high-performance file system. He supports the modification of the file system as one submission to superimpose layer by layer, and at the same time, different directories can be mounted under the same virtual file system. The Union file system is the basis of Docker images. This file system feature: Load multiple file systems at the same time, but from the outside, only one file system can be seen. Joint loading will superimpose each layer of file system, so that the final file system will contain all the underlying files and directories.

### 2.3.3 Method of Image.

    The docker image is actually composed of a layered file system.
    
- Bootfs (boot file system) mainly includes bootloader and kernel. Bootloader is mainly to boot and load the kernel. The bootfs file system will be loaded when linux first starts. At the bottom of the docker image is bootfs. This layer is the same as the Linux/Unix system, including bootloader and kernel. When the boot is loaded, the entire kernel is stored in memory. At this time, the right to use memory has been transferred to the kernel by bootfs. Bootfs will be uninstalled at this time.

- Rootfs (root file system) contains standard directories and files such as /dev, /proc, /bin, /etc in a typical linux system. Rootfs is a variety of different operating system distributions, such as Ubuntu/CentOS and so on.

- The centos we usually install into the virtual machine are 1 to several GB. Why is docker only 200MB here? For a streamlined OS, rootfs can be very small, and it only needs to include the most basic commands, tools, and programs. Because the bottom layer directly uses the Host's kernel, you only need to provide rootfs. It can be seen that the different linux distributions have the same bootfs, but the rootfs will be different. Therefore, different distributions can share bootfs.
