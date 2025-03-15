package channel_pool

import "sync"

var channelPool = sync.Map{}

// 为每个scene生成一套初始化方法
// 所有的task下依赖和产出的字段的channel
