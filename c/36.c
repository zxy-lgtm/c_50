#include<stdio.h>

typedef struct person {
    char name[6];
    int year;
    int month;
    int day;
} a;

int compare_max (a cur, a max) {
    if (cur.year < max.year || (cur.year == max.year && cur.month < max.month) || (cur.year == max.year && cur.month == max.month && cur.day < max.day))
        return 1;
    else
    {
        return 0;
    }
    
}

int main () {

    int n;
    int count = 0;
    scanf("%d", &n);

    a min, max;
    a cur;
    max.year = 2014;max.month = 9; max.day = 6;
    min.year = 1814;min.month = 9; min.day = 6;

    for (int i = 0; i < n; i ++ ) {
        scanf("%s %d/%d/%d", &cur.name, &cur.year, &cur.month, &cur.day);
        if ((cur.year < 2014 || (cur.year == 2014 && cur.month < 9) || (cur.year == 2014 && cur.month == 9 && cur.day < 7))\
         && (cur.year > 1814 || (cur.year == 1814 && cur.month > 9) ||(cur.year == 1814 && cur.month ==9 && cur.day > 5))) {
            count ++;
        
            if (compare_max(cur,max) == 1) {
                max = cur;
            }

            if (compare_max(cur,min) == 0) {
                min = cur;
            }
        }
    }


    printf("%d %s %s\n", count, max.name, min.name);

    return 0;

}