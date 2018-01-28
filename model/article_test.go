package model

import (
	"testing"
	"fmt"
)


func TestArticle_AddCate(t *testing.T)  {
	fmt.Println(123)
	var m = Article{}
	var err = m.AddCate("哈哈哈",100)
	fmt.Print(err)
}
