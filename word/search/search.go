package search

import (
	"fmt"
)

type SearchInfo struct {
	RootDir  string
	FindWord string
}

func (si *SearchInfo) ToStringn() string {
	return fmt.Sprintf("start dir=%s , find = %s", (*si).RootDir, (*si).FindWord)
}
