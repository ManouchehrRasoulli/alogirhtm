#include <stdio.h>
#include <stdlib.h>

struct ListNode {
    int val;
    struct ListNode *next;
};

void push(struct ListNode** head, int value) {
    struct ListNode* temp;
    temp = (struct ListNode*)malloc(sizeof(struct ListNode));
    temp->val = value;
    temp->next = NULL;

    if (*head == NULL) {
        *head = temp;
    } else {
        struct ListNode* ptr = *head;
		while(ptr->next!=NULL)
		{
			ptr = ptr->next;
		}
		ptr->next = temp;
    }
}

struct ListNode* addTwoNumbers(struct ListNode* l1, struct ListNode* l2) {
    int carry = 0;

    struct ListNode* l1ptr = l1;
    struct ListNode* l1headptr = NULL;
    
    while(l1ptr != NULL && l2 != NULL) {
        int v = l1ptr->val + l2->val + carry;
        int addValue = v % 10;
        carry = v / 10;

        l1ptr->val = addValue;

        l1ptr = l1ptr->next;
        l2 = l2->next;
    }

    while(l1ptr != NULL) {
        int v = l1ptr->val + carry;
        int addValue = v % 10;
        carry = v / 10;

        l1ptr->val = addValue;
        l1ptr = l1ptr->next;
    }
    
    while(l2 != NULL) {
        int v = l2->val + carry;
        int addValue = v % 10;
        carry = v / 10;

        push(&l1, addValue);
        l2 = l2->next;
    }

    if (carry != 0) {
        push(&l1, carry);
    }

    return l1;
}

int main() {
    struct ListNode l1pp = {1, 0};
    struct ListNode l1p = {9, NULL};
    struct ListNode l2p = {7, &l1pp};

    struct ListNode* resault = addTwoNumbers(&l1p, &l2p);

    while(resault != NULL) {
        printf("-- %i\n", resault->val);
        resault = resault->next;
    }

    return 0;
}

