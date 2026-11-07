# sandbox-cardgame

Webanwendung für ein Sandbox-Kartenspiel bei dem Nutzer eigene Kartendecks hochladen und in Räumen nach eigenen Regeln spielen können

## Lokale Entwicklung

```shell
cd backend && go run .
```

## Docker

```shell
docker build backend -t cardgame-backend && docker run -p 5000:5000 cardgame-backend
```

## Formatting

```shell
cd backend && go fmt && cd ../frontend && npm run format && cd ..
```
