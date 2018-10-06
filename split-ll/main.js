let ll = {
    id: 1,
    next: { id: 2, next: { id: 3, next: { id: 4, next: null } } }
};

function splitLL(head) {
    let prev = null;
    let slow = head; // <= slow is now a reference to node 1
    let fast = head; // <= fast is now a reference to node 1
    while (fast !== null && fast.next !== null) {
        prev = slow; // prev = node 1
                     // prev = node 2
        slow = slow.next; // slow = node 2
                        // slow = node 3
        fast = fast.next.next;  // fast = node 3
                              // fast = node 4
    }

    prev.next = null; // this is now node 2, by setting this to null it edits head
    console.log("This is first half", head);
    console.log("This is second half", slow);
}
splitLL(ll);
