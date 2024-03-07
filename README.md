# wjdsgTools

go语言工具集





## stringUtils

| 函数名 | 说明                       |
| ------ | -------------------------- |
| Split  | 根据指定分割符，分割字符串 |
|        |                            |
|        |                            |

## DateUtils

| 函数名             | 说明                     |
| ------------------ | ------------------------ |
| ParseYMDHMS        | 将数据转换为年月日时分秒 |
| ParseYMD           | 将数据转换为年月日       |
| Format             | 字符串时间转time         |
| BeginOfDay         | 获取指定日期的开始时间   |
| EndOfDay           | 获取指定日期的结束时间   |
| Offset             | 计算偏移后的日期时间     |
| BeginOfMonth       | 获取指定月开始时间       |
| EndOfMonth         | 获取指定月结束时间       |
| GetTwoTimeRangeDay | 计算两个时间之间的天数   |

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

