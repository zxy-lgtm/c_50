#include <stdio.h>
int char_a(int i)
{
	for(int a=0;a<i;a++)
	putchar('*');
	printf("\n");
}
int char_b(int i)
{
	for (int a=0;a<i;a++)
	putchar(' ');
}
int main(void)
{
	int i=0;
	int j=7;
	for(i=0;i<3;i++)
	{
		j-=2;
		char_b(i);char_a(j);	
	}
	for(i=1;i>=0;i--)
	{
		j+=2;//printf("%d%d",j,i);
		char_b(i);char_a(j);
	}
	return 0;
}
