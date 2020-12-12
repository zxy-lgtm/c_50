#include<stdio.h>

int prime_number(int n) {
    
    int count = 0;

    for (int i = 2; i < n; i++) {
        if (n % i == 0) {
            count ++;
        }
    }

    if (count == 0) {
        return n;
    } else {
        return 0;
    }
}

int main () {

    int a, b;
    int min, max;

    printf("请输入一个数吧:");
    scanf("%d", &a);
    printf("请再输入一个数吧:");
    scanf("%d", &b);

    //比较a, b的大小
    min = a;
    max = a;
    if (min < b)
        max = b;
    else
        min = b;

    //自己定义一个足够长的素数数组.
    int c[10000] = {0};
    int count = 0;

    for (int i = 0; i < 10000 ; i++) {
        if (prime_number(i) != 0) {
            c[count] = prime_number(i);
            count ++;
        }    
    }

    //定义一个换行工具人
    int u = 0;

    for (int i = 0; i < max ; i++) {
        if (i >= min) {
            printf("%d ", c[i]);
            u ++;
        }
        if (u % 10 == 0) {
            printf("\n");
        }
    }
    
    return 0;

}