#include<stdio.h>
int main(void)
{
    int i,j;
    int n;
    while(scanf("%d",&n)!=EOF)
    {
        int sum=1;
        for (i=0;i<n;i++)
        {
            scanf("%d",&j);
            if(j%2==0)
                sum*=1;
            else 
                sum*=j;
        }
        printf("%d",sum);
    }
    return 0;
}