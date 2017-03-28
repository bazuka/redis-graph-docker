FROM bazuka/redis-unstable

RUN apt-get update -y && apt-get install -y git uuid-dev build-essential

# RUN cd /tmp && wget https://cmake.org/files/v3.7/cmake-3.7.2.tar.gz &&\
#     tar xf cmake-3.7.2.tar.gz &&\
#     cd cmake-3.7.2 &&\
#     ./configure &&\
#     make &&\
#     make install

EXPOSE 6379

WORKDIR /redis-graph

RUN mkdir module
COPY ./redis-module-graph/src/libmodule.so module/

# RUN git clone https://github.com/swilly22/redis-module-graph.git
# RUN cd redis-module-graph && cmake . && make module
# VOLUME $PWD:/redis-graph

COPY ./docker/redis-unstable/docker-entrypoint.sh .
ENTRYPOINT ["./docker-entrypoint.sh"]

RUN chown -R redis:redis /redis-graph

CMD [ "redis-server", "--loadmodule /redis-graph/module/libmodule.so" , "--appendonly yes", "--appendfilename redis-graph.aof", "--dir /data"]
