#include<stdio.h>

int main () {

    int N;
    int A, a, B, b;
    int suma = 0, sumb = 0;

    scanf("%d", &N);

    for(int i = 0; i < N; i ++ ) {
        scanf("%d %d %d %d", A, a, B, b);
        if (A + B == a  && A + B != b)
            suma ++;
        else if (A + B == b  && A + B != a)
            sumb ++;
    }

    printf("%d %d", suma, sumb);

    return 0;

}