package main

/*

#include <stdint.h>
#include <stdio.h>

// The gateway function
int callOnMeGo_cgo(int32_t inum)
{
    // определение типа функции из GO
    int32_t callOnMeGo(int32_t);

	// вызов функции GO
    return callOnMeGo(inum);
}
*/
import "C"
