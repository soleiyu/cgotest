package main

/* 
int cadd(int a, int b) {
	int hoge, i = 0;
	for (; i < b; i++) {
		hoge += a;
	}
	return hoge;
}
*/
import "C"
import "time"
import "fmt"

func main() {

	start := time.Now()
	println(" C :", C.cadd(11, 13))
	end := time.Now()
	fmt.Printf("%f\n", (end.Sub(start)).Seconds())

	start = time.Now()
	println("Go :", gadd(11, 13))
	end = time.Now()
	fmt.Printf("%f\n", (end.Sub(start)).Seconds())
}

func gadd(a, b int) int {
	hoge := 0
	for i:= 0; i < b; i++ {
		hoge += a
	}
	return hoge
}

