#include<stdio.h>

//求出从第一项到第n项的多项式的和;
double sum_ (int n) {

    int count = 2;
    double sum = 0.00;

    for (double i = 1; i <= n; i++ ) {
        if (count % 2 == 0) {
            sum += (1.00 / i);
        } else {
            sum -= (1.00 / i);
        }

        //让奇数项加,偶数项减
        count ++;
    }

    return sum;

}

int main () {

    int n;

    printf("你要输入几个数呢?");
    scanf("%d", &n);

    int a;
    double sum;
    char ch;

    printf("现在输入你的每个数吧!\n如果想结束的话就按下ctrl和D吧!");

    while (ch = getchar() != EOF) {
        scanf ("%d", &a);
        sum = sum_(a);
        printf("%.2f",sum);
    }

    return 0;

}