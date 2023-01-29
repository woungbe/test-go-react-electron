package main

import (
	"fmt"
	"log"

	"github.com/asticode/go-astikit"
	"github.com/asticode/go-astilectron"

	"woungbe.go-react-electron/feeHttp"
)

// 생각나는 모든 조각을 테스트 해본다.
// golang 으로 고 스탑을 만들어 본다 - 전체 시작, 종료 말이지..

// http 로 데이터 통신을 해본다.

func main() {

	feeHttp.Init()
	// electronService() // 맨 마지막에 실행해서 데이터가 종료되는 것을 막는다.

}

func electronService() {
	// Set logger
	l := log.New(log.Writer(), log.Prefix(), log.Flags())

	// Create astilectron
	a, err := astilectron.New(l, astilectron.Options{
		AppName:           "Test1",
		BaseDirectoryPath: "example",
	})
	if err != nil {
		l.Fatal(fmt.Errorf("main: creating astilectron failed: %w", err))
	}
	defer a.Close()

	// Handle signals
	a.HandleSignals()

	// Start
	if err = a.Start(); err != nil {
		l.Fatal(fmt.Errorf("main: starting astilectron failed: %w", err))
	}

	// New window
	var w *astilectron.Window
	if w, err = a.NewWindow("./html/index.html", &astilectron.WindowOptions{
		Center: astikit.BoolPtr(true),
		Height: astikit.IntPtr(700),
		Width:  astikit.IntPtr(700),
		Title:  astikit.StrPtr("웅비의 운영채널"),
	}); err != nil {
		l.Fatal(fmt.Errorf("main: new window failed: %w", err))
	}

	// Create windows
	if err = w.Create(); err != nil {
		l.Fatal(fmt.Errorf("main: creating window failed: %w", err))
	}

	// Blocking pattern

	a.Wait()
}
