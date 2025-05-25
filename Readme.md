# Class Reminder Backend
This is the backend service for Class Reminder, built with Go and MySQL. It provides scheduled class notifications and can be deployed using Docker or run directly via Go.

# Installation
- install mysql
- install docker
- install rundeck
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

# Deployment
- ssh to your server, clone the repository
- setup secret and variables in CI/CD github actions
- create a new service in /etc/systemd/system
 ```service
    # /etc/systemd/system/class-reminder.service
    [Unit]
    Description=Class Reminder Backend Service
    After=network.target

    [Service]
    ExecStart=/home/your-user/class-reminder-be/class-reminder-be
    WorkingDirectory=/home/your-user/class-reminder-be
    Restart=always
    EnvironmentFile=/home/your-user/class-reminder-be/.env
    User=your-user

    [Install]
    WantedBy=multi-user.target
```