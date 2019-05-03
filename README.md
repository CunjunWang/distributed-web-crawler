# distributed-web-crawler
course project for Introduction of Golang on imooc.
## Tech stack
1. Golang 1.11
2. Elasticsearch 6.5.4
3. Docker

## 单机版 / Single node version
/crawler

## 分布式版 / Distributed version
/crawler-distributed

## 前端页面 / Simple front end page
/frond-end

## 启动 / Run
### 单机版 / Single node version :
`docker run -d -p 9200:9200 elasticsearch:x.x.x (your es version)`

`(under project root directory) cd crawler`

`go run main.go`

### 分布式版 / Distributed version :
`docker run -d -p 9200:9200 elasticsearch:x.x.x (your es version)`

`(under crawler-distributed) cd persist`

`go run itemSaver.go`

`(under crawler-distributed) cd worker/server`

`go run worker.go (start as many server as you want, as long as you add port configuration and set them in config.go)`

`(under project root directory) cd crawler-distributed`

`go run main.go`

### 页面 / Simple front end page :
`(under project root directory) cd front-end`

`go run start.go`

## Todo List
1. Crawl more website, with css selector or xpath (instead of regular expression).
2. Handle with anti-crawl mechanism (qps limit, encrypted cookie), or follow robots agreement.
3. Login mechanism.
4. Put De-dup in a separate module (with Redis).
5. Optimize ES search quality.
6. A more handy front-end page.
7. Use these data to play with AI.
8. Use Docker + Kubernetes to package and deploy.