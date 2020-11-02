#include<stdio.h>

int main () {

    int  n, m, min, max, stand;
    scanf("%d %d %d %d %d", &n, &m, &min, &max, &stand);
    int num[n][m];

    for (int i = 0; i < n; i ++ ) {
        for (int j = 0; j < m; j ++ ) {
            scanf("%d ", &num[i][j]);
            if (num[i][j] >= min && num[i][j] <= max) {
                num[i][j] = stand * 111;
            }
        }
    }
 
    for (int i = 0; i < n; i ++ ) {
        for (int j = 0; j < m; j ++ ) {
            printf("%03d ", num[i][j]);
        }
        printf("\n");
    }

    return 0;


}