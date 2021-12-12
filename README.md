# chinese-holidays-go

[![Go Reference](https://pkg.go.dev/badge/github.com/bastengao/chinese-holidays-go.svg)](https://pkg.go.dev/github.com/bastengao/chinese-holidays-go)
![badge](https://github.com/bastengao/chinese-holidays-go/workflows/Go/badge.svg)

提供具有中国特色的休假安排或者工作日查询。

Minnanalee: 在原作者基础上，
增加了Holiday（time.Time)(bool, string, error)接口，
Book可以导出节日名称，Book的IsHoliday，普通周末，返回false
Cache、Multiple没有实现节日名称。



## Install

    go get github.com/bastengao/chinese-holidays-go

## Usage

Bundle Query

```go
import (
    "github.com/bastengao/chinese-holidays-go/holidays"
)

queryer, err := holidays.BundleQueryer()
if err != nil {
    panic(err)
}

d := time.Date(2019, 10, 1, 0, 0, 0, 0, china)
queryer.IsHoliday(d)    // true
queryer.IsWorkingday(d) // false
```

Cache Queryer is a Queryer that fetches online data and check updates every day.

```go
queryer, err := holidays.NewCacheQueryer()
if err != nil {
    panic(err)
}

queryer.IsHoliday(d)
```

Multiple Queryer is a Queryer that delegates query to underlying multiple Queryers.
Try each queryers in order until one returns a result.

```go
bundleQueryer, err := holidays.BundleQueryer()
if err != nil {
    panic(err)
}

cacheQueryer, err := holidays.NewCacheQueryer()
if err != nil {
    panic(err)
}

queryer := holidays.NewMultipleQueryer(cacheQueryer, bundleQueryer)
queryer.IsHoliday(d)
```

## Features

- [x] bundled data
  - support [2022](http://www.gov.cn/zhengce/content/2021-10/25/content_5644835.htm)
  - support [2021](http://www.gov.cn/zhengce/content/2020-11/25/content_5564127.htm)
  - support [2020](http://www.gov.cn/zhengce/content/2019-11/21/content_5454164.htm)
  - support [2019](http://www.gov.cn/zhengce/content/2018-12/06/content_5346276.htm) and 5.1 [changes](http://www.gov.cn/zhengce/content/2019-03/22/content_5375877.htm)
  - support [2018](http://www.gov.cn/zhengce/content/2017-11/30/content_5243579.htm)
  - support [2017](http://www.gov.cn/zhengce/content/2016-12/01/content_5141603.htm)
  - support 2016
- [x] [online data](https://github.com/bastengao/chinese-holidays-data)
