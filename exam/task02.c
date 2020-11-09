#include<stdio.h>

int main () {

    int weight;
    double money = 0.00;
    printf("请输入你要邮寄的物品的重量吧\n");
    scanf("%d", &weight);

    if (weight > 30) {
        printf("fail\n");
    } else {
        if (weight <= 10) {
            money = (double)weight * (0.80) + 0.20;
        } else if (weight <= 20 && weight > 10) {
            money = (double)(weight) * (0.75) + 0.20;
        } else {
            money = (double)(weight) * (0.70) + 0.20;
        }

        printf("一共需要%.2f元\n", money);
    }
}