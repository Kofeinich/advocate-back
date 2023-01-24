run-swagger:
	docker-compose -f "swagger-api/swagger.yml" up

run-redis:
	docker-compose -f "redis.yml" up
