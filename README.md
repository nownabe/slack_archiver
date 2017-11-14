Slack Archiver
==============

# Get Slack API Token
* Go to https://api.slack.com/apps and create new app.
* In Permissions functionality, select scopes and save them.
  * channels:history
  * channels:read
  * channels:write
* Click "Install App to Workspace" and then get OAuth Access Token.

# Development
```bash
make run
```

# Deploy
```bash
cp app.yaml.example app.yaml
vi app.yaml
cp cron.yaml.example cron.yaml
vi cron.yaml
go get google.golang.org/appengine
gcloud app deploy app.yaml cron.yaml
```
