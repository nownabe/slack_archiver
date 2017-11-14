export GOROOT=$(dirname $(which dev_appserver.py))/../platform/google_appengine/goroot-1.8

.PHONY: deploy
deploy:
	gcloud app deploy app/app.yaml

.PHONY: deploy-cron
deploy-cron:
	gcloud app deploy cron.yaml

.PHONY: run
run:
	export GOROOT=$(dirname $(which dev_appserver.py))/../platform/google_appengine/goroot-1.8
	dev_appserver.py app/app.yaml
