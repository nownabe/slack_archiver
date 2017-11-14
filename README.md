Slack Archiver
==============

# Get Slack API Token
* Go to https://api.slack.com/apps and create new app.
* In Permissions functionality, select scopes and save them.
  * channels:history
  * channels:read
  * channels:write
* Click "Install App to Workspace" and then get OAuth Access Token.

# Edit YAMLs
```bash
cp app/app.yaml.example app/app.yaml
vi app/app.yaml
```

```bash
cp cron.yaml.example cron.yaml
vi cron.yaml
```

# Development
```bash
make run
```

# Deploy

App:

```bash
make deploy
```

Cron:

```bash
make deploy-cron
```
