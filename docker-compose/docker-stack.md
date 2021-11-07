# Docker Stack
> For `docker-swarm`
> `traefik` - api gateway

```yaml
version: "3.3"
services:
  ${SERVICE}:
    image: ${IMAGE}:latest
    environment:
      - TZ=Asia/Bangkok
      - ENVELOPMENT=production
    deploy:
      mode: replicated
      replicas: 1      
      labels:
        - "traefik.enable=true"
        - "traefik.docker.network=traefik-public"
        - "traefik.http.routers.${SERVICE}.tls=true"
        - "traefik.http.routers.${SERVICE}.entrypoints=websecure"
        - "traefik.http.routers.${SERVICE}.rule=Host(`${DOMAIN}`) && PathPrefix(`${PREFIX}`)"
        - "traefik.http.services.${SERVICE}.loadbalancer.server.port=80"
    networks:
      - traefik-public
 
    
networks:
  traefik-public:
    external: true
```


> Written with [StackEdit](https://stackedit.io/).
<!--stackedit_data:
eyJoaXN0b3J5IjpbNDUzMzA3NzY1LC0zMTEzNTQyMzFdfQ==
-->