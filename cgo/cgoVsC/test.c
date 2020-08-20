#include <stdlib.h>
#include <stdio.h>

#include "test.h"

#define LENGTH 10

void* c_malloc(int size) {
    void* ptr = malloc(size);
    return ptr;
}

void c_free(void* ptr) {
  if (ptr != NULL) {
    free(ptr);
    ptr = NULL;
  }
}

int sum(int i, int j) {
	return i + j;
}

void assign(double ptr[]) {
    printf("Hello, c\n");

	for (int i = 0; i < LENGTH; i++) {
        ptr[i] = i;
	}
}

void cptr_test(void* ptr) {
  for (int i = 0; i < LENGTH; i++) {
    ((double *)ptr)[i] = i;
  }
}

void print_ptr(double* ptr) {
    for (int i = 0; i < LENGTH; i++) {
        printf("Print from c, %f\n", ptr[i]);
    }
}

void is_null_ptr(double* ptr) {
    if (ptr == NULL) {
        printf("is null ptr!\n");
    } else {
        printf("is not null ptr!\n");
    }
}

//int main() {
//    void* p = c_malloc(sizeof(double) * LENGTH);
//    cptr_test(p);
//    print_ptr((double*)p);
//    c_free(p);
//    return 0;
//}
