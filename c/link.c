#include<stdio.h>
#include<stdlib.h>//为了用malloc

typedef struct _node{
    int value;
    struct _node *next;
    
}Price;//这是一个结构放入了一个value还有一个尾部指针可以实现连接.并且将它重新命名为Price.
//输入
Price*  putchar_linklist(void){
    char ch;
    int number;
    Price*head=NULL;//头指针
    Price*p;
    Price*touch;
    do{
        Price *p=(Price*)malloc(sizeof(Price));
        scanf("%d",&number);
        p->value = number;
        if(head==NULL){
            head=p;
        }
        else{
            touch->next=p;//接到后面
        }
        touch=p;//接到前面
    }while(ch=getchar()!='\n');
    return head;
}
//输出+遍历.
void outputchar_linklist(Price*head){
   while(head!=NULL){
    printf("%d ",head->value);
    head=head->next;
    } 
}
//删除特定数字 从键盘输入
void delete(Price*lemon){
    int i;
    printf("输入你想删除的数字吧:");
    scanf("%d",&i);
    while(lemon!=NULL){
        if(lemon->value!=i){
        printf("%d ",lemon->value);
        lemon=lemon->next;
        }
        else
        lemon=lemon->next;
    }
}
//添加特定数字 从键盘输入:
Price* add(Price*lemon,int n1,int n2){
    Price*head=lemon;
    for(int i=1;i<n2;i++){
    head=head->next;
    }
    Price*w =(Price*)malloc(sizeof(Price));
    w->value=n1;
    w->next=head->next;
    head->next=w;
    printf("现在它变成这样啦:");
    while(head!=NULL){
        printf("%d ",head->value);
        head=head->next;
    }
}
//主函数
int main(void){
    Price*lemon;
    printf("输入一串数字吧\n");
    lemon=putchar_linklist();
    printf("你输入了\n");
    outputchar_linklist(lemon);
    printf("\n");
    delete(lemon);
    int i;
    int j;
    printf("输入你想加入的数字吧:");
    scanf("%d",&i);
    printf("提前说好要加到谁的后面呢?\n");
    scanf("%d",&j);
    printf("看,我们改成这样啦");
    add(lemon,i,j);
    return 0;
}