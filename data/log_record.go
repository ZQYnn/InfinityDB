package data

// LogRecordPos  数据内存索引, 表示数据在磁盘的位置
type LogRecordPos struct {
	Fid    uint32 // 文件ID， 表示将数据存储到那个文件当中
	Offset int64  //  偏移量， 表示文件存储到数据文件中的位置
}
