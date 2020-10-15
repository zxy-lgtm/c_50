#include <stdio.h>
int sp(int a,int b)
{
    if(a%b==0)
    return b;
    else
    sp(b,a%b);
}
int main()
{
    int c,d;
    printf("请输入一个数"); scanf("%d",&c);
    printf("请输入一个数"); scanf("%d",&d);
    printf("它们的最大公约数是:%d\n",sp(c,d));
    return 0;
}