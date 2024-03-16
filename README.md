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





## stringUtils字符串相关

| 函数名       | 说明                        |
| ------------ | --------------------------- |
| StringToJson | 字符串转json：map或者结构体 |
| JsonToSting  | json转字符串                |
| EncodeHexStr | 字符串转hex                 |
| DecodeHexStr | hex转字符串                 |

## DateUtils日期相关

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
| AgeOfNow       | 计算年龄                                     |
| IsLeapYear     | 是否闰年                                     |

## SecureUtils加密相关

| 函数名  | 说明                |
| ------- | ------------------- |
| Md5     | md5加密（无法解密） |
| Encrypt | 加密                |
| Decrypt | 解密                |

## StructUtils结构体相关

| 函数名             | 说明                             |
| ------------------ | -------------------------------- |
| CopyProperties     | 结构体转换                       |
| GetStructInfoList  | 获取结构体所有信息               |
| GetStructFieldList | 获取结构体所有字段               |
| GetStructFieldInfo | 根据结构体字段名字获取字段的详细 |

## IdUtils  id生成

| 函数名         | 说明                            |
| -------------- | ------------------------------- |
| GetId          | 根据时间戳生成id                |
| GetSnowflakeId | 雪花算法（Snowflake）的 ID 生成 |

## RandomUtils 随机数

| 函数名               | 说明       |
| -------------------- | ---------- |
| GenerateRandomString | 生成随机数 |

## CalculateUtils 计算相关

| 函数名         | 说明                   |
| -------------- | ---------------------- |
| PercentageCode | 求两数的百分比(带符号) |
| Percentage     | 求两数的百分比         |

## FileUtils 文件相关

| 函数名 | 说明         |
| ------ | ------------ |
| Copy   | 文件拷贝     |
| Read   | 读取文件内容 |

## InterfaceUtils 接口类型相关

| 函数名  | 说明                             |
| ------- | -------------------------------- |
| GetType | 获取接口实际类型                 |
| IsType  | 检查接口的类型是否为某个特定类型 |

## IdCardUtils 身份证相关

| 函数名        | 说明                                                |
| ------------- | --------------------------------------------------- |
| IsValidCard18 | 判断身份证是否合法:只能判断中国大陆第二代身份证18位 |
| IdCardGetAge  | 根据身份证计算年龄                                  |
| IdCardGetAREA | 根据身份证获取所属省份                              |
| IdCardSEX     | 根据身份证获取性别                                  |

## DesensitizedUtils 信息脱敏相关

| 函数名           | 说明       |
| ---------------- | ---------- |
| ChineseName      | 中文名称   |
| Email            | 电子邮件   |
| MobilePhone      | 手机号     |
| IdCard           | 身份证     |
| Address          | 地址       |
| PossWord         | 密码       |
| SecureStructList | 结构体脱敏 |

