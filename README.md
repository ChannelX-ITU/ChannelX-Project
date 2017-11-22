# ChannelX-Project

Run following:

```
docker build -t chx .
```

For the first run:

```
docker run -it -v $(pwd)/backend:/app/backend -v $(pwd)/frontend:/app/frontend -v /app/frontend/node_modules -e MYSQL_DATABASE=ChannelX -e MYSQL_ROOT_PASSWORD=35792030 -e PKG_NAME=github.com/ChannelX-ITU/ChannelX-Project/backend -p 4200:4200 -p 6969:6969 --add-host="my.sql:127.0.0.1" --name chx chx
```

It will create a container named "chx". In order to keep using the same one after stopping the container use the following:


```
docker start chx
```
