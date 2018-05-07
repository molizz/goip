## goip

将各大ip提供商进行了一定的封装，直接使用

### 注意

高德（amap）只支持国内ip

### 使用方式


`go get github.com/molizz/goip`

```golang

import "github.com/molizz/goip"

taobao := platform.NewTaobao()
tencent := platform.NewTencent("key")
goip.AddPlatform(taobao)
goip.AddPlatform(tencent)

location, err := goip.GetLocation("35.185.191.24")
fmt.Println(location.City)

```

你可以通过AddPlatform添加多个platform，比如上面添加了2个平台，当淘宝的无法使用时，将自动使用腾讯的。

因为大部分平台都是有并发限制的。

所以建议获取ip时操作时，请异步处理（如果你的并发比较高的话）

### 有的平台需要Key

有的平台是需要申请key的

- 腾讯地图api： http://lbs.qq.com/webservice_v1/guide-ip.html
- 高德地图api： http://lbs.amap.com/api/webservice/guide/api/ipconfig

请自行申请