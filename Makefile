run:
	@docker-compose -f config/compose.yml -p tweet up --build -d
stop:
	@docker-compose -f config/compose.yml -p tweet down --remove-orphans
