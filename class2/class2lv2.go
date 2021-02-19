package main

import "fmt"

func Receiver(v interface{})  {//判断参数类型
	switch v.(type){
	case string:
		fmt.Println("这个是string")
	case int:
		fmt.Println("这个是int")
	case bool:
		fmt.Println("这个是bool")
	}
}
func main()  {
	Receiver("你好吗")//"这个是string"
	Receiver(32)	//"这个是int"
	Receiver(true) //"这个是bool"
}

