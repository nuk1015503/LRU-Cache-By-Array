package main

import "fmt"

func main() {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Get(1)
	cache.Put(3, 3)
	cache.Get(2)
	cache.Put(4, 4)
	cache.Get(1)
	cache.Get(3)
	cache.Get(4)
}

// Data data
type Data struct {
	Key int
	Val int
}

// LRUCache LRUCache
type LRUCache struct {
	len      int
	dataList []Data
}

// Constructor 初始化
func Constructor(capacity int) LRUCache {
	var cache LRUCache
	cache.len = capacity
	return cache
}

// Show 顯示資料
func (c *LRUCache) Show() {
	for _, val := range c.dataList {
		fmt.Println(val.Key, "val", val.Val)
	}
}

// Get 根據key 取得 資料 若不存在回傳-1
func (this *LRUCache) Get(key int) int {
	isExist := false
	index := 0
	tmp := Data{}
	for i, val := range this.dataList {
		if val.Key == key {
			isExist = true
			index = i
			tmp = val
			break
		}
	}

	if isExist {
		for ; index < len(this.dataList)-1; index++ {
			this.dataList[index] = this.dataList[index+1]
		}
		this.dataList[index] = tmp
		return tmp.Val
	}

	return -1
}

// Put 根據key 存放 value 若是超過上限將移除最遠沒有使用到的資料
func (this *LRUCache) Put(key int, value int) {
	isExist := false
	index := 0
	for i, val := range this.dataList {
		if val.Key == key {
			isExist = true
			index = i
			break
		}
	}

	if isExist {
		for ; index < len(this.dataList)-1; index++ {
			this.dataList[index] = this.dataList[index+1]
		}
		this.dataList[index] = Data{Key: key, Val: value}
	} else {
		if len(this.dataList) == this.len {
			for i := 0; i < this.len-1; i++ {
				this.dataList[i] = this.dataList[i+1]
			}
			this.dataList[this.len-1] = Data{Key: key, Val: value}
		} else {
			this.dataList = append(this.dataList, Data{Val: value, Key: key})
		}
	}
}
