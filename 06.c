#include <stdio.h>
#include <math.h>
#define NUMBER 100
int main()
{
    int k;
    while (scanf("%d",&k)==1 && k!= 0)
    {
        int a[NUMBER];
        int i=0;
        while (k--)
        {
            scanf("%d",&a[i++]);
        }
        for (int j=0;j<i;j++)
        {
            for(int w=j;w<i;w++)
            {
                if(fabs(a[i])>fabs(a[j]))
                {
                    int s=a[i];
                    a[i]=a[j];
                    a[j]=s;
                }
            }
        }
        for (int j=0;j<i;j++)
        {
            printf("%d",a[j]);
        }
        printf("\n");
    }
    return 0;
}