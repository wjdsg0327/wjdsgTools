# wjdsgTools

go语言工具集

使用教程

```
go get github.com/wjdsg0327/wjdsgTools
```

```go
func main() {

	id := wjdsgtools.Utils{}.IdUtils().GetId()

	fmt.Println(id)

}
```





## stringUtils

| 函数名       | 说明                        |
| ------------ | --------------------------- |
| Split        | 根据指定分割符，分割字符串  |
| StringToJson | 字符串转json：map或者结构体 |
| JsonToSting  | json转字符串                |

## DateUtils

| 函数名         | 说明                                         |
| -------------- | -------------------------------------------- |
| ParseYMDHMS    | 将数据转换为年月日时分秒                     |
| ParseYMD       | 将数据转换为年月日                           |
| Format         | 字符串时间转time                             |
| BeginOfDay     | 获取指定日期的开始时间                       |
| EndOfDay       | 获取指定日期的结束时间                       |
| Offset         | 计算偏移后的日期时间                         |
| BeginOfMonth   | 获取指定月开始时间                           |
| EndOfMonth     | 获取指定月结束时间                           |
| Between        | 计算两个时间之间的年月日                     |
| GetBetweenDate | 获取两个日期之间的所有日期，格式：yyyy-MM-dd |
| Yesterday      | 获取昨天的日期                               |
| Tomorrow       | 获取昨天的日期                               |
| LastWeek       | 获取上周的日期                               |
| NextWeek       | 获取下周的日期                               |
| LastMonth      | 获取上个月的日期                             |
| NextMonth      | 获取下个月的日期                             |
| GetZodiac      | 根据月日计算星座                             |

## Md5Utils

| 函数名  | 说明                |
| ------- | ------------------- |
| Md5     | md5加密（无法解密） |
| Encrypt | 加密                |
| Decrypt | 解密                |

## StructUtils

| 函数名             | 说明                             |
| ------------------ | -------------------------------- |
| CopyProperties     | 结构体转换                       |
| GetStructInfoList  | 获取结构体所有信息               |
| GetStructFieldList | 获取结构体所有字段               |
| GetStructFieldInfo | 根据结构体字段名字获取字段的详细 |

## IdUtils

| 函数名         | 说明                            |
| -------------- | ------------------------------- |
| GetId          | 根据时间戳生成id                |
| GetSnowflakeId | 雪花算法（Snowflake）的 ID 生成 |

## RandomUtils

| 函数名               | 说明       |
| -------------------- | ---------- |
| GenerateRandomString | 生成随机数 |

## CalculateUtils

| 函数名         | 说明                   |
| -------------- | ---------------------- |
| PercentageCode | 求两数的百分比(带符号) |
| Percentage     | 求两数的百分比         |

