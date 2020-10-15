#include<stdio.h>
int sp(int n)
{
    int count = 0;
    for (int j=2;j<n;j++){
        if (n%j==0)
            count ++;
    } 
    if (count==0)
    return n; 
    else
    {
        return 0;
    }
    

}
int main(void)
{
    int n=100;
    int count=0;
    for (int i = 0;i < 99;i++){
        if (sp(n)!=0){
            printf("%d ",sp(n));
            count ++;
        }
        if (count%5==0)
        printf("\n");
        n --;
    }
    return 0;
}