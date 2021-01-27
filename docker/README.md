# Docker  
  
  
## Traefik log 
>  Log Rotation by `logrotate` 

Check version:
```shell
logrotate --version
```
Create config log rotation:
```shell
vi /etc/logrotate.d/traefik
```

Config:
```raw
/var/log/traefik/*.log {
  daily
  rotate 30
  missingok
  notifempty
  compress
  dateext
  dateformat .%Y-%m-%d
  create 0644 root root
  postrotate
   docker kill --signal="USR1" <contrainer id or name>
  endscript
}
```

Run:
```shell
logrotate /etc/logrotate.conf --debug 
logrotate /etc/logrotate.conf
```
<!--stackedit_data:
eyJoaXN0b3J5IjpbLTgyMjA0MTQyMl19
-->