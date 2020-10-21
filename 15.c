#include<stdio.h>
int main(void)
{
    int T,a[1000];
    printf("请输入一个整数T:"); scanf("%d",&T);
    printf("请输入%d个整数并以number 的形式输入",T);
    for(int i=0;i<T-1;i++){
        do{
        scanf("%d ",&a[i]);
        }while(a[i]<32);
    }
    for(int i=0;i<T;i++)
        printf("%c",a[i]);
    return 0;
}