DBNAME:=treasure_app
# https://docs.docker.com/docker-for-mac/networking/#use-cases-and-workarounds
DOCKER_DNS:=db
FLYWAY_CONF?=-url=jdbc:mysql://$(DOCKER_DNS):3306/$(DBNAME) -user=root -password=password
ZOOM_CLIENT_ID            :=
ZOOM_CLIENT_SECRET        :=
ZOOM_REDIRECT_URI         :=
PAYJP_PK_TEST             :=
PAYJP_SK_TEST             :=
SENDGRID_APIKEY           :=
SENDGRID_SERVICE_EMAIl    :=

export DATABASE_DATASOURCE:=root:password@tcp($(DOCKER_DNS):3306)/$(DBNAME)?parseTime=true&time_zone=%27Asia%2FTokyo%27&loc=Local
export GOOGLE_APPLICATION_CREDENTIALS=$(HOME)/.config/gcloud/treasure_app_service_account.json
export ZOOM_CLIENT_ID:=$(ZOOM_CLIENT_ID)
export ZOOM_CLIENT_SECRET:=$(ZOOM_CLIENT_SECRET)
export ZOOM_REDIRECT_URI:=$(ZOOM_REDIRECT_URI)
export PAYJP_PK_TEST:=$(PAYJP_PK_TEST)
export PAYJP_SK_TEST:=$(PAYJP_SK_TEST)
export SENDGRID_APIKEY:=$(SENDGRID_APIKEY)
export SENDGRID_SERVICE_EMAIL:=$(SENDGRID_SERVICE_EMAIL)

firebase.env:
	cp .firebase.env.example .firebase.env

zoom.env:
	cp .zoom.env.example .zoom.env

payjp.env:
	cp .payjp.env.example .payjp.env

sendgrid.env:
	cp .sendgrid.env.example .sendgrid.env

docker-compose/build:
	docker-compose build

docker-compose/up:
	docker-compose up

docker-compose/up/service:
	docker-compose up $(service)

docker-compose/down:
	docker-compose down

docker-compose/logs:
	docker-compose logs -f

frontend/deps:
	docker-compose exec frontend npm install

DB_SERVICE:=db
mysql/client:
	docker-compose exec $(DB_SERVICE) mysql -uroot -hlocalhost -ppassword $(DBNAME)

mysql/init:
	docker-compose exec $(DB_SERVICE) \
		mysql -u root -h localhost -ppassword \
		-e "create database \`$(DBNAME)\`"

__mysql/drop:
	docker-compose exec $(DB_SERVICE) \
		mysql -u root -h localhost -ppassword \
		-e "drop database \`$(DBNAME)\`"

MIGRATION_SERVICE:=migration
flyway/info:
	docker-compose run --rm $(MIGRATION_SERVICE) $(FLYWAY_CONF) info

flyway/validate:
	docker-compose run --rm $(MIGRATION_SERVICE) $(FLYWAY_CONF) validate

flyway/migrate:
	docker-compose run --rm $(MIGRATION_SERVICE) $(FLYWAY_CONF) migrate

flyway/repair:
	docker-compose run --rm $(MIGRATION_SERVICE) $(FLYWAY_CONF) repair

flyway/baseline:
	docker-compose run --rm $(MIGRATION_SERVICE) $(FLYWAY_CONF) baseline

# 以下 prod rds 用 flyway ( local で叩いても繋がらないよ
FLYWAY_DOCKER:=flyway/flyway
CODEBUILD_SRC_DIR?=$(shell pwd)
PROD_FLYWAY_CMD = \
  prod/flyway/info \
  prod/flyway/validate \
  prod/flyway/migrate \
  prod/flyway/repair \
  prod/flyway/baseline

$(PROD_FLYWAY_CMD):
	@echo run $(@F) target in prod
	@docker run -v $(CODEBUILD_SRC_DIR)/database/migration/schema:/flyway/sql -i --rm $(FLYWAY_DOCKER) $(FLYWAY_CONF) $(@F)
