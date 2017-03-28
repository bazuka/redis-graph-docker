FROM rikorose/gcc-cmake

RUN apt-get update -y && apt-get install uuid-dev

RUN mkdir /builder
WORKDIR /builder

RUN git clone https://github.com/swilly22/redis-module-graph.git

RUN cd redis-module-graph &&\
    cmake . && make module

VOLUME $PWD:/builder/dest
# RUN cp -R ./redis-module-graph/src /builder/dest/