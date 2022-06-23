# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def mergeTwoLists(self, list1: Optional[ListNode], list2: Optional[ListNode]) -> Optional[ListNode]:
        head = cur = None
        if not list1:
            return list2
        
        if not list2:
            return list1
        
        while list1 or list2:
            node = None
            if not list1:
                node = ListNode(list2.val)
                list2 = list2.next
            elif not list2:
                node = ListNode(list1.val)
                list1 = list1.next
            elif list1.val < list2.val:
                node = ListNode(list1.val)
                list1 = list1.next
            else: 
                node = ListNode(list2.val)
                list2 = list2.next
                
            if not head:
                head = cur = node
            else:
                cur.next = node
                cur = cur.next
        
        return head
        
        