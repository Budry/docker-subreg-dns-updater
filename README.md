# Docker Subreg DNS updater

## Docker

```shell script
docker run \
  -v /var/run/docker.sock:/var/run/docker.sock \
  -e SUBREG_PASSWORD=<password> \
  -e SUBREG_USER=<password> \ 
  budry/docker-subreg-dns-updater 
```

## Docker compose

```yaml
version: '3'

services:
    dns-updater:
      image: budry/docker-subreg-dns-updater
    vol

```