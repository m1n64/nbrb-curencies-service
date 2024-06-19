## Go REST API NBRB currency rates project

****
Backend: Go 1.22, Gin-gonic, GORM, air and delve

Database: MySQL, Redis

Server, proxy: nginx

Documentation: swagger

*****
```shell
cp .env.example .env
```

Dev-mode:
```shell
docker-compose -f docker-compose.yml -f docker-compose.dev.yml up -d
```

Production:

```shell
docker-compose up -d
```

*****
API: [http://localhost/api/](http://localhost/api/)
Documentation: [http://localhost/swagger/](http://localhost/swagger/)
