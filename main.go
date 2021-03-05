package main

import (
	"fmt"

	. "sqlingo_study/generated/sqlingo"

	_ "github.com/go-sql-driver/mysql"
	sqlingodrive "github.com/lqs/sqlingo"
)

func main() {
	db, err := sqlingodrive.Open("mysql", "root@/test")
	if err != nil {
		panic(err)
	}
	var posts []*PostsModel
	db.SelectFrom(Posts).
		Where(Posts.Id.In(1, 2)).
		OrderBy(Posts.FirstName.Desc()).
		FetchAll(&posts)
	for _, v := range posts {
		fmt.Println(*v)
	}
}
