
# Traefik Configuration

## logrotate
```
# add logrotate config
# _> sudo vi /etc/logrotate.d/traefik
# to make log-rotate change take effect launch 
# sudo logrotate /etc/logrotate.d/traefik
# or reboot
/var/log/traefik.log
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




> Written with [StackEdit](https://stackedit.io/).
<!--stackedit_data:
eyJoaXN0b3J5IjpbLTUxODc0NjA0OF19
-->