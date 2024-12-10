package fio

const DataFilePerm = 0644

type IOManager interface {
	// Read 读取指定位置的文件
	Read([]byte, int64) (int, error)

	// Write 写入字节数组到文件中
	Write([]byte, int64)

	// Sync  持久化数据
	Sync() error

	// Close 关闭文件
	Close() error
}
