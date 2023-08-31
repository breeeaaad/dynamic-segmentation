export POSTGRES_PORT=5432
export POSTGRES_HOST=postgres
export POSTGRES_DB=avitointern
export POSTGRES_USER=postgres
export POSTGRES_PASSWORD=12345

build:
    docker-compose up --build

run: build
    docker-compose up --remove-orphans

stop: 
    docker-compose down --remove-orphans