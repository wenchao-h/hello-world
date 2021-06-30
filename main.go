package main
//该处需要导入一个打印支持的包
//Java与之类似的是jdk自带的输出
import "fmt"

//与Java类似的是，一个go文件的运行同样需要main方法，不过go写法更加简单，且不需要做static处理，同级别下都能访问
func main() {
	//打印输出，后面‘ ；’可有可无，go语言追求一种简洁和高效性
	fmt.Println("hellword")
}
