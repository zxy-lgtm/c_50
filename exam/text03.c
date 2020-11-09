#include<stdio.h>
#include<stdlib.h>

typedef struct node {
    int Val;
    struct node *Prev;
    struct node *Next;
}listNode;

listNode* receive (listNode *ptr, int n) {
    for (int i = 0; i < n; i ++ ) {
        listNode* node = (listNode*)malloc(sizeof(listNode));
        scanf("%d", &node->Val);
        node -> Next = ptr;
        node -> Prev = NULL;
        if (ptr != NULL)
            ptr -> Prev = node;
        ptr = node;
    }

    return ptr;
}

void f_or_t (listNode *ptr, int n) {

    int tool = 0;
    int count = 0;
    listNode *head = ptr;
    do {
        ptr = ptr ->Next;
        count ++ ;
    } while( count != n - 1);

    count = 0;
    do {
        listNode*before = ptr;
        listNode*behind = head;
        //只要有一组相对的数不同,tool就会改变
        if((before -> Val) != (behind -> Val)) {
            tool ++;
        }
        before = before->Prev;
        behind = behind->Next;
        count ++;
    }while(count != n - 1);

    if (tool == 0) {
        printf("true\n");
    } else {
        printf("false\n");
    }
}

int main () {

    int n;
    scanf("%d\n", &n);

    listNode*ptr = NULL;
    ptr = receive(ptr, n);
    f_or_t(ptr, n);

    return 0;
}