# Highlander API

A simple API that is able process just one request at once. Not more and not less. Exactly one.

```
there can be only one
```

The highlander API can be packed as a normal binary or using docker.



# Why?

1. Just for fun. I have created it after we were kidding about slow APIs. The question was how hard is it to write a API like this in go or node?
2. One the other side it could be great to have a boilderplate for request throtteling. For example for people who are flooding your APIs.



# Status

This is a very early version of highlander API and is only processing GET requests. In future I could Image to extend to to a reverse proxy with highlander capabilities.



# How to ... ?


## How to build

The normal build is building the highlander docker container.

```
./build.sh
```

You can also set the variable `DEV` in case you want to just build the software and not package it using docker. 

```
DEV=1 ./build.sh
```


## How to run

The normal run is running highlander in a docker container

```
./run.sh
```

You can also set the variable `DEV` in case you want to just run the software and run it using docker.

```
DEV=1 ./run.sh
```


## How to access

```
curl -k https://localhost:8444
```


## How to configure

Put a file `highlander.yaml` to the currect directory. The content can look like this:
For a configuration example, please have a look into the file `examples/highlander.yaml`



## How to test the performance of highlander API

I have tested highlander API using [siege](https://www.joedog.org/siege-home/)
You can run performances tests easily using 

```
siege -r 1 -c 100 --no-follow https://localhost:8444
```


## Hot to debug

You can inspect the running docker container using

```
docker log -f highlander
```


## How to contribute

fork -> create branch -> create pull request

