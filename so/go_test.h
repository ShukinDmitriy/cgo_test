#ifndef __GO_TEST_H__
#define __GO_TEST_H__

typedef int32_t (*PCALLBACK_PROC)(int32_t inum);
int32_t go_runme(PCALLBACK_PROC pcall, int32_t inum);

#endif