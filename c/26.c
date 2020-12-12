#include<stdio.h>

int main () {

    int a, b;
    int min = 0;

    printf("请输入一个数");
    scanf("%d", &a);
    printf("请再输入一个数吧");
    scanf("%d", &b);

    //最小公倍数一定小于或等于max
    int max = a * b;

    for (int i = 1; i <= max; i++ ) {
        if (i % a == 0 && i % b == 0) {
            min = i;
            break;
        } 
    }

    printf("它们的最小公倍数是%d", min);

    return 0;

}