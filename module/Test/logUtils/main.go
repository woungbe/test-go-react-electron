package main

import (
	"fmt"
	"time"

	"woungbe.go-react-electron/module/logUtils"
)

// defind 잡고서 전역에서 사용하게 만들면 편하긴 하겠네.  한쪽으로 모이니까.
// 관리만 잘하면 어떻게 된 되겠군...

func main() {

	logs := new(logUtils.LogModule)
	logs.Init(func(msg string) {
		fmt.Println("msg : ", msg)
	})

	done := make(chan bool)
	i := 0
	go func() {
		for {
			time.Sleep(time.Second)
			str := fmt.Sprintf("i : %d", i)
			logs.Print(str)
			i++
		}
	}()

	<-done

}
