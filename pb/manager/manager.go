package manager

import (
	"encoding/json"
	"fmt"
	"goproject/demo/pb/pbtype"
	"io"
	"os"
)

func Add(m *pbtype.PhoneBook, rm *RollbackMgr) {
	var user pbtype.User
	fmt.Println("------------------------------------------")
	fmt.Println("1. 이름")
	fmt.Scan(&user.Name)
	fmt.Println("2. 전화번호")
	fmt.Scan(&user.Number)
	fmt.Println("3. 주소")
	fmt.Scan(&user.Address)
	fmt.Println("------------------------------------------")

	m.Data = append(m.Data, user)
	(*m).Count++

	// fmt.Println("aaaaa")
	// clip
	(*rm).Add(&user)
}

func Del(m *pbtype.PhoneBook) bool {
	var delName string
	fmt.Println("------------------------------------------")
	fmt.Println("1. 삭제할 이름")
	fmt.Scan(&delName)
	fmt.Println("------------------------------------------")

	r, em, d := search(m, delName)
	fmt.Println("result :", r, " , message : ", em, " , index : ", d)

	if r {
		for i := d; i < (*m).Count; i++ {
			next := i + 1
			if next == (*m).Count {
				break
			} else {
				(*m).Data[i] = (*m).Data[next]
			}
		}
		(*m).Count--
		if (*m).Count > 0 {
			m.Data = m.Data[:m.Count]
		}

	}

	return r
}

func Edit(m *pbtype.PhoneBook) {
	if m.Count == 0 {
		fmt.Println("전화부에 저장된 정보가 없습니다")
		return
	}
	for {
		fmt.Println("========================================")
		fmt.Println("편집대상 찾기 , 이름을 입력하세요")
		var searchName string
		fmt.Scan(&searchName)
		fmt.Println("========================================")

		// var findIdx int = -1

		r, _, findIdx := search(m, searchName)

		if r {
			findU := &m.Data[findIdx]
			fmt.Println(findU.ToString())

			fmt.Println("========================================")
			var cNumber, cAddress string
			fmt.Println("change Number")
			fmt.Scan(&cNumber)

			fmt.Println("change Address")
			fmt.Scan(&cAddress)
			findU.Edit(cNumber, cAddress)
			fmt.Println("========================================")
		} else {
			fmt.Println(searchName, "는 저장된 사용자가 아닙니다.")
		}

		fmt.Println("========================================")
		fmt.Println("Quit?(Y,N)")
		var quitYn string
		fmt.Scan(&quitYn)
		fmt.Println("========================================")
		if quitYn == "Y" {
			break
		}

	}

}

func Display(m *pbtype.PhoneBook, rm *RollbackMgr) {

	if (*m).Count == 0 {
		fmt.Println("데이타가 없습니다")
	} else {
		for i := 0; i < (*m).Count; i++ {
			fmt.Printf("[%d] : Name : %s , Number: %s , Address : %s \n", i+1, (*m).Data[i].Name, (*m).Data[i].Number, (*m).Data[i].Address)
		}

		fmt.Println("--------- Clip Board -------------")
		(*rm).Print()
	}
}

func Search(m *pbtype.PhoneBook) {
	var cnt int = (*m).Count
	if cnt == 0 {
		fmt.Println("저장된 사용자가 없습니다")
	} else {

		for {
			var searchName string
			fmt.Println("------------------------------------------")
			fmt.Println("1. 검색을 이름을 입력하세요( 종료는 quit )")
			fmt.Scan(&searchName)
			fmt.Println("------------------------------------------")

			if searchName == "quit" {
				break
			}

			r, _, d := search(m, searchName)
			if r {
				u := (*m).Data[d]
				fmt.Printf("[%d]Name :%s , Number : %s , Addr : %s ", d, u.Name, u.Number, u.Address)
			}
		}

		fmt.Println("------------ clip info ----------------")
		Cache.Print()

	}
}

func search(m *pbtype.PhoneBook, searchName string) (bool, string, int) {

	// sort
	

	if (*m).Count == 0 {
		return false, "저장된 사용자가 없습니다", -1
	} else {
		fmt.Println("=== Count :", (*m).Count, " -- len :", len((*m).Data))
		for i := 0; i < (*m).Count; i++ {
			if (*m).Data[i].Name == searchName {
				return true, "", i
			}
		}

		return false, "입력한 사용자 정보를 찾을수 없습니다.", -1
	}
}

