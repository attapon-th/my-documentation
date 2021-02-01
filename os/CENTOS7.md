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
&& sudo yum install docker-ce docker-ce-cli containerd.io \
&& sudo systemctl start docker \
&& sudo systemctl enable docker \
&& sudo docker swarm init

```


> Written with [StackEdit](https://stackedit.io/).
<!--stackedit_data:
eyJoaXN0b3J5IjpbMjg0MTgyNzIzLDEyMTcxNTYxMDQsNjgxNT
k3ODE2XX0=
-->