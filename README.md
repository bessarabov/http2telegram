# http2telegram

This is a simple software that can be used to debug HTTP requests.
It was created to simplify the debug of webhooks.

This is a web server that sends data about any HTTP request
that it get to telegram.

You can run this code on your own server, or deploy it to Google Cloud Run
(or to some other cloud provider that allow running docker containers).

You must specify env variables:

 * `TELEGRAM_TOKEN` — secret token for you telegram bot (you can get it from https://t.me/BotFather )
 * `CHAT_ID` — id of telegram chat where to send message
 * `PORT` — port that web server will use (this env variable is optional, the default is 8080)

```
docker run --rm -e TELEGRAM_TOKEN="bot123...:AAA..." -e CHAT_ID="12345..." -e PORT=8080 --publish=8080:8080 --name=http2telegram http2telegram
```
