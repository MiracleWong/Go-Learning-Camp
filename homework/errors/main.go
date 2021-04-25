package main

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
)

/*
	因为 github.com/pkg/errors 兼容了 Go1.13 中的errors 相关的内容，如wrap、unwrap、errors.Is 和 errors.As
	这里使用github.com/pkg/errors 作为errors 使用，而不是go 自带的errors
 */

var err error

func getUsers() error {
	// wrap 错误，并带入相关信息
	return errors.Wrap(sql.ErrNoRows, "sql: select id, name, sex from users;")
}

func test1() error {
	return getUsers()
}

func test2() error {
	return test1()
}


func main() {
	err = test2()
	if errors.Is(err, sql.ErrNoRows) {
		fmt.Printf("original err: %T %v\n", errors.Cause(err), errors.Cause(err))
		fmt.Printf("stack trace:\n%+v\n", err)
		return
	}
}


