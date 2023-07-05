package main
 
import (
	"fmt"
 
	"github.com/zhaozhongyuan425/hello/v2" // 引入v2版本
)
 
func main() {
	fmt.Println("现在是假期时间...")
	hello.SayHi("张三") // v2版本的SayHi函数需要传入字符串参数
}
