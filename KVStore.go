package KVStore

type Item struct {
	Key string
	Val int
}

//type Store []Item

type Store map[string]int

// Checks if the item exists in the store. If the item exists it returns index or else returns -1.
func (s Store) checkIfItemsExistsInStore(key string) int {
	if value, ok := s[key]; !ok{
		return -1
	} else {
		return value
	}
	/*
	for index, item := range s {
		if item.Key == key {
			return index
		}
	}
	return -1*/
}

// Get returns value associated with key, if the item exists or else returns -1.
func (s Store) Get(key string) int {
	return s.checkIfItemsExistsInStore(key)
	/*
	index := s.checkIfItemsExistsInStore(key)
	if index != -1 {
		return s[index].Val
	}
	return -1*/
}

// Put inserts the given item in the store. On success, it returns a 0 and if the item already exists it returns a -1.
func (s Store) Put(item Item) int {
	index := s.checkIfItemsExistsInStore(item.Key)
	if index == -1 {
		s[item.Key] = item.Val
		return 0
	}
	return -1

	/*index := s.checkIfItemsExistsInStore(item.Key)
	if index == -1 { // Item doesn't exist. So the item could be safely inserted into store.
		s = append(s, item)
		return 0
	}
	return -1*/
}
