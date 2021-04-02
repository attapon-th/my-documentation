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
```
.PHONY all

PUB_ATTAPON="ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQD81dBPJDLvfimAjpKNlEvCp/O4CRDV072qcMcttFV8p4mV1aD9nXuyoVF4yFb5AFJ/Y8L4mzuiBpiEcLfnhGp2Z8Fox7Z9b3LfXHHdSkKY/EgrwS/6QXx/m2ItHO6E4o0rkCnwpjwy/YEZCZTIJPF7Gunt25IMLruAyCTnpAinFlGGXSn9w3zusuMQYrN7eiUILuZCZU+Ty4s+BF+7Sh2hy26TpiNZSOkCq7/L/xzfuhRdz2fp9a3uUp+CsOTpkgk/Agpg2MoJKM9xUcipETnUKFlzUTplyW+IL2E9IpR9CRbwKN3dkXcmMfTxDzyDb8EPnVIGeP/dHDZ88rrR+mv1"
PUB_JETSADA="ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQC33mYYArt5wyVD38Zmy+Tt4q0FHYNpc9un5rXeRJebV8WJCYUldONhBssF+IAIQMz9/Oa45bepHhMdHT/ZyRly4xAwLVGwMmRDbRbY+FqeAkyflWQ/U0btzdVXq66PWMElAe0664ZBaMNo/xTqinMr5sObgEX8D8DuW9BRsR15rK4IJAG8yKfFZwko4IuxadfcpiUsOzMFPcWr+MLE2YCCqeW+ey/37XeuKT7mzySnkjq5gzme2JAlBswNGX2rRbGHSp+oYIlCDIC++Ikv/uJgtL09mpbp1QYg0nmTE1Np/CTb8T6l4WX4aI+klNdyf+HjM9JFVxCfXBx/h73lRi33 ict_moph@DESKTOP-DSSADHV"
PUB_nattawuth="ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDS+wSv7xvQea8ZFTnnFU+e1yQRTrQAmsYfejgE9n1xPsjtG+kIzu07IXZaZXbPzdCZfBYMycC/8Ybwf6oxt6Rhu3jfOUY2TWjC0By9kqdPzEW/mdbRz3BWTnL6rEhM4sWz5AIlgSuqJ1Cx4KiPnthLJ+WrQ0E/j6ztZMkJA9QMwvAcJIvbDeUa/33VfNUbqXCnTfV+TB1ADLIJRWcwb1oB7E20umj9cqbPjmFJChu4qU1Q6cy1DifJ47/+DMEwkLHEromIq54e2Ibaxsb9XR+hPXP+sttNTKihuHluV/waP4hK3hT2fWtS0P5wVxqsFfOdsFHDDy7MRpk7p7r2cuUt master@DESKTOP-5T8DDCD"

attapon:
	adduser attapon
	usermod -a -G wheel attapon
	usermod -a -G docker attapon
	su - attapon
	echo ${PUB_ATTAPON} >> ~/.ssh/authorized_keys
	chmod -R go= ~/.ssh
	chmod 600 ~/.ssh/authorized_keys

jetsada:
	adduser jetsada
	usermod -a -G wheel jetsada
	usermod -a -G docker jetsada
	mkdir -p /home/jetsada/.ssh
	echo ${PUB_JETSADA} >> /home/jetsada/.ssh/authorized_keys
	chown -R jetsada:jetsada /home/jetsada/.ssh
	chmod -R go= /home/jetsada/.ssh
	chmod 600 /home/jetsada/.ssh/authorized_keys

nattawuth:
	adduser nattawuth
	usermod -a -G wheel nattawuth
	usermod -a -G docker nattawuth
	mkdir -p /home/nattawuth/.ssh
	echo ${PUB_nattawuth} >> /home/nattawuth/.ssh/authorized_keys
	chown -R nattawuth:nattawuth /home/nattawuth/.ssh
	chmod -R go= /home/nattawuth/.ssh
	chmod 600 /home/nattawuth/.ssh/authorized_keys
```

> Written with [StackEdit](https://stackedit.io/).
<!--stackedit_data:
eyJoaXN0b3J5IjpbLTE5MjkwODI0NTcsMjg0MTgyNzIzLDEyMT
cxNTYxMDQsNjgxNTk3ODE2XX0=
-->
