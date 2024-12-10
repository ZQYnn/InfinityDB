package index

import (
	"bytes"
	"github.com/ZQYnn/InfinityDB/data"
	"github.com/google/btree"
)

// Index 定义抽象接口 后续加入其他的数据结构 可以直接实现这个接口
// 基本实现的方式
type Index interface {
	Put(key []byte, pos *data.LogRecordPos) bool
	Get(key []byte) *data.LogRecordPos
	Delete(key []byte) bool
}

type Item struct {
	key []byte
	pos *data.LogRecordPos
}

func (ai *Item) Less(bi btree.Item) bool {
	return bytes.Compare(ai.key, bi.(*Item).key) == -1
}
