
// https://leetcode.com/problems/reverse-integer/

pub fn reverse(mut x: i32) -> i32 {

    let mut revered= 0;

    while x != 0 {
        revered = revered * 10 + x%10;
        x /= 10;
    }
    
    if revered > std::i32::MAX || revered < std::i32::MIN {
        return 0
    }

    revered
}

fn main() {

    let x = -123;

    println!("{:?}", reverse(x));
    println!("{}", std::i32::MAX > 1056389759)
}
