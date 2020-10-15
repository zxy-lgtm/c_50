#include<stdio.h>
int sp(int n)
{
    int count = 0;
    int c;
    for (int j=2;j<n;j++){
        if (n%j==0)
            count +=1;
    }
    if (count==0)
        c = n;
    return c;
}
int main(void)
{
    int n=100;
    for (int i = 0;i < 100;i++){
        printf("%d ",sp(n));
        n--;
        if (n%5==0)
        printf("\n");
    }
    return 0;
}