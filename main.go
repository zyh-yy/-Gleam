package main

import (
	"log"
	"math/rand"
	"time"

	epool "github.com/zhaomin1993/pool/easy_go_pool"
)

type Score struct {
	Num int
}

func (s Score) Do() {
	if s.Num%10000 == 0 {
		log.Println("num:", s.Num)
	}
	if s.Num%2 == 0 {
		panic(s.Num)
	}
	time.Sleep(time.Millisecond * 100)
}

func main() {
	//创建协程池
	pool := epool.NewWorkerPool(1000, 1100)
	defer pool.Close() //关闭协程池
	pool.OnPanic(func(msg interface{}) {
		//log.Println("error:", msg)
	})
	datanum := 100 * 100 * 10
	for i := 1; i <= datanum; i++ {
		//接收任务
		if err := pool.Accept(Score{Num: i}); err != nil {
			log.Println("err:\t", err)
			break
		}
		if i%10000 == 0 {
			log.Println("send num:", i)
			randNum := rand.Intn(10) + 1000
			//调整协程池大小
			pool.AdjustSize(uint16(randNum))
		}
	}
}
