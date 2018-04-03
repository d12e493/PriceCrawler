# API Server & Price Crawler Job
This project consists of RESTful API server and price crawler batch job.
The architecture of initital design is written on the first day.
Due to time considerations , I implement the system by the architecture of POC.

## Contents
* [Architectue](#Architectue)
* [How to build](#how-to-build) 
* [How to run](#how-to-run)
* [API](#api)
* [Other Improvement](#other-improvement)
---
## Architectue

* Initial design
1. Crawler and API Server
<br/>
<img src="doc/crawler_and_api_server.png" alt="Crawler and API Server" width="200">
</img>
2. Schema and Worker Pool
<br/>
<img src="doc/schema_and_worker_pool.png" alt="Schema and Worker Pool" width="200">
</img>
3. Implement for POC

## How to build
<pre>sh build.sh</pre>
## How to run
1. run api server (port:11968)
<pre>
sh script/api.sh start|stop|status
</pre>
2. price crawler
<pre>
sh script/priceCrawler.sh
</pre>

## API

## Other Improvement
1. Error handle and status code for API.
2. Modulize service and dao for project.
3. Optimize price crawler and cache some fixed value (like category,item...)
4. Monitor api server status and get metrics.
