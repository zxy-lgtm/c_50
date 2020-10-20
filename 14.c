#include<stdio.h>
int main(void)
{
    int n,min,count;
    int a[100]={0};
    int i,j;
    while(scanf("%d",&n) != EOF)
    {
        if(n==0)
        break;
        else
        for (i=0;i<n;i++){
            scanf("%d",&a[i]);
            min=a[0];
            count = 0;
        }
        for(i=0;i<n;i++)
        {
            if(min>a[i]){
                min=a[i];
                count=i;
            }
        }     
        a[count]=a[0];
        a[0]=min;
        for (i=0;i<n;i++){
            printf("%d",a[i]);
    }
    return 0;
}
