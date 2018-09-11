#include <stdio.h>
#include <stdlib.h>

int x = 0x1234;
int y;
char *s = "x = %x, y = %x\n";

int main()
{
    printf(s,x,y);
    return 0;
}