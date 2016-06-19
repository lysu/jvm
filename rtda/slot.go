package rtda

import "github.com/lysu/jvm/rtda/heap"

type Slot struct {
	num int32
	ref *heap.Object
}
