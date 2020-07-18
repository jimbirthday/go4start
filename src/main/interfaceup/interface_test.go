package httpup

import (
	"fmt"
	"testing"
)

//接口本身是调用方和实现方需要遵守的一中协议，大家按照统一的方法命名参数类型和数量来协调逻辑处理的过程
//在go语言中只有实现interface中的所有方法才能是实现，否则只是声明了一个自定义方法
//一个类型可以实现多个方法，同样的多个类型可以实现一个方法，这个在java里对应的思想就是继承和实现
//但是java中一个类型只能实现一个接口，也只能继承一个接口，如果想继承多个，只能按层级继承
//两者对比体现出go语言的灵活性，在向上封装的框架或者业务代码中，更精细化的构思能够减少大量重复代码的编写
//例如如果设计一个监听机制，这个监听机制不仅仅可用于框架启动时的监听，同样可以使用在db层的执行监听上
//在go语言中可能只需要设计一套接口或者机制，就可以在不同的地方各自实现
//在java中也能够使用一套机制，但是需要更加抽象的思想来提取两个地方的共同性质，才能封装，同样也可以使用泛型或者反射来实现，这是java的特质
//但是java中难做的也是这一点，无论使用java哪种方式实现，实际上都需要你具有一定的java代码的设计能力，这里可能会需要你对设计模式的一些知识
//在java代码中实现一个策略模式的时间，可能会比在go语言中实现要求的时间多一些，这是语言特性导致的，无论你是新手还是老江湖

type DataWriter interface {
	WriteData(data interface{}) error
	WriteData1(data interface{}) error
}

type file struct {
}

func (a *file) WriteData(data interface{}) error {
	fmt.Println("data is on write ~~~ ", data)
	return nil
}

func (a *file) WriteData1(data interface{}) error {
	return nil
}

func Test1(T *testing.T) {
	f := new(file)
	_ = f.WriteData("file is loading")
}
