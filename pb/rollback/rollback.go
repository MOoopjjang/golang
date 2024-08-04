package rollback

import (
	"fmt"
	"goproject/demo/pb/pbtype"
	"goproject/demo/pb/syncmanager"
)

type Node struct {
	Data *pbtype.User
	Next *Node
	Prev *Node
}

type RollbackMgr struct {
	Count int
	CNode Node
}

var cache RollbackMgr = RollbackMgr{0, Node{nil, nil, nil}}

func Cache() *RollbackMgr {
	return &cache
}

/*
func (rm *RollbackMgr) Initialzie() {
	(*rm).Count = 0
	(*rm).CNode = Node{nil, nil, nil}
}
*/

func Initialzie() {
	cache.Count = 0
	cache.CNode = Node{nil, nil, nil}
}

/**
*  go-routine 메소드
*
 */
func Add() {

	for u := range syncmanager.Sm().Ch() {
		fmt.Println(">>> Add >>", (*u).ToString())

		switch cache.Count {
		case 0:
			cache.CNode = Node{u, nil, nil}
			cache.Count++
		case 10:
			//가장오래된 데이타 삭제...
			fNode := firstNode(&cache)
			cache.CNode = *fNode.Next
			cache.CNode.Prev = nil

			fNode = &Node{}
			// fNode.Next = nil
			// fNode.Prev = nil

			//새로운 데이타 마지막에 추가
			lNode := lastNode(&cache)
			nNode := &Node{u, nil, lNode}
			lNode.Next = nNode
		default:
			//새로운 데이타 마지막에 추가
			lNode := lastNode(&cache)
			nNode := &Node{u, nil, lNode}
			lNode.Next = nNode

			cache.Count++
		}

	}

	/* go-routine done */
	syncmanager.Sm().Done()
}

func (rm *RollbackMgr) Rollback() (*pbtype.User, error) {

	if (*rm).Count == 0 {
		return nil, fmt.Errorf("01")
	}

	lNode := lastNode(rm)
	fmt.Printf("%p \n", lNode)

	if lNode.Prev != nil {
		lNode.Prev.Next = nil
	}

	(*rm).Count--
	// fmt.Println(lNode.Data)
	return lNode.Data, nil

}

func (rm *RollbackMgr) Print() {
	if (*rm).Count > 0 {
		tNode := &(*rm).CNode
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
