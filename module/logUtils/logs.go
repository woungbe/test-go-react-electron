package logUtils

// 어떤 로그가 실행되면 모와주는 것 만들기

// 설정 및 세팅
// 콜백 하나 있어야됨
// 데이터 받는 데가 있어야됨
// => 데이터를 받으면 어떻게 처리할지 리턴 주는 애가 있어야되는데 .

type sendLog func(msg string)

type LogModule struct {
	writeLogflg  bool    // 파일로 저장할꺼니 말꺼니 ?
	callBackFunc sendLog // 콜백 실행하기
}

func (ty *LogModule) Init(callBackFunc sendLog) {
	ty.callBackFunc = callBackFunc
}

func (ty *LogModule) Print(msg string) {
	ty.callBackFunc(msg)
}
