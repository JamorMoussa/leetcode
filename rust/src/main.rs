use std::{collections::HashMap, vec};


#[derive(Debug)]
struct Stack { 
    _stack: Vec<char>
}

impl Stack {

    fn new() -> Stack{
        Stack{_stack: Vec::<char>::new()}
    }
    
    fn push(&mut self, item: char) {
        self._stack.push(item);
    }

    fn from_str(items: &str) -> Stack {
        Stack {_stack: items.chars().collect::<Vec<char>>()}
    }

    fn pop(&mut self) -> Option<char>{
        self._stack.pop()
    }

    fn is_empty(&self) -> bool {
        self._stack.is_empty()
    }

    fn last(&self) -> Option<&char> {
        self._stack.last()
    }

    fn len(&self) -> usize {
        self._stack.len()
    }

}


fn is_valid(s: String) -> bool {
    let bracket_map: HashMap<char, char> =
        [(')', '('), (']', '['), ('}', '{')].iter().cloned().collect();

    let opening_brackets: Vec<char> = bracket_map.values().cloned().collect();

    let mut stack = Stack::new();

    for c in s.chars() {
        if opening_brackets.contains(&c) {
            stack.push(c);
        } else if let Some(&expected) = bracket_map.get(&c) {
            if stack.pop() != Some(expected) {
                return false;
            }
        }
    }

    stack.is_empty()
}


fn main() {
    println!("{}", is_valid("({})".to_string()));
}