func ReadFile(path string, m *pbtype.PhoneBook) {
	if f, err := os.Open(path); err == nil {
		defer f.Close()

		content, e1 := io.ReadAll(f)
		if e1 != nil {
			return
		}

		e2 := json.Unmarshal(content, &m.Data)
		if e2 != nil {
			panic(e2)
		}
		m.Count = len(m.Data)
	}

}

func SaveFile(path string, m *pbtype.PhoneBook) {
	if m.Count > 0 {
		doc, _ := json.Marshal(m.Data)
		if doc != nil {
			// 기존 파일삭제
			if f, err := os.Stat(path); err == nil {
				os.Remove(f.Name())
				// defer f.Close()
			}
			//새로운 정보 저장
			os.WriteFile(path, doc, 0777)

		}
	}

}

func PbRollback(m *pbtype.PhoneBook, rm *RollbackMgr) error {

	rmCount := (*rm).Count
	if rmCount == 0 {
		return fmt.Errorf("클립보드가 비었습니다.")
	}

	if ru, err := (*rm).Rollback(); err == nil {
		fmt.Println("복구될 정보 >>>", ru.ToString())

		m.Data = append(m.Data, *ru)
		(*m).Count++
		return nil
	} else {
		return fmt.Errorf("복구에 실패하였습니다")
	}

}

//--------------------------------------------------------------------

type Node struct {
	Data *pbtype.User
	Next *Node
	Prev *Node
}

type RollbackMgr struct {
	Count int
	CNode Node
}

var Cache RollbackMgr = RollbackMgr{0, Node{nil, nil, nil}}

// func GetCache() *RollbackMgr {
// 	return &Cache
// }

func (rm *RollbackMgr) Initialzie() {
	(*rm).Count = 0
	(*rm).CNode = Node{nil, nil, nil}
}

func (rm *RollbackMgr) Add(u *pbtype.User) {
	fmt.Println(">>> Add >>", (*u).ToString())

	count := (*rm).Count

	switch count {
	case 0:
		(*rm).CNode = Node{u, nil, nil}
		(*rm).Count++
	case 10:
		//가장오래된 데이타 삭제...
		fNode := firstNode(rm)
		(*rm).CNode = *fNode.Next
		(*rm).CNode.Prev = nil

		fNode = &Node{}
		// fNode.Next = nil
		// fNode.Prev = nil

		//새로운 데이타 마지막에 추가
		lNode := lastNode(rm)
		nNode := &Node{u, nil, lNode}
		lNode.Next = nNode

		// (*rm).Count++
	default:
		//새로운 데이타 마지막에 추가
		lNode := lastNode(rm)
		nNode := &Node{u, nil, lNode}
		lNode.Next = nNode

		(*rm).Count++
	}

	// fmt.Println("count >>", (*rm).Count)
}

func (rm *RollbackMgr) Rollback() (*pbtype.User, error) {

	if (*rm).Count == 0 {
		return nil, fmt.Errorf("01")
	}

	lNode := lastNode(rm)
	// dNode := Node{lNode.Data, nil, nil}
	lNode.Prev.Next = nil

	(*rm).Count--
	// fmt.Println(lNode.Data)
	return lNode.Data, nil

}

func (rm *RollbackMgr) Print() {
	if (*rm).Count > 0 {
		tNode := &(*rm).CNode
		// fmt.Println("Print>>", tNode.Data.ToString())
		// fmt.Println("Print2>>", tNode.Next)
		for {
			if tNode.Next == nil {
				fmt.Println(tNode.Data.ToString())
				return
			}

			fmt.Println(tNode.Data.ToString())

			tNode = tNode.Next
		}
	}
}

func lastNode(rm *RollbackMgr) *Node {
	fNode := &rm.CNode
	for {
		if fNode.Next == nil {
			return fNode
		} else {
			fNode = fNode.Next
		}
	}
}

func firstNode(rm *RollbackMgr) *Node {
	fNode := &rm.CNode
	for {

		if fNode.Prev == nil {
			// fmt.Println("firstNode()>>>", fNode.Data.ToString())
			return fNode
		} else {
			fNode = fNode.Prev
		}
	}
}
