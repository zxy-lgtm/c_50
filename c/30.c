#include<stdio.h>
#include<string.h>

int sum_ (char a[], int num, int len) {

    int sum = 0;
    int number = 1;

    for (int i = 0; i < len ; i++)
    {
        if((a[i] - '0') == num ) {

            sum += number * num;
            number *= 10; 
        }
    }

    return sum;
}

int main () {

    char a1[1000], a2[1000], len1, len2;
    int num1, num2;
    int T = 2;

    scanf("%s", a1);
    scanf("%s", a2);
    scanf("%d", &num1);
    scanf("%d", &num2);

    len1 = strlen(a1);
    len2 = strlen(a2);

    int sum1 = sum_(a1, num1, len1);
    int sum2 = sum_(a2, num2, len2);

    int true_sum = sum1 + sum2; 

    printf("%d", true_sum);

    return 0;
}