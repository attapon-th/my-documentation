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


add_sshkey(){
	USER=$1
	KEY_PUB=$2	
	USER_HOME=/home/$USER
	sudo mkdir -p $USER_HOME/.ssh
	sudo echo "$KEY_PUB" >> $USER_HOME/.ssh/authorized_keys
	sudo chown -R $USER:$USER  $USER_HOME/.ssh
	sudo chmod -R go= $USER_HOME/.ssh
	sudo chmod 600 $USER_HOME/.ssh/authorized_keys
}


```
<!--stackedit_data:
eyJoaXN0b3J5IjpbMTg2OTEwNjA4NiwyNTY1NTU3MzVdfQ==
-->