#include<stdio.h>
#include<string.h>

int main () {

    int n;
    char a[1000];
    int num[4] = {0};

    scanf("%d", &n);

    int sum[1000] = {0};

    for (int i = 0; i < n; i ++ ) {

        scanf("%s", a);
        char *head = a;

        if (strlen(a) > 7 && strlen(a) < 17) {
            while ( *head != '\0') {
                if(*head > 64 && *head < 91)
                    num[0] = 1; 
                
                else if (*head > 96 && *head < 123)
                    num[1] = 1; 
                
                else if (*head > 47 && *head < 58)
                    num[2] = 1; 
                
                else if (*head == 127 || *head == 33 || *head == 36 || *head == 37 || *head == 64 || *head ==94)
                    num[3] = 1; 

                head ++;
            }
        }
    }

    for (int i = 0; i < n; i ++ ) {
        for (int j = 0; j < 4; j ++ ) {
            sum[i] += num[j];
        }
    }
    for (int i = 0; i < n; i ++ ) {
        if (sum[i] > 2)
            printf("YES\n");
        else 
            printf("NO\n");
    }
    
    return 0;
}