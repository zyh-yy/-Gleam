package utils

// import(
// 	"github.com/zhaomin1993/pool"
// )

// //封装后的线程池

// func NewPool(poolSize int)  {
// 	pool.NewWorkerPool()
// }

type interface 


type Loader struct{
	channel []interface{} input
	channel []interface{} output
}

func (l*Loader)Do(){
	
}


//会根据idl生成出req的proxy，所有对象都需要通过get方法调用，只有可以写入某个对象的方法才能

//支持配置大字段，但是生成会按照每个字段生成

//初始化所有channel

//一个dag里上下是一对channel

//只能拿到值对象

//写入只能写入上下文可用对象/res结果对象

//结果对象需要注册的时候写入path，最后统一合并









