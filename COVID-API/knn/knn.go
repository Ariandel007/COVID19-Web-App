package knn

import (
	"fmt"
	"strconv"
	"math"
	"sync"
)
func euclidean_distance(row1 ,row2 []float32) (float32){
	distance:=0.0
	for i:=0; i<len(row1)-2;i++{
		distance+=math.Pow(float64(row1[i]) - float64(row2[i]),2)
	}
	return float32(math.Sqrt(distance))
}
func get_neighbors(train [][]float32, test_row []float32,num_neighbors int)([][]float32){
	sem := make(chan struct{}, 4)
	wg := sync.WaitGroup{}

	for k,train_row:=range train{
		sem<-struct{}{}
		wg.Add(1)
		go func(train_row []float32,k int) {
			ktmp:=k
			distance:=euclidean_distance(test_row,train_row)
			train[ktmp][len(train_row)-1]=distance
			<-sem
			wg.Done()
		}(train_row,k)
	}
	wg.Wait()

	trainOrdered:=RunMultiMergesortWithSem(train)
	
	neighbors:=make([][]float32,num_neighbors)
	for i:=0;i<num_neighbors;i++{
		neighbors[i]=trainOrdered[i]
	}
	return neighbors
}

func get_most_repeated_elem(data []float32)int{
	dict:=make(map[string]int)
	for _,val:=range data{
		key := fmt.Sprintf("%f", val)
		if v,ok:=dict[key];ok{
			dict[key]=v+1
		}else{
			dict[key]=1
		}
	}
	higher_val:=-1
	higher_output:=""
	for k, v := range dict { 
		if v>higher_val{
			higher_val=v
			higher_output=k
		}
	}
	res,_:=strconv.ParseFloat(higher_output,32)
	return int(res)
}

func Predict_classification(train [][]float32,test_row []float32,num_neighbors int)(int, [][]float32){
	neighbors:=get_neighbors(train,test_row,num_neighbors)
	output_values:=make([]float32,num_neighbors)
	for k,v:= range neighbors{
		output_values[k]=v[len(train[0])-2]
	}
	prediction:=get_most_repeated_elem(output_values)
	return prediction,neighbors
}