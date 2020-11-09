#include<stdio.h>
#include<stdlib.h>

typedef struct {
    int stuid;
    int score;
}student;

void getinto(student a[], int n) {
    
    for(int i = 0; i < n; i ++ ) {
        scanf("%d", &a[i].stuid);
        scanf("%d", &a[i].score);
    }
}

void scoresort(student a[], int n) {

    //冒泡排序法
    int i, j ;
    student tool;
    for (i = 0; i < n - 1; i ++ ) {
        for (j = 0; j < n - 1 - i; j ++ ) {
            if (a[j].score > a[j + 1].score) {
                tool = a[j];
                a[j + 1] = a[j];
                a[j] = tool;
            }
        }
    }
    return ;
}

void scoredasc(student a[], int n) {

    for (int i = 0; i < n; i ++ ) {
        printf("%d %d\n", a[i].stuid, a[i].score );
    }
}

void scoreasc(student a[], int n) {

    for (int i = n - 1; i >= 0; i -- ) {
        printf("%d %d\n", a[i].stuid, a[i].score);
    }
}

int main () {

    int n;
    printf("请输入一共有几个学生吧");
    scanf("%d", &n);

    student a[n];

    getinto(a, n);
    scoresort(a, n);
    printf("升序输出\n");
    scoredasc(a, n);
    printf("降序输出\n");
    scoreasc(a, n);

    return 0;
}