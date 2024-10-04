# Build & Push Docker image

## Fetcher
cd backend/fetcher

docker build -t rsiegfanz\home-control-fetcher:latest -f backend\fetcher\Dockerfile .

docker push rsiegfanz/home-control-fetcher:latest

## ingestor

## seeder

## server (including frontend)
