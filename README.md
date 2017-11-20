ChannelX-Project

Run following:

```
docker build -t chx .
```

```
docker run -it -v $(pwd)/backend:/app/backend -v $(pwd)/frontend:/app/frontend -e MYSQL_DATABASE=ChannelX -e MYSQL_ROOT_PASSWORD=35792030 -e PKG_NAME=github.com/ChannelX-ITU/ChannelX-Project/backend --add-host="mysql:127.0.0.1" chx
```
