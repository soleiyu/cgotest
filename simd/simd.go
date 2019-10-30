package main

/* 
#cgo CFLAGS: -mavx2
#include <stdio.h>
#include <immintrin.h>

void test() {
	printf("hoge\n");
}

__m256i i8x32add(__m256i a, __m256i b) {
	return _mm256_add_epu8(a, b);
}
*/
import "C"

func main() {

	C.test()
}


