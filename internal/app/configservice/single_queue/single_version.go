package single_queue

//这里维护一个关于版本号的map集合以借助mysql完成简单的消息队列
var v map[string]uint64 = make(map[string]uint64)

func GetV() map[string]uint64 {
	return v
}
