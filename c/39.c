#include<stdio.h>

int need (int sum) {

    int count = 0;
    int money[7] = {100,50,20,10,5,2,1};

    for(int i = 0; i < 7; i ++ ) {
        count += sum / money[i];
        sum = sum % money[i];
    }

    return count;
}

int main () {

    int n;
    int sum = 0;
    int number;

    scanf("%d", &n);

    while (number != 0) {
        scanf("%d", &number);
        sum += need(number);
    }

    printf("%d", sum);

    return 0;
}