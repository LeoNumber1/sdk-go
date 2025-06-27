# sdk for go


## Getting started

This sdk is only for go in ikaopu.cn.

See examples in [tests](tests).

You can contact liyuan@ikaopu.cn if you have any question.

- Install
```bash
go get -u github.com/LeoNumber1/sdk-go@v0.0.1-alpha2
```

- then import it in your code
```go
import "github.com/LeoNumber1/sdk-go"
```

- use it
```go
kpClient := client.NewClient(client.Config{
    BaseURL:   "http://localhost:3000",
    AccessKey: "YOUR AK",
    SecretKey: "YOUR SK",
}
insClient := api.NewInstanceAPI(kpClient)
insClient.GetInstance("instanceId")

marketClient := api.NewMarketAPI(kpClient)
marketClient.ListMarketCommon()
```

Try it yourself and enjoy your code!