#include <stdio.h>
#include <stdlib.h>

struct stack {
    int maxsize;
    int top;
    int *items;
};

struct stack* new_stack(int capacity) {
    struct stack *pt = (struct stack *) malloc(sizeof(struct stack));

    pt->maxsize = capacity;
    pt->top = -1;
    pt->items = (int *) malloc(sizeof(int) * capacity);

    return pt;
}

int size(struct stack * pt) {
    return pt->top + 1;
}

int is_empty(struct stack * pt) {
    return pt->top == -1;
}

int is_full(struct stack * pt) {
    return pt->top == pt->maxsize - 1;
}

void push(struct stack * pt, int value) {
    if (is_full(pt)) {
        printf("stack overflow !");
        exit(EXIT_FAILURE);
    }

    printf("inserting %d \n", value);

    ++pt->top;
    pt->items[pt->top] = value;
}

int peek(struct stack * pt) {
    if (is_empty(pt)) {
        printf("stack undeflow !\n");
        exit(EXIT_FAILURE);
    }

    return pt->items[pt->top];
}

void pop(struct stack * pt) {
    if (is_empty(pt)) {
        printf("stack underflow ! \n");
        exit(EXIT_FAILURE);
    }

    printf("popping %d \n", peek(pt));

    pt->items[pt->top--];
}

int main() {
    struct stack * s = new_stack(5);

    push(s, 1);
    push(s, 2);
    push(s, 3);
    
    printf("the top element of stack is %d \n", peek(s));
    printf("the size of stack is %d \n", s->top);

    pop(s);
    pop(s);
    pop(s);

     if (is_empty(s)) {
        printf("The stack is empty");
    }
    else {
        printf("The stack is not empty");
    }
 
    return 0;
}

