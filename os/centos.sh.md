add_user()  { 
	USER=$1
	PASSWORD=$2
	echo "Adding user $USER ..." echo useradd -c $USER
	echo passwd $USER $PASSWORD
	echo "Added user $USER And pass $PASSWORD"  
}
add_group(){
	GROUP=$2
	USER=$1
	usermod -aG $GROUP $USERNAME
}
<!--stackedit_data:
eyJoaXN0b3J5IjpbOTAyNjMyMTA2XX0=
-->