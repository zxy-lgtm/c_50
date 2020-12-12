#include <stdio.h>
#include<math.h>//为了使用sin(),cos()函数 在VS Code上编译需要指明路径 gcc xx -lm

int main () {

    double R1, P1, R2, P2;
    
    printf("请输入R1");
    scanf("%lf", &R1);
    printf("请输入P1");
    scanf("%lf", &P1);
    printf("请输入R2");
    scanf("%lf", &R2);
    printf("请输入P2"); 
    scanf("%lf", &P2);

    double A1 = R1 * cos(P1);
    double B1 = R1 * sin(P1);
    double A2 = R2 * cos(P2);
    double B2 = R2 * sin(P2);

    double A = (A1 * A2) - (B1 * B2);
    double B = (A1 * B2) + (A2 * B1);

    printf("%.2f%.2fi", A, B);

    return 0;

}