# Create Login with `rsa_key`

## URL: [How to set up ssh key](https://www.digitalocean.com/community/tutorials/how-to-set-up-ssh-keys-on-centos7)

## Generate SSH-KEY
1. Gen Key
```shell
$ ssh-keygen

Output
Generating public/private rsa key pair.
Enter file in which to save the key (/your_home/.ssh/id_rsa):
```
OR
```
ssh-keygen -t rsa -b 4096 -f ./id_rsa
```

## ***How to set permission

1. เพิ่ม id_rsa.pub ลงใน `~/.ssh/authorized_keys`
```shell
echo public_key_string >> ~/.ssh/authorized_keys

```
OR
```shell

cat ~/.ssh/id_rsa.pub | ssh username@remote_host "mkdir -p ~/.ssh && touch ~/.ssh/authorized_keys && chmod -R go= ~/.ssh && cat >> ~/.ssh/authorized_keys"
```
2. __***__ Add can login by key 
```
chmod -R go= ~/.ssh
```
3. เปลี่ยน [user] ให้ตรงกัน login 
```
chown -R [user]:[user] ~/.ssh
```
4. __optional__: เปลี่ยน `authorized_keys` เป็น `600` only
```
chmod 600 ~/.ssh/authorized_keys
```