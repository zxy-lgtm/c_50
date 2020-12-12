#include<stdio.h>

int main () {
       char ch;
       int width, len;

       scanf("%d %c", &width, &ch);

       len = width * 0.5; 
       
       for(int i = 0; i < width; i ++ ) {
           printf("%c", ch);
       }
       printf("\n");

       for (int i = 0; i < len - 2; i++ ) {

           int count = 1;

           for (int j = 0; j < width ;j++ ) {
               if (count == 1 || count == width) {
                   printf("%c", ch);
               } else {
                   printf(" ");
               }
               count ++;
           }
           printf("\n");
       }

        for(int i = 0; i < width; i ++ ) {
           printf("%c", ch);
       }
       printf("\n");

       return 0;
}