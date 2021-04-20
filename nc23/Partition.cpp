
struct ListNode {
    int val;
    struct ListNode *next;
};
 

class Solution {
public:
    /**
     * 
     * @param head ListNode类 
     * @param x int整型 
     * @return ListNode类
     */
    ListNode* partition(ListNode* head, int x) {
        // write code here

        if(head == nullptr || head->next == nullptr) {
            return head;
        }

        ListNode low = ListNode();
        ListNode high = ListNode();
        ListNode* p = &low;
        ListNode* q = &high;
        for(ListNode* curr = head; curr != nullptr; curr = curr->next) {
            if(curr->val < x) {
                p->next = curr;
                p = curr;
            } else {
                q->next = curr;
                q = curr;
            }
        }

        p->next = high.next;
        q->next = nullptr;

        return low.next;
    }
};


int main(int argc, char** argv) {
    return 0;
}