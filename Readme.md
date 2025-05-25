# Installation
- install mysql
- install docker
- execute the sql file in dump/db_class_reminder.sql

# Build docker image
```docker
docker build -t class-reminder-be .
```

# Run docker with docker compose
```docker
docker-compose up -d
```

# Build & run with golang directly (without docker)
- build: `go build`
- run: `./class-reminder-be`