#include<stdio.h>

int main () {
    
    int n;
    int a[1000] = {0};
    a[0] = 0;
    a[1] = 2;
    a[2] = 2;
    a[3] = 2;

    scanf("%d", &n);

    //判断是否可以跳回去
    if (n < 2) {
        printf("不能跳回去哦");
    } else {

        //创建一个循环,因为前三个数字已经确定啦,所以直接从4开始循环递归
        for (int i = 4; i <= 1000; i++ ) {
            a[i] = ((2 * a[i - 2]) + a[i - 1]);
        }

        printf("一共有%d种方法哦", a[n]);
    }

    return 0;

}