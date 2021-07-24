package common

import (
	"crypto/md5"
	"fmt"
	"io"
)

type password struct{}

// Password , Support Contains Append Method to Array
var Password = password{}

// Contains , 检查是否包含某String
func (a *password) Valid(str string, passwordstr string) bool {
	return a.Generate(str) == passwordstr
}

func (a *password) Generate(str string) string {
	w := md5.New()
	io.WriteString(w, str)
	md5str := fmt.Sprintf("%x", w.Sum(nil))
	return md5str
}
