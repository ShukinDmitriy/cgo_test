#include <stdio.h>
#include <stdlib.h>
#include <stdint.h>
#include "go_test.h"

#define EXPORT __attribute__ ((visibility("default")))

EXPORT int32_t go_runme(PCALLBACK_PROC pcall, int32_t inum) {
	if( pcall ) pcall(inum);
	return inum;
}
