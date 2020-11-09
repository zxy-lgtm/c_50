#include <stdio.h>
int main(void) {

    int n;
    int count = 0;
    
    do {
        printf("请输入一个数:");
        scanf("%d", &n);
        if (n > 1000)
            printf("请输入一个小于1000的数");
        //else
            //printf("%d\n", n);
    } while (n > 1000);

    do {
        if (n % 2 == 0) {
            n = n / 2;
            count++;
        } else {
            n = 3 * n + 1;
            count++;
        }
    } while (n != 1);

    printf("%d\n", count);
    
    return 0;
}