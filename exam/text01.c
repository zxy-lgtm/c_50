#include<stdio.h>
#include<stdlib.h>
//int的最大值,我用了计算器...
#define MAX 2147483647

int MyAtoi(const char* c) {
    
    //如果字符串开头有空格,就一直跳过
    while (*c == ' ') {
        c ++ ;
    }

    int tool = 1;
    //如果字符串开头空格过后不是数字
    if (*c < '0' || *c > '9') {
        //字符串的第一个是正还是负,默认他是正数(tool = 1)
        if (*c == '-') {
        tool = -1;
        c ++ ;
        } else {
            return 0;
        }
    } 

    //将字符串的数变成整数,因为考虑溢出问题,我不妨直接用long型整数
    long sum = 0;
    //定义一个工具人来记录位数,最多就十位多了就肯定溢出了
    int count = 0;

    for ( ; *c >= '0' && *c <= '9' && count < 11; c ++ ) {
        
        sum = sum * 10 + (*c - '0');
        //判断溢出,因为溢出的话,这个值就变成负的了,和原来不一样所以直接返回.
        if ((int) sum != sum) {
            return -1;
        }
        count ++ ;

    }

    sum = (int)sum * tool;

    return (int)sum;
}

int main () {
    
    char c[100];

    scanf("%s", c);

    int number1 = atoi(c);
    int number2 = MyAtoi(c);

    printf("%d %d\n", number1, number2);
    
    return 0;
}