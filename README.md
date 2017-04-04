# Testing Redis graph module

[Roi Lipman](https://github.com/swilly22) created awesome lib for [Redis](https://github.com/antirez/redis) which using Hexastore concept to build relations between entities.

Get it from [Redis-module-graph](https://github.com/swilly22/redis-module-graph).

## Data and testing

First build all containers with `make build`.

To test redis-module-graph i've been using the open data from [Stanford Large Network Dataset Collection](https://snap.stanford.edu/data/).

These files contains anonymized profiles of Slovak social network "Pokec" [(get it here)](https://snap.stanford.edu/data/soc-pokec.html).

It contains 1632803 profiles with 30622564 relations. Each profile has 59 attributes.
The `./data` directory already contains file `columns.txt` which describes Profile's attributes.

Download and unzip files `soc-pokec-profiles.txt.gz` and `soc-pokec-relationships.txt.gz` into `./data` directory and then import data with `make import` .
