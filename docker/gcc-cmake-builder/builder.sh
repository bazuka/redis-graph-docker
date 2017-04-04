#!/bin/bash
git clone https://github.com/swilly22/redis-module-graph.git

cd redis-module-graph && cmake . && make module

cp /builder/redis-module-graph/src/libmodule.so /builder/host/libmodule.so