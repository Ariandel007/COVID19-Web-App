package knn

import (
	"sync"
)

func Merge(ldata [][]float32, rdata [][]float32) (result [][]float32) {
	result = make([][]float32, len(ldata)+len(rdata))
	lidx, ridx := 0, 0

	for i := 0; i < cap(result); i++ {
		switch {
		case lidx >= len(ldata):
			result[i] = rdata[ridx]
			ridx++
		case ridx >= len(rdata):
			result[i] = ldata[lidx]
			lidx++
		case ldata[lidx][len(ldata[0])-1] < rdata[ridx][len(ldata[0])-1]:
			result[i] = ldata[lidx]
			lidx++
		default:
			result[i] = rdata[ridx]
			ridx++
		}
	}
	return
}

func SingleMergeSort(data [][]float32) [][]float32 {
	if len(data) < 2 {
		return data
	}
	middle := len(data) / 2
	return Merge(SingleMergeSort(data[:middle]), SingleMergeSort(data[middle:]))
}
func MultiMergeSortWithSem(data [][]float32, sem chan struct{}) [][]float32 {
	if len(data) < 2 {
		return data
	}

	middle := len(data) / 2

	wg := sync.WaitGroup{}
	wg.Add(2)

	var ldata [][]float32
	var rdata [][]float32

	select {
	case sem <- struct{}{}:
		go func() {
			ldata = MultiMergeSortWithSem(data[:middle], sem)
			<-sem
			wg.Done()
		}()
	default:
		ldata = SingleMergeSort(data[:middle])
		wg.Done()
	}

	select {
	case sem <- struct{}{}:
		go func() {
			rdata = MultiMergeSortWithSem(data[middle:], sem)
			<-sem
			wg.Done()
		}()
	default:
		rdata = SingleMergeSort(data[middle:])
		wg.Done()
	}

	wg.Wait()
	return Merge(ldata, rdata)
}

func RunMultiMergesortWithSem(data [][]float32) [][]float32 {
	sem := make(chan struct{}, 4)
	return MultiMergeSortWithSem(data, sem)
}