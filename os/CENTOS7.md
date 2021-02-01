# Package `CentOS 7, 8`


## Initialization
[ShellScript `centos.sh`](./centos.sh.md)
```shell
./centos.sh add_user <user> <password>
./centos.sh add_group <user> <group>
./centos.sh add_sshkey <user> <pubilc_key>
```

## Recommend
```shell
sudo yum -y install update \
&& sudo yum -y install https://dl.fedoraproject.org/pub/epel/epel-release-latest-7.noarch.rpm \
&& sudo yum -y install \
	epel-release \
	htop \
	rsyne \
	jq	
```

## Docker
```shell
sudo yum install -y yum-utils \
&& sudo yum-config-manager \
    --add-repo \
    https://download.docker.com/linux/centos/docker-ce.repo \
&& 


```


> Written with [StackEdit](https://stackedit.io/).
<!--stackedit_data:
eyJoaXN0b3J5IjpbMjA0MzIyNjc5OCwxMjE3MTU2MTA0LDY4MT
U5NzgxNl19
-->