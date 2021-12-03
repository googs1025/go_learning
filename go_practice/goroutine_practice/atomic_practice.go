package main

/*
	代码中的加锁操作涉及到内核态与上下文切换会比较耗较高。
	原子操作是Go语言提供的方法它在用户态即可完成，具体性能比加锁性能好。
	sync/atomic
 */

func main() {

}
