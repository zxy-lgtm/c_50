#include<stdio.h>

int main () {

    int n;
    int a[9] = {1,2,3,4,5,6,7,8,9};
    int s, b, h;

    scanf("%d", &n);
    
    b = n/100;
    s = (n % 100) / 10;
    h = (n % 100) % 10;

    for(int i = 0; i < b; i++ ) {
        printf("B");
    }

    for(int i = 0; i < s; i++ ) {
        printf("S");
    }

    for(int i = 0; i < h; i++ ) {
        printf("%d", a[i]);
    }

    return 0;

}