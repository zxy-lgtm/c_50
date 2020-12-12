#include <stdio.h>

int main () {

    int T;

    scanf("%d", &T);

    int A[T], B[T];
    int c[T];

    for (int i = 0; i < T; i++ ) {
        scanf("%d", &A[i]);
        scanf("%d", &B[i]);

        if(A[i] % B[i] == 0) {
            c[i] = 1;
        } else {
            c[i] = 0;
        }
        
    }

    for(int i = 0; i < T; i++ ) {
        if(c[i] == 1) {
            printf("Yes");
        } else {
            printf("No");
        }
    }
    return 0;
}