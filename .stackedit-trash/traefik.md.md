
# Traefik Configuration

## logrotate
	sudo vi /etc/logrotate.d/traefik

Config:
```
/var/log/traefik/*.log
{
    compress
    create 0640 root root
    daily
    delaycompress
    missingok
    notifempty
    rotate 90

    postrotate
        kill -USR1 `pgrep traefik`
    endscript
}
```
	sudo logrotate /etc/logrotate.d/traefik --debug
	
	sudo logrotate /etc/logrotate.d/traefik




> Written with [StackEdit](https://stackedit.io/).
<!--stackedit_data:
eyJoaXN0b3J5IjpbMTIxMDQzNDE2OF19
-->