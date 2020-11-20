#include<stdio.h>
int month (int x,int w2){
    int i;
    for (i = 0; i < w2 * 4; i ++) {
        putchar(' ');
    }
    for(i = 1; i <= x; i ++ ) {
        printf("%-3d",i);
        if((i + w2) % 7 == 0)
            putchar('\n');
    }
    if ((x + w2) % 7 != 0)
        putchar('\n');
    return (x - 28 + w2) % 7;
}

void main () {
    int n, c, y, w, days = 28;
    scanf("%d",&n);
    c = n / 100;
    y = n % 100 - 1;
    w = (c/4-2*c+y+y/4+14*13/5)%7;
    if(n%400==0||n%4==0&&n%100!=0){
        days = 29;
    }
    puts("==========================Jan=========================\nSun Mon Tue WednThu Fri Sat");
    w = month(31,w);

    puts("==========================Feb=========================\nSun Mon Tue WednThu Fri Sat");
    w = month(days,w);

    puts("==========================Mar=========================\nSun Mon Tue WednThu Fri Sat");
    w = month(31,w);

    puts("==========================Apr=========================\nSun Mon Tue WednThu Fri Sat");
    w = month(30,w);

    puts("==========================May=========================\nSun Mon Tue WednThu Fri Sat");
    w = month(31,w);

    puts("==========================Jun=========================\nSun Mon Tue WednThu Fri Sat");
    w = month(30,w);

    puts("==========================Jul=========================\nSun Mon Tue WednThu Fri Sat");
    w = month(31,w);

    puts("==========================Aug=========================\nSun Mon Tue WednThu Fri Sat");
    w = month(31,w);

    puts("==========================Sep=========================\nSun Mon Tue WednThu Fri Sat");
    w = month(30,w);

    puts("==========================Oct=========================\nSun Mon Tue WednThu Fri Sat");
    w = month(31,w);

    puts("==========================Nov=========================\nSun Mon Tue WednThu Fri Sat");
    w = month(30,w);

    puts("==========================Dec=========================\nSun Mon Tue WednThu Fri Sat");
    w = month(31,w);

}
