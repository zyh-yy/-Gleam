package utils

import (
	"fmt"
	"sync"
	"time"

	"github.com/sony/sonyflake" // 改进版雪花算法库
)

// 原生实现（包含时钟回拨处理）
type Snowflake struct {
	mu            sync.Mutex
	epoch         time.Time // 起始时间
	lastTimestamp int64     // 上次生成时间
	nodeID        uint16    // 节点ID
	sequence      uint16    // 序列号（0-4095）
}

const (
	nodeBits     = 10 // 节点ID位数（支持1024节点）
	sequenceBits = 12 // 序列号位数
	timeShift    = nodeBits + sequenceBits
)

func NewSnowflake(nodeID uint16, epoch time.Time) *Snowflake {
	return &Snowflake{
		epoch:  epoch,
		nodeID: nodeID,
	}
}

func (s *Snowflake) Generate() (uint64, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	now := time.Since(s.epoch).Milliseconds()

	// 时钟回拨处理
	if now < s.lastTimestamp {
		return 0, fmt.Errorf("时钟回拨 %d 毫秒", s.lastTimestamp-now)
	}

	if now == s.lastTimestamp {
		s.sequence = (s.sequence + 1) & ((1 << sequenceBits) - 1)
		if s.sequence == 0 { // 当前毫秒序列号用尽
			for now <= s.lastTimestamp {
				now = time.Since(s.epoch).Milliseconds()
			}
		}
	} else {
		s.sequence = 0
	}

	s.lastTimestamp = now

	return uint64(now)<<timeShift | uint64(s.nodeID)<<sequenceBits | uint64(s.sequence), nil
}

// SonyflakeId 生产唯一id
func SonyflakeId() uint64 {
	// 设置起始时间为2020-01-01
	startTime, _ := time.Parse("2006-01-02", "2020-01-01")
	settings := sonyflake.Settings{
		StartTime: startTime,
		MachineID: func() (uint16, error) {
			return 1, nil // 动态获取机器ID
		},
	}

	sf := sonyflake.NewSonyflake(settings)
	id, _ := sf.NextID()
	return id
}
