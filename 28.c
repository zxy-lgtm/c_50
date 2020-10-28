#include <stdio.h>
#include <string.h>

int main()
{

    int n;
    int count = 0;
    char *ch[1000];

    memset(ch, '\0', 1000);
    gets(ch);

    for (int j = 0; ch[j] != '\0'; j++)
    {
        if ((int)ch[j] >= '0' && (int)ch[j] <= '9')
        {
            count++;
        }
    }
    printf("%d", count);

    return 0;
}