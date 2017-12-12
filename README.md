# ChannelX-Project

Run following the first time (It will log the aggregate output to the terminal):

```
docker-compose up
```

For the following runs:

```
docker-compose start && docker-compose logs -ft
```

It will start the services and will start printing aggregate outputs to the same terminal. Ctrl+C will exit from the log tail but will NOT stop any services.

By default, dev versions of the backend and frontend will be deployed on watch mode with volumes attached to local files: Whenever you make any changes, both Angular and Go files will be recompiled and server will be restarted automatically. However, anything else won't be run (any sql scripts or anything), so run them independently.

If you need to access any of the services, run the following while the services are running:

```
docker-compose exec <service_name: backend | mysql> bash
```

After you are done with everything, stop the services with following:

```
docker-compose stop
```
