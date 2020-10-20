#include<stdio.h>
int swap(int *pa,int *pb)
{
    int temp=*pa;
    *pa=*pb;
    *pb=temp;
    return 0;
}
int main(void)
{
    int n,min;
    int count=0;
    int a[100]={0};
    int i;
    while(scanf("%d",&n) != EOF)
    {
        for (i=0;i<n;i++){
            scanf("%d",&a[i]);
        }
        for(i=0;i<n;i++)
        {
            min=a[0];
            if(min>a[i])
                count++;
        }
        swap(&a[count],&a[0]);
        for (i=0;i<n;i++){
            printf("%d",a[i]);
    }
    return 0;
}
