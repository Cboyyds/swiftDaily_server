package utils

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// 原来这个随机数生成器，是我自己生成的啊，不是qq邮箱自动生成给我返回来
func GenerateVerificationCode(length int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return fmt.Sprintf("%0*d", length, r.Intn(int(math.Pow10(length))))
}
