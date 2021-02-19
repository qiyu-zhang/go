package main
import "fmt"
//鸽子
type  dove interface {
	gugugu()
}
//复读机
type repeater interface {
	repeat(string)
}
//柠檬精
type  lemon interface {
	acid()
}
//真香怪
type smellnice interface {
	nice()
}
//人类
type person struct {
	name   string // 姓名
	age    int  // 年龄
	gender  string // 性别
}
//咕咕
func (p *person) gugugu(){
	fmt.Println(p.name," 咕咕咕")
}
//复读
func (p *person) repeat(word string) {
	fmt.Println(word)
}
//恰柠檬
func (p *person) acid(){
	fmt.Println("恰柠檬")
}
//真香
func (p *person) nice(){
	fmt.Println("真香 ")
}
func main(){
	p := &person{
		name:"balabala",
		age:18,
		gender: "male",
	}
	//鸽子
	p.gugugu()
	//复读机
	p.repeat("go go go go")
	//柠檬精
	p.acid()
	//真香怪
	p.nice()
}