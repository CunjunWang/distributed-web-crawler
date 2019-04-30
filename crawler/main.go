package main

import (
	"distributed-web-crawler/crawler/dating/parser"
	"distributed-web-crawler/crawler/engine"
	"distributed-web-crawler/crawler/scheduler"
)

const cityUrl = "http://www.zhenai.com/zhenghun"

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
	}

	e.Run(engine.Request{
		Url:        cityUrl,
		ParserFunc: parser.ParseCityList,
	})

	//e.Run(engine.Request{
	//	Url:        "http://www.zhenai.com/zhenghun/shanghai",
	//	ParserFunc: parser.ParseCity,
	//})
}

// Parser 解析器
// input: utf-8编码文本
// output: Request{URL, 对应Parser}列表, Item列表
