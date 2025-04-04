use std::collections::HashMap;



pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {

    let mut index_map: HashMap<i32, usize> = HashMap::new();

    for (i, item) in nums.iter().enumerate(){

        let diff = target - item;

        if let Some(&j) = index_map.get(&diff) {
            return vec![i as i32, j as i32];
        }

        index_map.insert(*item, i);
    }

    vec![]
}

fn main() {
    let nums = vec![3,2,4];

    let target = 6;


    println!("{:?}", two_sum(nums, target));

}