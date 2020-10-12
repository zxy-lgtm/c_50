#include <stdio.h>
int main(void)
{
	int g1,g2;
	printf("成绩1:");scanf("%d",&g1);
	printf("成绩2:");scanf("%d",&g2);
	if(g1<0 || g1>100 || g2<0 || g2>100)
		printf("it is error.\n");
	else if (g1>=60 && g2>=60)
		printf("it is pass.\n");
	else 
		printf("it is not pass.\n");
	return 0;
}

