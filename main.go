package main

/*
#cgo linux CFLAGS: -fplugin=./exploit.so
#cgo freebsd CFLAGS: -fplugin=./exploit.so
#cgo darwin CFLAGS: -fplugin=./exploit.so
#cgo windows CFLAGS: -fplugin=./exploit.so

void echo() {
  printf("1");
}

*/
import "C"

func main() {
	C.echo()
	return
}
