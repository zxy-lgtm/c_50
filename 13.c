#include<stdio.h>
int main(void)
{
    int T;
    int i;
    do{
        scanf("%d",&T);
        if(T>10)
        printf("请输入一个小于等于10的数\n");
    }while(T>10);
    for (i=0;i<T;i++){
       long int A,B,C;
        printf("请以number number number的形式输出");
        scanf("%d %d %d",&A,&B,&C);
        int D=A+B;
        if(D>C)
        printf("Case #x: ture\n");
        else
        printf("Case #x: false\n");
    }
    return 0;
}