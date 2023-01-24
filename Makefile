run-swagger:
	docker-compose -f "swagger-api/swagger.yml" up

run-redis:
	docker-compose -f "internal/repository/redis/docker-compose-only-redis.yml" up
