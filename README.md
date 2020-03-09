# 分布式爬虫设计
go get -u github.com/simonzs/crawler_go
### persist
ItemSaver 和 Engine模块之间通信
- RPCService: ItemSaver
- RPCClient: Engine模块

### Worker
```
type Request struct {
	URL        string
	ParserFunc ParserFunc
}

// ParserResult 提取结果
type ParserResult struct {
	Reuqests []Request
	Items    []Item
}
```

- 需要传递函数: ParserFunc
需要序列化和反序列化
#### 序列话
- Parser -> Serialized Parser -> Json字符串
#### 反序列化
- Parser <-  Serialized Parser <- Json字符串

# 运行
#### 运行Worker
- go run worker/server/worker.go
#### 运行ItemSaver
- go run persist/server/itemserver.go
### 执行engine
- go run main.go

# 命令运行
- go run main.go --itemsaver_host=":1234" --worker_host="9003"
- go run worker/server/worker.go  --port=9001
- go run persist/server/itemserver.go --port=1234