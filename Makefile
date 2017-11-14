export GOROOT=$(dirname $(which dev_appserver.py))/../platform/google_appengine/goroot-1.8

.PHONY: run
run:
	export GOROOT=$(dirname $(which dev_appserver.py))/../platform/google_appengine/goroot-1.8
	dev_appserver.py app/app.yaml
