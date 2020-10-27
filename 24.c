#include<stdio.h>
#include<string.h>

int main () {

    char ch;
    int t_sum = 0;
    int sum[10] = {0};
    int num[3] = {0};
    int count = 1;

    scanf("%s", ch);
    
    while ((ch = getchar()) != '\n')
    {
        if (ch >= '0' && ch <= '9') {
            sum[ch - '0'] += (ch - '0');
            count ++;
        }
    }

    for (int i = 0; i < count; i++ ) {
        t_sum += sum[i];

    }

    

    //因为自然数小于十的一百次方,那么最多有一百位,最大和为900,所以三个就行
    num[1] = t_sum / 100;
    num[2] = (t_sum % 100) / 10;
    num[3] = (t_sum % 100) % 10;
    

    for (int i = 1; i < 4; i++) {
        switch (num[i]) {
        case 1 : printf("yi ");  break;
        case 2 : printf("er ");  break;
        case 3 : printf("san "); break;
        case 4 : printf("si ");  break;
        case 5 : printf("wu ");  break;
        case 6 : printf("liu "); break;
        case 7 : printf("qi ");  break;
        case 8 : printf("ba ");  break;
        case 9 : printf("jiu "); break;
        case 0 : printf("ling ");break;
        }
    }

    //为了让最后一格没空
    printf("\b");

    return 0;

}