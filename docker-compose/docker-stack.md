# Docker Stack
> For `docker-swarm`
> `traefik` - api gateway

```yaml
version: "3.3"
services:
  ${}:
    image: ${IMAGE}:latest
    environment:
      - TZ=Asia/Bangkok
      - "ENVELOPMENT=production"
      - "SERVER.LISTEN=0.0.0.0:80"
      - "SERVER.PREFIX=/dev-cobed"
      - "DB.COWARD.URI=${COWARD_DB_URI}"
      - "DB.CB.URI=${COBED_DB_URI}"
    deploy:
      mode: replicated
      replicas: 1      
      labels:
        - "traefik.enable=true"
        - "traefik.docker.network=traefik-public"
        - "traefik.http.routers.co_bed_api.tls=true"
        - "traefik.http.routers.co_bed_api.entrypoints=websecure"
        - "traefik.http.routers.co_bed_api.rule=Host(`indev.moph.go.th`) && PathPrefix(`/dev-cobed`)"
        - "traefik.http.services.co_bed_api.loadbalancer.server.port=80"
    networks:
      - traefik-public
  test_co_bed_frontend:
    image: registry.gitlab.com/indev-moph/co-bed-frontend:latest
    environment:
      - TZ=Asia/Bangkok
      - "ENVELOPMENT=production"
    networks:
      - traefik-public
    deploy:
      mode: replicated
      replicas: 1      
      labels:
        - "traefik.enable=true"
        - "traefik.docker.network=traefik-public"
        - "traefik.http.routers.test_co_bed_frontend.tls=true"
        - "traefik.http.routers.test_co_bed_frontend.entrypoints=websecure"
        - "traefik.http.routers.test_co_bed_frontend.rule=Host(`indev.moph.go.th`) && PathPrefix(`/co-bed-test`)"
        - "traefik.http.services.test_co_bed_frontend.loadbalancer.server.port=3000"
    
networks:
  traefik-public:
    external: true
```


> Written with [StackEdit](https://stackedit.io/).
<!--stackedit_data:
eyJoaXN0b3J5IjpbLTMxMTM1NDIzMV19
-->