# Build & Push Docker image

## Fetcher
docker build -t rsiegfanz/home-control-fetcher:latest -f backend/fetcher/Dockerfile .

docker push rsiegfanz/home-control-fetcher:latest

## ingestor
docker build -t rsiegfanz/home-control-ingestor:latest -f backend/ingestor/Dockerfile .

docker push rsiegfanz/home-control-ingestor:latest

## seeder
docker build -t rsiegfanz/home-control-seeder:latest -f backend/seeder/Dockerfile .

docker push rsiegfanz/home-control-seeder:latest

## server (including frontend)
docker build -t rsiegfanz/home-control-server:latest -f backend/server/Dockerfile .

docker push rsiegfanz/home-control-server:latest
