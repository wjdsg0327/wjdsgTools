package RandomUtils

import "math/rand"

type RandomUtils struct {
}

// GenerateRandomString 生成随机数
// Capital：大写字母
// Lowercase：小写字母
// Symbol ：符号
// Number ：数字
// length ：需要生成多少位
func (RandomUtils) GenerateRandomString(Capital bool, Lowercase bool, Symbol bool, Number bool, length int) string {

	charset := ""

	if Capital == true {
		charset = charset + capital()
	}
	if Lowercase == true {
		charset = charset + lowercase()
	}
	if Symbol == true {
		charset = charset + symbol()
	}
	if Number == true {
		charset = charset + number()
	}

	randomString := make([]byte, length)
	for i := range randomString {
		randomString[i] = charset[rand.Int63n(int64(len(charset)))]
	}

	return string(randomString)

}
