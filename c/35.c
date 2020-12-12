#include<stdio.h>

int main () {

    int t, num[100] = {0};
    double cases[][100] = {0}, max[100] = {0};

    scanf("%d", &t);

    for (int i = 0; i < t; i++ ) {
        scanf("%d", &num[i]);
        for (int j = 0; j < num[i]; j++ ) {
            scanf("%lf", &cases[i][j]);
            max[i] = cases[i][0];
            if(max[i] < cases[i][j]) {
                max[i] = cases[i][j];
            } 
        }
    }

    for (int i = 0; i < t; i++ ) {
        printf("%.2f\n", max[i]);
    }

    return 0;

}