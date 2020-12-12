#include<stdio.h>
int m(int n)
{
    int month2;
    if(n%4==0)
        {
            month2=29;
        }
    else
        {
            month2=28;
        }
    return month2;
}
int d(int year,int month,int day)
{
    int count;
    int sum=0;
       int sum2=m(year)+31+day;
            for (count=2;count<month-1;count++)
            {
                if((month%2==0 && month<8) || (month%2==1 && month>8))
                {
                    sum+=30;
                }
                else
                {
                    sum+=31;
                }
            }
            sum=sum+sum2;
        return sum;
}
int main(void)
{
    int sum;
    int year,month,day;
    printf("请以year/month/day的形式输出");
    while(scanf("%d/%d/%d",&year,&month,&day)!=EOF)
    {
       
        if(month==1)
        {
            int sum1=day;
            printf("%d\n",sum1);
        }
        else if(month==2)
        {
            int sum1=31+day;
            printf("%d\n",sum1);
        }
        else
        {
          printf("%d\n",d(year,month,day));  
        }
    }
    return 0;
}
