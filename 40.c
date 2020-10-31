#include<stdio.h>

typedef struct person {
    int n1;
    int result;
    int money;
    int n2;
}a;

int judge (int n1, int n2) {
    
    if (n1 > n2)
        return 0;
    else
        return 1;
    
}

int main () {

    int your_money, n;

    scanf("%d", &your_money);
    scanf("%d", &n);

    for (int i = 0; i < n; i ++ ) {

        a chr;
        
        scanf("%d %d %d %d", &chr.n1, &chr.result, &chr.money, &chr.n2);
        if (chr.money <= your_money) {
            if (chr.result == judge(chr.n1, chr.n2)) {
                your_money += chr.money;
                printf("Win! %d, Total = %d\n", chr.money, your_money);
            } else {
                your_money -= chr.money;
                printf("Lose %d, Total = %d\n", chr.money, your_money);
            }    
        } else if (your_money == 0) {
            printf("Game Over!\n");
            break;
        } else {
            printf("Not enough token! Total = %d\n", your_money);
        }
    }

    return 0;
}