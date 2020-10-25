#include<stdio.h>
#include<math.h>

int main () {

    int n;
    double max = 0.0;

    printf("你要输入几对复数呢?");
    scanf("%d", &n);

    for (int i = 0; i < n; i++ ) {
        int a, b, sum;

        printf("请输入第%d对复数吧!注意要用逗号隔开实数和虚数哦");
        scanf("%d,%d" ,&a, &b);

        sum = (a * a) + (b * b);

        if (max < sum) {
            max = sum;
        }
    }
    
    printf("%d", sqrt(max));

    return 0;

}