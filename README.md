## Go REST API template project

****
Backend: Go 1.22, Gin-gonic, GORM, air and delve

Database: Postgresql, Redis

Server, proxy: nginx

*****

Dev-mode:
```shell
docker-compose -f docker-compose.yml -f docker-compose.dev.yml up -d
```

Production:

```shell
docker-compose up -d
```