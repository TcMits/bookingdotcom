# bookingdotcom

# Quick start

run production server with seeding data (i assume that you have already installed docker):
```sh
docker compose -f production.yml up
```

run development server with seeding data (i assume that you have already installed docker):
```sh
docker compose -f local.yml up
```

fetch swagger docs (after starting server):
```sh
curl 0.0.0.0:8080/docs/index.html
```

use mongo express (after starting server):
```sh
curl 0.0.0.0:8081
```
