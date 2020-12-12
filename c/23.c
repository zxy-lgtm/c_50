#include<stdio.h>

int main () {

    int ch;
    int sum[10] = {0};

    printf("输入你的数吧");
    // scanf("%d", &ch);

    while ((ch = getchar()) != '\n') {
        if (ch >= '0' && ch <= '9') {
            sum[ch - '0'] ++;
        }
    }

    puts("看!");
    for (int i = 0; i < 10; i++ ) {
        printf("'%d':%d\n", i, sum[i]);
    }

    return 0;

}