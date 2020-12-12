#include<stdio.h>
#include<stdlib.h>

typedef struct stduent {
    char num[17];
    int set;
    int true_set;
}a;

int main () {
    int n;

    scanf("%d", &n);

    a students[n];

    for (int i = 0; i < n; i++ ) {
        scanf("%s %d %d",students[i].num, &students[i].set, &students[i].true_set);
    }

    int e;

    scanf("%d", &e);

    int r[e];

    for (int i = 0; i < e; i++ ) {
        scanf("%d", &r[i]);
    }

   for (int count = 0; count < e; count ++ ) {
       for (int i = 0; i < n; i ++ ) {
            if(r[count] == students[i].set) {
            printf("%s %d\n", students[i].num, students[i].true_set);
        }
       }
   }

    return 0;

}
