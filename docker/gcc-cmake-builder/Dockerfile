FROM rikorose/gcc-cmake

MAINTAINER Alex Mushket <uzibazuka@gmail.com>

# RUN apt-get update -y && apt-get install uuid-dev

RUN mkdir /builder

WORKDIR /builder
ADD builder.sh .
RUN chmod +x builder.sh