package main

import (
	"fmt"
	"goproject/demo/pb/manager"
	"goproject/demo/pb/pbtype"

	"goproject/demo/pb/rollback"
	"goproject/demo/pb/syncmanager"
)

func main() {
	var pbm pbtype.PhoneBook
	// slice type으로 생성
	pbm.Data = make([]pbtype.User, 100, 100)

	rollback.Initialzie()

	/** go */
	syncmanager.Sm().Initialzie()
	syncmanager.Sm().Add(1)
	go rollback.Add()

	manager.ReadFile("data.json", &pbm)
	for {
		fmt.Println("------------------------------------------")
		fmt.Println("1. 추가")
		fmt.Println("2. 제거")
		fmt.Println("3. 수정")
		fmt.Println("4. 찾기")
		fmt.Println("5. 출력")
		fmt.Println("6. 되돌리기")
		fmt.Println("7. 종료")
		fmt.Println("------------------------------------------")
		var nSelect int
		n, err := fmt.Scanf("%d", &nSelect)
		if err != nil {
			fmt.Println(n, err)
		} else {
			var isExist bool = false

			switch nSelect {
			case 1:
				// manager.Add(&pbm, &manager.Cache)
				manager.Add(&pbm)
			case 2:
				manager.Del(&pbm)
			case 3:
				manager.Edit(&pbm)
			case 4:
				manager.Search(&pbm)
			case 5:
				manager.Display(&pbm)
			case 6:
				manager.PbRollback(&pbm)
			default:
				fmt.Println("Program end")
				manager.SaveFile("data.json", &pbm)
				isExist = true

			}

			if isExist {
				break
			}
		}
	}

	syncmanager.Sm().CloseAndWait()
}
