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
	char_a(5);
	char_b(1),char_a(3);
	char_b(2),char_a(1);
	char_b(1),char_a(3);
	char_a(5);
	return 0;
}
