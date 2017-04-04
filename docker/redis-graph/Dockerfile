FROM bazuka/redis-unstable

MAINTAINER Alex Mushket <uzibazuka@gmail.com>

EXPOSE 6379

WORKDIR /redis-graph

RUN mkdir -p module

COPY libmodule.so /redis-graph/module/
COPY ./docker/redis-unstable/docker-entrypoint.sh .

RUN chown -R redis:redis /redis-graph

ENTRYPOINT ["./docker-entrypoint.sh"]
CMD [ "redis-server", "--loadmodule /redis-graph/module/libmodule.so" , "--appendonly yes", "--appendfilename redis-graph.aof", "--dir /data"]