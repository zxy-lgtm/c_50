#include <stdio.h>
int char_a(int i)
{
	for(int a=0;a<i;a++)
	putchar('*');
}
int char_b(int i)
{
	for (int a=0;a<i;a++)
	putchar(' ');
}
int main(void)
{
	char_a(5);
	printf("\n");
	char_b(1),char_a(3);
	printf("\n");
	char_b(2),char_a(1);
	printf("\n");
	char_b(1),char_a(3);
	printf("\n");
	char_a(5);
	printf("\n");
	return 0;
}
