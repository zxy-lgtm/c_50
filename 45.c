#include<stdio.h>
#include<string.h>

int main () {

    char p[100][100];
    char ch;
    do {

        int i = 0;
        int a[3] = {0};
        scanf("%s", p[i]);
        int len = sizeof(p[i]);
        char chll[len + 1];
        char *q = chll;
        char *head = chll;

        for (int j = 0; j < len; j ++ ) {
            if (p[i][j] == 'Z') {
                a[0] ++ ;
            } else if (p[i][j] == 'O') {
                a[1] ++ ;
            } else if (p[i][j] == 'J') {
                a[2] ++ ;
            }
        }

        for (int j = 0; j < len + 1; j ++ ) {
            if (a[0] > 0 && a[1] > 0 && a[2] > 0) {
                printf("ZOJ");
            } else if (a[0] <= 0 && a[1] > 0 && a[2] > 0) {
                printf("OJ"); 
            } else if (a[0] > 0 && a[1] <= 0 && a[2] > 0) {
                printf("ZJ");  
            } else if (a[0] > 0 && a[1] > 0 && a[2] <= 0) {
                printf("ZO");
            } else if (a[0] <= 0 && a[1] <= 0 && a[2] > 0) {
                printf("J");
            } else if (a[0] <= 0 && a[1] > 0 && a[2] <= 0) {
                printf("O");
            }else if (a[0] > 0 && a[1] <= 0 && a[2] <= 0) {
                printf("Z");
            }

            a[0] --;
            a[1] --;
            a[2] --;
        }

        printf("\n");
        i ++ ;


    }while (ch = getchar() != 'E');

    return 0;
}