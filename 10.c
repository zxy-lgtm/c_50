#include<stdio.h>
int compare(int i,int a[50][5],int b[5])
{
    int count=0;
    for(int j=0;j<i;j++){
        if(a[i][5]>b[5]){
            count++;
        }
    }
    return count;
}
int main(void)
{
    int m,n;
    int average,number;
    scanf("%d%d",&n,&m);
    int a[50][5]={0};
    int aver[5]={0};
    for (int student=0;student<n;student++){
        for (int score=0;score<m;score++){
            scanf("%d",a[50][5]);
        }
    }
    for (int i=0;i<n;i++){
        average=a [i][5]/n;
        aver[i]=average;
        printf("%d",aver[i]);
    }
    printf("%d",compare(n,a[50][5],aver[5]));
    return 0;
}