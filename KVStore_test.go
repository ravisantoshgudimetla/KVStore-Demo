package KVStore

import ("testing"
	"strconv"
	"fmt"
)

var cacheSize= int(50000)
var store = make(Store, cacheSize)

func populateKVStore(){
	var item Item
	for i:= 0; i<cacheSize; i++ {
		item.Key = strconv.Itoa(i)
		item.Val = i
		store.Put(item)
	}
}


func BenchmarkStore_Put(b *testing.B) {
	populateKVStore()
}

func BenchmarkStore_Get(b *testing.B) {
	for i:=0; i< cacheSize; i++ {
		fmt.Println(store.Get(strconv.Itoa(i)))
	}
}