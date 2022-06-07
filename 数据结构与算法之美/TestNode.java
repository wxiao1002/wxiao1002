class TestNode{

  static class Node {
      int var;
      Node next;
      Node prev;

      public Node(int var, Node next) {
          this.var = var;
          this.next = next;
      }

      public Node(int var, Node next, Node prev) {
          this.var = var;
          this.next = next;
          this.prev = prev;
      }

  }

  // 链表检测环
    public boolean hasCycle (Node head){
      if (node == null || node.next == null){
          return false;
      }
      Node fast = head.next,slow = head;

      while (fast != slow){
          if (fast == null || fast.next == null){
              return false;
          }
          fast = fast.next.next;
          slow = slow.next;
      }
      return true;



    }
}