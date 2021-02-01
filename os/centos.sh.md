```shell
#!$SHELL
add_user()  { 
	USER=$1
	PASSWORD=$2
	echo "Adding user $USER ..." echo useradd -c $USER
	echo passwd $USER $PASSWORD
	echo "Added user $USER And pass $PASSWORD"  
}
add_group(){
	USER=$1
	GROUP=$2	
	usermod -aG $GROUP $USERNAME
}


add_key(){
	USER=$1
	KEY_PUB=$2	
	sudo mkdir -p /home/$USER/.ssh
	sudo echo "$KEY_PUB" >> /home/$USER/.ssh/authorized_keys
	sudo chown $USER:$USER
}


```
<!--stackedit_data:
eyJoaXN0b3J5IjpbLTEzNjIxMDYwMjJdfQ==
-->