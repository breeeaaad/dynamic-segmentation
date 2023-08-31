build:
    chmod +x ./envs.sh && ./envs.sh
    docker-compose up --build

run: build
    docker-compose up --remove-orphans

stop: 
    docker-compose down --remove-orphans