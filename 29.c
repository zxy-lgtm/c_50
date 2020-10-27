#include<stdio.h>

int main () {

    int number[10], N;
    int sum = 0;

    scanf("%d", &N);

    for (int i = 0; i < N; i++) {
        scanf("%d", &number[i]);
    }

    for (int i = 0; i < N; i++ ) {
        for(int j = 0; j <N; j++ ) {
            if (j != i && number[j] != number[i])
                sum += number[j] * 10 + number[i];
        }
    }

    printf("%d", sum);

    return 0;
}