CFLAGS = -Wall -Werror -fno-strict-aliasing -fvisibility=hidden
CC=cc

TARGETS=go_test.so

all: $(TARGETS)

go_test.so: go_test.o
	$(CC) $(CFLAGS) $(LDFLAGS)-o $@ $^ -shared

clean:
	rm -rf *.c~ *.o *.a $(TARGETS)

build:
	gcc -c go_test.c && ar cru libgo_test.a go_test.o