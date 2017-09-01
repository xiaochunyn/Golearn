## Linux 网卡bond的七种模式

###什么是bond
   网卡bond是通过多张网卡绑定为一个逻辑网卡，实现本地网卡的冗余，带宽扩容和负载均衡，在生产场景中是一种常用的技术。Kernels 2.4.12及以后的版本均供bonding模块，以前的版本可以通过patch实现。可以通过以下命令确定内核是否支持 bonding：
```
#cat /boot/config-2.6.32-573.el6.x86_64 |grep -i bonding
CONFIG_BONDING=m
```

## bond的七种模式介绍：
### 1、mode=0(balance-rr)(平衡抡循环策略)
链路负载均衡，增加带宽，支持容错，一条链路故障会自动切换正常链路。交换机需要配置聚合口，思科叫port channel。
- 特点：传输数据包顺序是依次传输（即：第1个包走eth0，下一个包就走eth1….一直循环下去，直到最后一个传输完毕），此模式提供负载平衡和容错能力；但是我们知道如果一个连接或者会话的数据包从不同的接口发出的话，中途再经过不同的链路，在客户端很有可能会出现数据包无序到达的问题，而无序到达的数据包需要重新要求被发送，这样网络的吞吐量就会下降
- 表示负载分担round-robin，并且是轮询的方式比如第一个包走eth0，第二个包走eth1，直到数据包发送完毕。
 -   优点：流量提高一倍
 -   缺点：需要接入交换机做端口聚合，否则可能无法使用

### 2、mode=1(active-backup)(主-备份策略)
这个是主备模式，只有一块网卡是active，另一块是备用的standby，所有流量都在active链路上处理，交换机配置的是捆绑的话将不能工作，因为交换机往两块网卡发包，有一半包是丢弃的。
- 特点：只有一个设备处于活动状态，当一个宕掉另一个马上由备份转换为主设备。mac地址是外部可见得，从外面看来，bond的MAC地址是唯一的，以避免switch(交换机)发生混乱。
此模式只提供了容错能力；由此可见此算法的优点是可以提供高网络连接的可用性，但是它的资源利用率较低，只有一个接口处于工作状态，在有 N 个网络接口的情况下，资源利用率为1/N
- 优点：冗余性高
- 缺点：链路利用率低，两块网卡只有1块在工作

### 3、mode=2(balance-xor)(平衡策略)
表示XOR Hash负载分担，和交换机的聚合强制不协商方式配合。（需要xmit_hash_policy，需要交换机配置port channel）
- 特点：基于指定的传输HASH策略传输数据包。缺省的策略是：(源MAC地址 XOR 目标MAC地址) % slave数量。其他的传输策略可以通过xmit_hash_policy选项指定，此模式提供负载平衡和容错能力

### 4、mode=3(broadcast)(广播策略)
表示所有包从所有网络接口发出，这个不均衡，只有冗余机制，但过于浪费资源。此模式适用于金融行业，因为他们需要高可靠性的网络，不允许出现任何问题。需要和交换机的聚合强制不协商方式配合。
- 特点：在每个slave接口上传输每个数据包，此模式提供了容错能力
- 必要条件：
        条件1：ethtool支持获取每个slave的速率和双工设定
        条件2：switch(交换机)支持IEEE802.3ad Dynamic link aggregation
        条件3：大多数switch(交换机)需要经过特定配置才能支持802.3ad模式

### 5、mode=4(802.3ad)(IEEE 802.3ad 动态链接聚合)
表示支持802.3ad协议，和交换机的聚合LACP方式配合（需要xmit_hash_policy）.标准要求所有设备在聚合操作时，要在同样的速率和双工模式，而且，和除了balance-rr模式外的其它bonding负载均衡模式一样，任何连接都不能使用多于一个接口的带宽。
- 特点：创建一个聚合组，它们共享同样的速率和双工设定。根据802.3ad规范将多个slave工作在同一个激活的聚合体下。
外出流量的slave选举是基于传输hash策略，该策略可以通过xmit_hash_policy选项从缺省的XOR策略改变到其他策略。需要注意的 是，并不是所有的传输策略都是802.3ad适应的，
尤其考虑到在802.3ad标准43.2.4章节提及的包乱序问题。不同的实现可能会有不同的适应 性。
- 必要条件：
条件1：ethtool支持获取每个slave的速率和双工设定
条件2：switch(交换机)支持IEEE 802.3ad Dynamic link aggregation
条件3：大多数switch(交换机)需要经过特定配置才能支持802.3ad模式

