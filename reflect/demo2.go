package main

import (
	"fmt"
	"reflect"
)

type B struct {
	thing int
}

func (b *B) change(i int) { b.thing = i }

func (b B) write() string {
	b.thing = 100
	return fmt.Sprint(b)
}

// get set
type Person struct {
	firstName string
	lastName  string
}

func (p *Person) FirstName() string {
	return p.firstName
}

func (p *Person) SetFirstName(newName string) {
	p.firstName = newName
}

func main() {

	var b1 B // b1是值
	b1.change(1)
	fmt.Println(b1.write())
	fmt.Println("b1:", b1)

	//试着在 write() 中改变接收者b的值：将会看到它可以正常编译，但是开始的 b 没有被改变。
	b2 := new(B) // b2是指针
	b2.change(2)
	fmt.Println(b2.write())
	fmt.Println("b2:", *b2)
	fmt.Println("====================")

	var x float64 = 3.4
	fmt.Println("type:", reflect.TypeOf(x))
	// 值复制
	//当 v := reflect.ValueOf(x) 函数通过传递一个 x 拷贝创建了 v，那么 v 的改变并不能更改原始的 x。
	//要想 v 的更改能作用到 x，那就必须传递 x 的地址 v = reflect.ValueOf(&x)
	v := reflect.ValueOf(x)
	fmt.Println("value:", v)
	fmt.Println("type:", v.Type())
	fmt.Println("kind:", v.Kind())
	fmt.Println("value:", v.Float())
	fmt.Println(v.Interface())

	//变量 v 的 Interface() 方法可以得到还原（接口）值，所以可以这样打印 v 的值：fmt.Println(v.Interface())
	fmt.Printf("value is %5.2e\n", v.Interface())
	y := v.Interface().(float64)
	fmt.Println(y)

	//panic: reflect: reflect.Value.SetFloat using unaddressable value
	fmt.Println("canAddr:", v.CanAddr(), " canSet:", v.CanSet())
	//v.SetFloat(3.1415)

	////panic: reflect: reflect.Value.SetFloat using unaddressable value
	//v = reflect.ValueOf(&x)
	//v.SetFloat(3.1415)

	v = reflect.ValueOf(&x)
	// v需要是指针
	v = v.Elem()
	fmt.Println("canAddr:", v.CanAddr(), " canSet:", v.CanSet())

	v.SetFloat(3.1415)
	fmt.Println("v:", v)
	fmt.Println("v.Interface:", v.Interface())

}
