# GO語言實作pub-sub設計模式

## project structure
```
─┬ go-publisher-subscriber
 └─┬ go.mod  
   │  └─用指令建立, go mod init go-publisher-subscriber
   ├ go.sum
   │  └─
   ├ hub.go
   │  └─
   ├ main.go 
   │  └─建立 subscriber, hub 並執行 publish, subscribe functions
   ├ main_test.go  
   │  └─單元測試, 使用 testify 套件
   └ subscriber.go
      └─
```