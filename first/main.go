package main

import (
	"database/sql"
	"github.com/pkg/errors"
	"log"
	"net/http"
)

const DataNotFound = 10001 // 数据不存在

type MyError struct {
	Message string
	Code    int
	Err     error
}

func GetSql() error {
	return errors.Wrap(sql.ErrNoRows, "data not found")
	//return errors.WithStack(sql.ErrNoRows)
}

func Call() error {
	return errors.WithMessage(GetSql(), "data not found")
}

func GetSqlNoErr() MyError {
	// 伪造错误
	err := sql.ErrNoRows
	myMsg := MyError{
		Err:     nil,
		Message: "ok",
		Code:    http.StatusOK,
	}
	if errors.Is(err, sql.ErrNoRows) {
		myMsg.Err = err
		myMsg.Code = DataNotFound
		myMsg.Message = "data not found"
	}
	return myMsg
}

func main() {
	//err := Call()
	//if errors.Cause(err) == sql.ErrNoRows {
	//	// 打印堆栈信息
	//	fmt.Printf("data not found, %v\n", err)
	//	fmt.Printf("%+v\n", err)
	//	return
	//}
	//if err != nil {
	//	// unknown error
	//}
	err := GetSqlNoErr()
	if err.Code == DataNotFound {
		log.Fatal(err.Message)
	}
}