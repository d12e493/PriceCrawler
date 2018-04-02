# PriceCrawler
## API Server & Batch Job

## Architectue

### Initial design
1. Crawler and API Server
![Crawler and API Server](doc/crawler_and_api_server.png)

2. Schema and Worker Pool
![Schema and Worker Pool](doc/schema_and_worker_pool.png)
### Implement for POC

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
