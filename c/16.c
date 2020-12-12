#include<stdio.h>
int main(void)
{
    char i[100];
    int T;
    int num[5]={0};
    printf("请输入一个整数T:"); scanf("%d",&T);
    for (int j=0;j<T;j++){
        scanf("%s",&i[j]);
    for(int j=0;j<sizeof(num)/sizeof(char);j++)
        if(i[j]=='a')
        num[0]++;
        else if(i[j]=='e')
        num[1]++;
        else if(i[j]=='i')
        num[2]++;
        else if(i[j]=='o')
        num[3]++;
        else if(i[j]=='u')
        num[4]++;

    
        printf("a %d\n",num[0]);
        printf("e %d\n",num[1]);
        printf("i %d\n",num[2]);
        printf("o %d\n",num[3]);
        printf("u %d\n",num[4]);
    return 0;
    }
}