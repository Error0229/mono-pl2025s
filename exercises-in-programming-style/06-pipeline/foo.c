#include<stdio.h>
// dangling pointer in c
// can be solved by escape and garbage collection in Go
int* foo() {
    int x = 1;
    return &x;
}

int main() {
    int* p = foo();
    printf("%d\n", *p);
    return 0;
}