### 6、mode=5(balance-tlb)(适配器传输负载均衡)
是根据每个slave的负载情况选择slave进行发送，接收时使用当前轮到的slave。该模式要求slave接口的网络设备驱动有某种ethtool支持；而且ARP监控不可用。
- 特点：不需要任何特别的switch(交换机)支持的通道bonding。在每个slave上根据当前的负载（根据速度计算）分配外出流量。如果正在接受数据的slave出故障了，另一个slave接管失败的slave的MAC地址。
- 必要条件：
ethtool支持获取每个slave的速率

### 7、mode=6(balance-alb)(适配器适应性负载均衡)
在5的tlb基础上增加了rlb(接收负载均衡receive load balance).不需要任何switch(交换机)的支持。接收负载均衡是通过ARP协商实现的.
- 特点：该模式包含了balance-tlb模式，同时加上针对IPV4流量的接收负载均衡(receive load balance, rlb)，而且不需要任何switch(交换机)的支持。接收负载均衡是通过ARP协商实现的。bonding驱动截获本机发送的ARP应答，并把源硬件地址改写为bond中某个slave的唯一硬件地址，从而使得不同的对端使用不同的硬件地址进行通信。
来自服务器端的接收流量也会被均衡。当本机发送ARP请求时，bonding驱动把对端的IP信息从ARP包中复制并保存下来。当ARP应答从对端到达 时，bonding驱动把它的硬件地址提取出来，并发起一个ARP应答给bond中的某个slave。
使用ARP协商进行负载均衡的一个问题是：每次广播 ARP请求时都会使用bond的硬件地址，因此对端学习到这个硬件地址后，接收流量将会全部流向当前的slave。这个问题可以通过给所有的对端发送更新 （ARP应答）来解决，应答中包含他们独一无二的硬件地址，从而导致流量重新分布。
当新的slave加入到bond中时，或者某个未激活的slave重新 激活时，接收流量也要重新分布。接收的负载被顺序地分布（round robin）在bond中最高速的slave上
当某个链路被重新接上，或者一个新的slave加入到bond中，接收流量在所有当前激活的slave中全部重新分配，通过使用指定的MAC地址给每个 client发起ARP应答。下面介绍的updelay参数必须被设置为某个大于等于switch(交换机)转发延时的值，从而保证发往对端的ARP应答 不会被switch(交换机)阻截。
- 必要条件：
条件1：ethtool支持获取每个slave的速率；
条件2：底层驱动支持设置某个设备的硬件地址，从而使得总是有个slave(curr_active_slave)使用bond的硬件地址，同时保证每个bond 中的slave都有一个唯一的硬件地址。如果curr_active_slave出故障，它的硬件地址将会被新选出来的 curr_active_slave接管
其实mod=6与mod=0的区别：mod=6，先把eth0流量占满，再占eth1，….ethX；而mod=0的话，会发现2个口的流量都很稳定，基本一样的带宽。而mod=6，会发现第一个口流量很高，第2个口只占了小部分流量。
 
### 小结
mode5和mode6不需要交换机端的设置，网卡能自动聚合。mode4需要支持802.3ad。mode0，mode2和mode3理论上需要静态聚合方式。
但实测中mode0可以通过mac地址欺骗的方式在交换机不设置的情况下不太均衡地进行接收。

