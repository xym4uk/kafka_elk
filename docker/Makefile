up:
	@docker-compose up -d

down:
	@docker-compose down

producer:
	@docker run -it --network docker_elk docker_app /app/producer $(count)

consumer:
	@docker run -it --network docker_elk docker_app /app/consumer