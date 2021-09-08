// 快速排序
package alg

func quickSort(data []int) {
	// 快速排序
	quickSortC(data,0,len(data) - 1)
}

func quickSortC(data []int,startPtr int,endPtr int) {
	if startPtr >= endPtr {
		return
	}

	partitionPtr := partition(data,startPtr,endPtr)
	quickSortC(data,startPtr,partitionPtr - 1)
	quickSortC(data,partitionPtr + 1,endPtr)
}

// 获取分区点
func partition(data []int,startPtr int,endPtr int) int {
	pivot := data[endPtr]
	i := startPtr

	for j := startPtr; j < endPtr; j++ {
		if data[j] < pivot {
			n := data[i]
			data[i] = data[j]
			data[j] = n
			i = i + 1
		}
	}

	n := data[i]
	data[i] = data[endPtr]
	data[endPtr] = n

	return i
}
