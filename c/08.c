#include<stdio.h>
int main(void)
{
    int i,j,n;
    double aver;
    while(scanf("%d",&n)!=EOF)
    {
        int sum = 0;
        for (i=0;i<n;i++)
        {
            scanf("%d",&j);
            sum+=j;
        }
       aver = sum/n ;
       printf("%.2f\n",aver);
    }
    return 0;
}
