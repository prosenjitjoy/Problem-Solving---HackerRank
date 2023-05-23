 Node* insert(Node *head,int data)
      {
          //Complete this method
          Node *node = new Node(data);
          
          if (head == NULL) {
              head = node;
              return head;
          }
          
          Node *tmp = head;
          while (tmp->next) {
              tmp = tmp->next;
          }
          tmp->next = node;
          
          return head;
      }