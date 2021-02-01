# Package `CentOS 7, 8`


## Initialization
```shell
add_user()  { 
	USER=$1
	PASSWORD=$2
	echo "Adding user $USER ..." echo useradd -c $USER
	echo passwd $USER $PASSWORD
	echo "Added user $USER And pass $PASSWORD"  
}
add_group(){
	GROUP=$2
	USERNAME=$1
	usermod -aG ${GROUP} ${}
}
```

## Recommend
```shell

```
## ssh

## htop

## docker

## Python

## rsync


> Written with [StackEdit](https://stackedit.io/).
<!--stackedit_data:
eyJoaXN0b3J5IjpbMTM1Njk0NjQxMyw2ODE1OTc4MTZdfQ==
-->