
# Traefik Configuration

## logrotate
```
# /etc/logrotate.d/traefik

# to make log-rotate change take effect launch sudo logrotate /etc/logrotate.d/traefik
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

```shell
sudo vi /etc/logrotate.d/traefik
```


> Written with [StackEdit](https://stackedit.io/).
<!--stackedit_data:
eyJoaXN0b3J5IjpbLTQ0NDYwNzcxMV19
-->