chcp 65001
cd project-user
docker build -f Dockerfile -t project-user:latest .
docker build -f Dockerfile-2 -t project-user-2:latest .
cd ..
cd project-project
docker build -f Dockerfile -t project-project:latest .
docker build -f Dockerfile-2 -t project-project-2:latest .
cd ..
cd project-api
docker build -t project-api:latest .
docker-compose up -d