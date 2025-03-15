package core

type manager struct {
	nodeCnt int //节点数量
}

// 所有node都是线程池，等待接收消息，然后执行
// 每个节点可以灵活控制吞吐量
// 执行完触发一个任务到下一个节点的执行队列中

type node struct {
	Name    string
	DepNode []string
}
