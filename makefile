DEV_REQUIRED_EXECUTABLES := go docker air
CURRENT_SHELL := $(SHELL)

check_required_dev_executables:
	@ for exec in $(DEV_REQUIRED_EXECUTABLES); do \
		if ! which $$exec > /dev/null; then \
			echo "$$exec is not in PATH"; \
			exit 1; \
		fi; \
	done
	@ echo "All required executables ($(DEV_REQUIRED_EXECUTABLES)) are available in PATH."


# dev/local commands
dev_server:
	@ $(MAKE) check_required_dev_executables
	@ docker compose -f ./docker/server/docker-compose.yml up --build --remove-orphans

# migrations
m_up:
	@ docker compose -f ./docker/docker-compose-local.yml run --rm local_migrate \
	  -source file://migrations \
	  -database 'mysql://user:pass@tcp(local_mysql:3306)/db' \
	  -verbose up 1

m_status: 
	@ docker compose -f ./docker/docker-compose-local.yml run --rm local_migrate \
	  -source file://migrations \
	  -database 'mysql://user:pass@tcp(localhost:3306)/db' \
	  version

m_down: 
	@ docker compose -f ./docker/docker-compose-local.yml run --rm local_migrate \
	  -source file://migrations \
	  -database 'mysql://user:pass@tcp(localhost:3306)/db' \
	  -verbose down 1

m_create:
	@if [ -z "$(NAME)" ]; then \
		echo "Error: NAME is not set. Please provide a value like: make goose_create NAME=mig_xx1"; \
		exit 1; \
	fi
	@ docker compose -f ./docker/docker-compose-local.yml run --rm local_migrate \
	  create \
	  -dir file://migrations \
	  -ext sql \
	  -seq \
	  $(NAME)

# sqlc
sqlc_gen:
	@ docker compose -f ./docker/docker-compose-local.yml run --rm local_sqlc generate
	@ echo "sql schema generated"

sqlc_verify:
	@ docker compose -f ./docker/docker-compose-local.yml run --rm local_sqlc verify
	@ echo "sql schema verified"


# prod db tunnels
# @ssh -i <path_to_ssh_pem_files> -L <localhost:port>:<remote_localhost:port> root@<remote_ip>
tunnel_mongo:
	@ssh -i <path_to_ssh_pem_files> -L localhost:27018:localhost:27017 root@<remote_ip>

tunnel_mysql:
	@ssh -i <path_to_ssh_pem_files> -L localhost:3306:localhost:3306 root@<remote_ip>
