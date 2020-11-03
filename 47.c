#include<stdio.h>
#include<stdlib.h>

typedef struct t {
    int number;
    int score;
}a;

int main () {

    int n;
    scanf("%d", &n);
    a num[n];
    int max = 0;
    a ture_sum = {0,0};

    for (int i = 0; i < n; i ++ ) {
        scanf("%d %d", &num[i].number, &num[i].score);
        if (max < num[i].number)
            max = num[i].number;
    }

    a sum[1000] = {0,0};

    for (int i = 0; i < n; i ++ ) {
        sum[num[i].number].score += num[i].score;
        sum[num[i].number].number = num[i].number;
    }

    for (int i = 0; i < max; i ++ ) {
        if (ture_sum.score < sum[i].score) {
            ture_sum.score = sum[i].score;
            ture_sum.number = sum[i].number;
        }
    }

    printf("%d %d", ture_sum.number, ture_sum.score);

    return 0;
}
