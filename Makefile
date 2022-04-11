install:
	go install github.com/swaggo/swag/cmd/swag@latest

generate:
	swag init -ot go --parseDependency

docker-clean:
	docker system prune -a

docker-build:
	docker buildx build --platform linux/amd64 -t kedai-itemsvc .
	docker tag kedai-itemsvc alganbr/kedai-itemsvc
	docker push alganbr/kedai-itemsvc

docker-rebuild:
	make docker-clean
	make docker-build