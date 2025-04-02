
#[derive(Debug, Clone)]
pub struct ListNode {
    val: i32,
    next: Option<Box<ListNode>> 
}

impl ListNode {
    
    fn new(val: i32) -> ListNode {
        ListNode { val: val, next: None}
    }

    fn add_node(&mut self, node: ListNode){
        let mut current = self;

        while let Some(ref mut next_node) = current.next {
            current = next_node.as_mut();
        }

        current.next = Some(Box::new(node));
    }

    fn print(&self) { 

        let mut current = Some(self);

        while let Some(node) = current {
            println!("{}", node.val);
            current = node.next.as_deref();
        }
    }

}

fn main() {

    let mut head = ListNode::new(1);
    
    head.add_node(
        ListNode::new(2)
    );
    
    head.add_node(
        ListNode::new(2)
    );

    head.print();

}
