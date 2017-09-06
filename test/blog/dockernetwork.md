#Docker Network

##1. Docker 的4种网络模式

执行docker run命令时，可以用--net选项指定容器的网络模式，Docker有以下4中网络模式：

- host模式： 容器和宿主机共享Network namespace
- container模式： 容器和另外一个容器共享Network namespace
- none模式： 容器有独立的Network namespace，但并没有对其进行任何网络设置，如分配veth pair 和网桥连接，配置IP等
- bridge模式： Bridge模式是Docker的默认模式


### 1.1	host模式
Docker使用了Linux的Namespaces技术来进行资源隔离，如PID Namespace隔离进程，Mount Namespace隔离文件系统，Network Namespace隔离网络等。一个Network Namespace提供了一份独立的网络环境，包括网卡、路由、Iptable规则等都与其他的Network Namespace隔离。一个Docker容器一般会分配一个独立的Network Namespace。但如果启动容器的时候使用host模式，那么这个容器将不会获得一个独立的Network Namespace，而是和宿主机共用一个Network Namespace。容器将不会虚拟出自己的网卡，配置自己的IP等，而是使用宿主机的IP和端口。但是，容器的其他方面，如文件系统、进程列表等还是和宿主机隔离的。

### 1.2	container模式
container模式指定新创建的容器和已经存在的一个容器共享一个Network Namespace，而不是和宿主机共享。新创建的容器不会创建网卡，配置IP，而是和一个指定的容器共享IP、端口范围等。同样，两个容器除了网络方面，其他的如文件系统、进程列表等还是隔离的。共享Network Namespace的两个容器可以通过lo设备设备通信。

### 1.3	none模式
none模式下，容器拥有独立的Network Namespace，但是不会为容器配置网络，也就是说，这个容器没有网卡、IP、路由等信息。
### 1.4	bridge模式
bridge模式是Docker默认的网络设置，此模式会为每一个容器分配Network Namespace、设置IP等，并将一个主机上的Docker容器连接到一个虚拟网桥上。
## 2 创建容器
### 2.1	host模式
host模式下容器使用宿主机的网络环境，但这并不表示容器拥有宿主机的文件系统，所以Docker deamo会将标识网络信息的文件从宿主机添加到容器里面。
主要有以下两个文件：
宿主机的主机名（hostname）；
宿主机上/etc/hosts文件，用于配置IP地址以及主机名。
其中，宿主机的主机名hostname用于创建container.Config中的Hostname与Domainname属性。

Docker Daemon在容器rootfs内创建hostname文件，并在文件中写入Hostname与Domainname；然后创建hosts文件，并写入宿主机上/etc/hosts内的所有内容。

Hostname文件中的内容为容器ID（下同）。

对应容器文件路径为：/var/lib/docker/containers/containerID/hosts(hostname)
### 2.2	container模式
1.	这里假设将容器c2 链接到容器nc
2.	根据HostConfig.NetworkMode字段找到nc的containerID，如果nc.containerID等于c2.containerID或nc不处于running状态则返回空
3.	将nc容器对象的HostsPath、ResolveConfPath、Hostname和Domainname赋值给当前容器对象container，实现代码如下：
			container.HostsPath = nc.HostsPath
			container.ResolvConfPath = nc.ResolvConfPath
			container.Config.Hostname = nc.Config.Hostname
			container.Config.Domainname = nc.Config.Domainname

### 2.3	none模式
1.	设置container.Config.NetworkDisabled为true
2.	创建hosts文件，内容为“127.0.0.1”，容器只能使用本机网络


### 2.4	bridge模式
1.	通过allocateNetwork函数为容器分配网络接口设备需要的资源信息（包括IP、bridge、Gateway等），并赋值给container对象的NetworkSettings
2.	为容器创建hostname以及创建Hosts等文件

