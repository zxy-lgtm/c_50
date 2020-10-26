#include<stdio.h>

#define CLK_TCK 100

int main () {

    int C1, C2;
    //double型是为了四舍五入
    double m = 0;

    printf("输入C1和C2吧!\n注意C1<C2哦");
    scanf("%d %d", &C1, &C2);

    //题目规定了C1<C2
    m = (C2 - C1); 

    //+0.5是为了四舍五入
    int sum_second = (int)((m / CLK_TCK) + 0.5);
    int hour = sum_second / 3600;
    int min = (sum_second % 3600) / 60;
    int second = (sum_second % 3600) % 60;

    printf("%d:%d:%d", hour, min, second);

    return 0;
    
}