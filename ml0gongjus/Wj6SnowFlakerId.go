package ml0gongjus

import (
	"errors"
	"fmt"
	"strconv"
	"sync"
	"time"
)

const (
	SequenceBits       = uint(10)     //自增ID 所占用位置
	DistrictIdBits     = uint(5) //区域 所占用位置
	DistrictIdShift    = SequenceBits + NodeIdBits
	NodeId             = int64(2)     //自定义NodeId
	NodeIdBits         = uint(9)      //节点 所占位置
	NodeIdShift        = SequenceBits //左移次数
	SequenceMask       = -1 ^ (-1 << SequenceBits)
	MaxDistrictId      = -1 ^ (-1 << DistrictIdBits) //最大区域范围
	MaxNextIdsNum      = 100                         //单次获取ID的最大数量
	MaxNodeId          = -1 ^ (-1 << NodeIdBits)     //节点 ID 最大范围
	Twepoch            = int64(1417937700000)        // 默认起始的时间戳 1449473700000 。计算时，减去这个值
	TimestampLeftShift = SequenceBits + NodeIdBits + DistrictIdBits
)

var Idworker, _ = NewIdWorker()

func HuoquId() int64 {
	nextid, _ := Idworker.NextId()
	return nextid
}
func HuoQuIdZiFu() string {
	nextidyonghu := HuoquId()
	idZiFuyonghu := strconv.FormatInt(nextidyonghu, 10)
	return idZiFuyonghu
}

type IdWorker struct {
	sequence      int64 //序号
	lastTimestamp int64 //最后时间戳
	nodeId        int64 //节点ID
	twepoch       int64
	districtId    int64
	mutex         sync.Mutex
}

// NewIdWorker new a snowflake id generator object.
func NewIdWorker() (*IdWorker, error) {
	var districtId int64
	districtId = 1 //暂时默认给1 ，方便以后扩展
	idWorker := &IdWorker{}
	if NodeId > MaxNodeId || NodeId < 0 {
		fmt.Sprintf("NodeId Id can't be greater than %d or less than 0", MaxNodeId)
		return nil, errors.New(fmt.Sprintf("NodeId Id: %d error", NodeId))
	}
	if districtId > MaxDistrictId || districtId < 0 {
		fmt.Sprintf("District Id can't be greater than %d or less than 0", MaxDistrictId)
		return nil, errors.New(fmt.Sprintf("District Id: %d error", districtId))
	}
	idWorker.nodeId = NodeId
	idWorker.districtId = districtId
	idWorker.lastTimestamp = -1
	idWorker.sequence = 0
	idWorker.twepoch = Twepoch
	idWorker.mutex = sync.Mutex{}
	fmt.Sprintf("worker starting. timestamp left shift %d, District id bits %d, worker id bits %d, sequence bits %d, workerid %d",
		TimestampLeftShift, DistrictIdBits, NodeIdBits, SequenceBits, NodeId)
	return idWorker, nil
}

// timeGen generate a unix millisecond.
func timeGen() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

// tilNextMillis spin wait till next millisecond.
func tilNextMillis(lastTimestamp int64) int64 {
	timestamp := timeGen()
	for timestamp <= lastTimestamp {
		timestamp = timeGen()
	}
	return timestamp
}

// NextId get a snowflake id.
func (id *IdWorker) NextId() (int64, error) {
	id.mutex.Lock()
	defer id.mutex.Unlock()
	return id.nextid()
}

// NextIds get snowflake ids.
func (id *IdWorker) NextIds(num int) ([]int64, error) {
	if num > MaxNextIdsNum || num < 0 {
		fmt.Sprintf("NextIds num can't be greater than %d or less than 0", MaxNextIdsNum)
		return nil, errors.New(fmt.Sprintf("NextIds num: %d error", num))
	}
	ids := make([]int64, num)
	id.mutex.Lock()
	defer id.mutex.Unlock()
	for i := 0; i < num; i++ {
		ids[i], _ = id.nextid()
	}
	return ids, nil
}
func (id *IdWorker) nextid() (int64, error) {
	timestamp := timeGen()
	if timestamp < id.lastTimestamp {
		//    fmt.Sprintf("clock is moving backwards.  Rejecting requests until %d.", id.lastTimestamp)
		return 0, errors.New(fmt.Sprintf("Clock moved backwards.  Refusing to generate id for %d milliseconds",
			id.lastTimestamp-timestamp))
	}
	if id.lastTimestamp == timestamp {
		id.sequence = (id.sequence + 1) & SequenceMask
		if id.sequence == 0 {
			timestamp = tilNextMillis(id.lastTimestamp)
		}
	} else {
		id.sequence = 0
	}
	id.lastTimestamp = timestamp
	return ((timestamp - id.twepoch) << TimestampLeftShift) |
		(id.districtId << DistrictIdShift) |
		(id.nodeId << NodeIdShift) |
		id.sequence, nil
}
