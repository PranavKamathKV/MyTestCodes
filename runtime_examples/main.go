package main

import(
	"fmt"
	"runtime"
)

func stackExample(){
	stackSlice:= make([]byte, 256)
	s:=runtime.Stack(stackSlice, false)
	fmt.Println(stackSlice[0:s])
}

func first(){
	second()
}

func second(){
	third()
}

func third(){
	for c:=0;c<5;c++{
		unintrptr_pc, filepath, linenumber, ok := runtime.Caller(c)
		fmt.Println(unintrptr_pc, filepath, linenumber, ok)
         	det:= runtime.FuncForPC(unintrptr_pc)
	//	fmt.Println(runtime.FuncForPC(unintrptr_pc))
		fmt.Println(det, det.Name())
	}
}

func main(){
	fmt.Println("Demo for Runtime")
	stackExample()
	fmt.Println("Demo for Caller")
	first()
}
