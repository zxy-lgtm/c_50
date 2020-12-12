#include<stdio.h>
#include<stdlib.h>

int main () {

    int n, m;
    
    scanf("%d %d", &n, &m);

    int score[m];
    int answer[m];
    int students[n][m];

    for (int i = 0; i < m; i ++ ) {
        scanf("%d", &score[i]);
    }

    for (int i = 0; i < m; i ++ ) {
        scanf("%d", &answer[i]);
    }

    for (int i = 0; i < n; i ++ ) {
        for (int j = 0; j < m; j ++ ) {
            scanf("%d", &students[i][j]);
        }
    }

    int sum[n];

    for (int i = 0; i < n; i ++ ) {

        sum[i] = 0;

        for (int j = 0; j < m; j ++ ) {
            if(students[i][j] == answer[j]) {
                sum[i] += score[j];
            }
        }

    }

    for (int i = 0; i < n; i ++ ) {
        printf("%d\n", sum[i]);
    }

    return 0;

}