# API Server & Price Crawler Job
This project consists of RESTful API server and price crawler batch job.
The architecture of initital design is written on the first day.
Due to time considerations , I implement the system by the architecture of POC.

## Contents
* [Architectue](#Architectue)
* [How to build](#how-to-build) 
* [How to run](#how-to-run)
* [API](#api)
* [Other Improvement](#other-improvement)
---
## Architectue

* Initial design
1. Crawler and API Server
![Crawler and API Server](doc/crawler_and_api_server.png)

2. Schema and Worker Pool
![Schema and Worker Pool](doc/schema_and_worker_pool.png)
* Implement for POC

## How to build
<pre>sh build.sh</pre>
## How to run
1. run api server
<pre>
sh script/api.sh start|stop|status
</pre>
2. price crawler
<pre>
sh script/priceCrawler.sh
</pre>

## API

## Other Improvement