## 二、bond的配置实例
1、首先要看linux是否支持bonding,大部分发行版都支持
```
# modinfo bonding |more
filename:       /lib/modules/2.6.32-431.el6.x86_64/kernel/drivers/net/bonding/bonding.ko
author:         Thomas Davis, tadavis@lbl.gov and many others
description:    Ethernet Channel Bonding Driver, v3.6.0
version:        3.6.0
license:        GPL
srcversion:     353B1DC123506708446C57B
depends:        8021q,ipv6
vermagic:       2.6.32-431.el6.x86_64 SMP mod_unload modversions
```
如输出以上信息，则说明支持bonding，如果没有,说明内核不支持bonding,需要重新编译内核
2、网卡配置文件
两个物理网口分别是：eth0,eth1 绑定后的虚拟口是：bond0
```
[root@jacken ~]# cat /etc/sysconfig/network-scripts/ifcfg-eth0 
DEVICE=eth0
HWADDR=EC:F4:BB:DC:4C:0C
TYPE=Ethernet
UUID=669f0694-9c52-4792-bd67-22c9d2c17acb
ONBOOT=yes
NM_CONTROLLED=no
BOOTPROTO=none
MASTER=bond0
SLAVE=yes
[root@jacken ~]# cat /etc/sysconfig/network-scripts/ifcfg-eth1
DEVICE=eth1
HWADDR=EC:F4:BB:DC:4C:0D
TYPE=Ethernet
UUID=1d2f30f4-b3f0-41a6-8c37-54f03115f7bd
ONBOOT=yes
NM_CONTROLLED=no
BOOTPROTO=none
MASTER=bond0
SLAVE=yes
[root@jacken ~]# cat /etc/sysconfig/network-scripts/ifcfg-bond0
DEVICE=bond0
NAME='System bond0'
TYPE=Ethernet
NM_CONTROLLED=no
USERCTL=no
ONBOOT=yes
BOOTPROTO=none
IPADDR=192.168.1.100
NETMASK=255.255.255.0
BONDING_OPTS='mode=1 miimon=100'
IPV6INIT=no
```
开机自动加载模块到内核
```#echo 'alias bond0 bonding' >> /etc/modprobe.d/dist.conf
#echo 'options bonding mode=0 miimon=200' >> /etc/modprobe.d/dist.conf
#echo 'ifenslave bond0 eth0 eth1' >>/etc/rc.local
miimon=100
```
每100毫秒 (即0.1秒) 监测一次路连接状态，如果有一条线路不通就转入另一条线路； Linux的多网卡绑定功能使用的是内核中的"bonding"模块
如果修改为其它模式，只需要在BONDING_OPTS中指定mode=Number即可。USERCTL=no --是否允许非root用户控制该设备
查看bond0状态：可以看到调用的是哪几个物理网卡
```
#cat /proc/net/bonding/bond0

[root@compute05 ~]#  cat /proc/net/bonding/bond0
Ethernet Channel Bonding Driver: v3.7.1 (April 27, 2011)
Bonding Mode: fault-tolerance (active-backup)
Primary Slave: None
Currently Active Slave: eth1
MII Status: up
MII Polling Interval (ms): 100
Up Delay (ms): 0
Down Delay (ms): 0
Slave Interface: eth0
MII Status: up
Speed: 1000 Mbps
Duplex: full
Link Failure Count: 0
Permanent HW addr: ec:f4:bb:dc:4c:0c
Slave queue ID: 0
Slave Interface: eth1
MII Status: up
Speed: 1000 Mbps
Duplex: full
Link Failure Count: 0
Permanent HW addr: ec:f4:bb:dc:4c:0d
Slave queue ID: 0
```
三、扩展
上边是两个网卡(eth0、eth1)绑定成一个bond0，如果我们要设置多个bond口，比如物理网口eth0和eth1组成bond0，eth2和eth3组成bond1，那么网口设置文件的设置方法和上面
是一样的，只是/etc/modprobe.d/dist.conf文件就不能叠加了。正确的设置方法有两种:
1、第一种
```
alias bond0 bonding
alias bond1 bonding
options bonding max_bonds=2 miimon=200 mode=1
```
这样所有的绑定只能使用一个mode了。
2、第二种
```
alias bond0 bonding
options bond0 miimon=100 mode=1
install bond1 /sbin/modprobe bonding -o bond1 miimon=100 mode=0
install bond2 /sbin/modprobe bonding -o bond2 miimon=100 mode=1
install bond3 /sbin/modprobe bonding -o bond3 miimon=100 mode=0
```
这种方式不同的bond口可以设定为不同的mode,注意开机自动启动/etc/rc.d/rc.local文件的设置
```
ifenslave bond0 eth0 eth1
ifenslave bond1 eth2 eth3
ifenslave bond2 eth4 eth5
ifenslave bond3 eth6 eth7
```

###参考
http://lixin15.blog.51cto.com/3845983/1769338

http://linuxnote.blog.51cto.com/9876511/1680315
