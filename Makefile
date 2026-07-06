.PHONY: make\:migration

# Capture everything after the target name as an argument
ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
$(eval $(ARGS):;@:)

# Escape the colon with a backslash
make\:migration:
	@if [ -z "$(ARGS)" ]; then \
		echo "❌ Error: Please provide a migration name."; \
		echo "Usage: make make:migration <migration_name>"; \
		exit 1; \
	fi
	@go run cmd/migrate-make/main.go $(ARGS)