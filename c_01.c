#include <stdio.h>
int main(void)
{
	float a,b,c;
	printf("请输入实数a:");scanf("%f",&a);
	printf("请输入实数b:");scanf("%f",&b);
	printf("请输入实数c:");scanf("%f",&c);
	if (a==b && a==c)
		printf("1\n");
	else if(a==b || a==c)
		printf("2\n");
	else if (a+b>c && a+c>b && b+c>a )
		printf("3\n");
	else
		printf("0\n");
	return 0;
}
