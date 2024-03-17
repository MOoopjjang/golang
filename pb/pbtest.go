package main

import (
	"fmt"
	"goproject/demo/pb/manager"
	"goproject/demo/pb/pbtype"
)

func rollback(cnt int, cache *manager.RollbackMgr) {

	for i := 0; i < cnt; i++ {
		if r, err := cache.Rollback(); err == nil {
			fmt.Println("[", i, "]", r.ToString())
		}

	}
}

func main() {
	fmt.Println("pbtest")

	fmt.Println("add test")
	manager.Cache.Initialzie()
	u1 := &pbtype.User{"cwkim", "1111", "incheon"}
	manager.Cache.Add(u1)
	// manager.Cache.Print()
	// fmt.Println("-------------")

	u1 = &pbtype.User{"ejkim", "2222", "incheon"}
	manager.Cache.Add(u1)
	// manager.Cache.Print()
	// fmt.Println("-------------")

	u1 = &pbtype.User{"bhkim", "3333", "incheon"}
	manager.Cache.Add(u1)
	// manager.Cache.Print()
	// fmt.Println("-------------")

	u1 = &pbtype.User{"khlee", "4444", "incheon"}
	manager.Cache.Add(u1)
	// manager.Cache.Print()
	// fmt.Println("-------------")

	u1 = &pbtype.User{"khlee", "5555", "incheon"}
	manager.Cache.Add(u1)
	// manager.Cache.Print()
	// fmt.Println("-------------")

	u1 = &pbtype.User{"khlee", "6666", "incheon"}
	manager.Cache.Add(u1)
	// manager.Cache.Print()
	// fmt.Println("-------------")

	u1 = &pbtype.User{"khlee", "7777", "incheon"}
	manager.Cache.Add(u1)
	// manager.Cache.Print()
	// fmt.Println("-------------")

	u1 = &pbtype.User{"khlee", "8888", "incheon"}
	manager.Cache.Add(u1)
	// manager.Cache.Print()
	// fmt.Println("-------------")

	u1 = &pbtype.User{"khlee", "9999", "incheon"}
	manager.Cache.Add(u1)
	// manager.Cache.Print()
	// fmt.Println("-------------")

	u1 = &pbtype.User{"khlee", "1000", "incheon"}
	manager.Cache.Add(u1)
	// manager.Cache.Print()
	// fmt.Println("-------------")

	u1 = &pbtype.User{"khlee", "1001", "incheon"}
	manager.Cache.Add(u1)
	// manager.Cache.Print()
	u1 = &pbtype.User{"khlee", "1002", "incheon"}
	manager.Cache.Add(u1)

	manager.Cache.Print()

	fmt.Println("######### Rollback ###########")

	rollback(2, &manager.Cache)
	manager.Cache.Print()

	rollback(5, &manager.Cache)
	manager.Cache.Print()

	fmt.Println("######### Add ###########")

	u1 = &pbtype.User{"khlee", "1000", "incheon"}
	manager.Cache.Add(u1)

	u1 = &pbtype.User{"khlee", "1001", "incheon"}
	manager.Cache.Add(u1)

	u1 = &pbtype.User{"khlee", "1002", "incheon"}
	manager.Cache.Add(u1)

	manager.Cache.Print()
}
