#include<stdio.h>
int money[100000]={0};
int swap(int *pa,int *pb)
{
    int t = *pa;
    *pa = *pb;
    *pb = t;
    return;
}
void quicksort(int a,int b)
{
    if (a>b)
    return;
    int i,j,t,temp;
    i = a;
    j = b;
    temp = money[a];
    while(i!=j)
    {
        while(money[j]<=temp && i<j)
        j--;
        while(money[i]>=temp && i<j)
        i++;
        if(i<j)
        swap(&money[i],&money[j]);
    }
    money[a]=money[i];
    money[i]=temp;
    quicksort(a,i-1);
    quicksort(i+1,b);
    return;
}
int main(void)
{
    int m=0,n=0;
    while(1)
    {
        printf("请以number number的形式输出");
        scanf("%d %d",&n,&m);
        if(m==0 && n==0)
        break;
        for(int i=0;i<n;i++)
        scanf("%d",&money[i]);
        quicksort(0,n);
        for(int i=0;i<m;i++)
        printf("%d",money[i]);
        printf("\b\n");
    }
    return 0;
}