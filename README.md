Go는 정적 컴파일 언어.

JAVA나 Python처럼 동적 컴파일언어가 대세인 상태에서 Go는 정적으로 돌아감
왜? 속도 때문에

go build를 사용하면 컴파일을 할 수 있는데
Linux에서는 exe가 사용되지 아니함. 
.exe가 없음

GOOS="windows" go build
=> 다른 플랫폼의 실행파일을 쉽게 생성할 수 있음.