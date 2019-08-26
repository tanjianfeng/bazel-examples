package main

/*
extern int clib_hello(void);
*/
import "C"

func main() {
    C.clib_hello()
}
