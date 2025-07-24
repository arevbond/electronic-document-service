dev-run:
	docker compose -f compose.dev.yml up
dev-restart:
	docker compose -f compose.dev.yml restart app