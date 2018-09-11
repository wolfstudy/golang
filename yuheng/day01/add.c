#include <stdio.h>
#include <stdlib.h>

int add(int x,int y)
{
    return x+y;
}

int main(int argc, char **argv)
{
    int x = 0x100;
    int y = 0x200;
    int z = add(x,y);
    printf("%d\n",z);
    return 0;
}

//gcc -O0 -o test1 add.c //gcc 的开关 把优化都关掉
//gcc -O2 -o test2 add.c // O1 O2 O3 各种优化级别