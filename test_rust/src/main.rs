use std::ops::Add;

#[derive(Debug, Clone)]
struct Tensor<T> {
    data: Vec<T>,
}

impl<T> Add for Tensor<T>
where
    T: Add<Output = T> + Copy,
{
    type Output = Self;

    fn add(self, other: Self) -> Self {
        let result: Vec<T> = 
        self
            .data
            .iter()
            .zip(other.data.iter())
            .map(|(&a, &b)| a + b)
            .collect();

        Tensor::new(result)
    }
}

// Define a module torch
mod torch {
    use super::Tensor;
    use std::ops::Mul;

    // Dot product function
    pub fn dot<T>(tensor1: &Tensor<T>, tensor2: &Tensor<T>) -> T
    where
        T: Mul<Output = T> + super::Add<Output = T> + Copy + Default,
    {
        tensor1
            .data
            .iter()
            .zip(tensor2.data.iter())
            .map(|(&a, &b)| a * b)
            .fold(T::default(), |sum, val| sum + val)
    }
}

// Implement Tensor methods
impl<T> Tensor<T> {
    fn new(data: Vec<T>) -> Tensor<T> {
        Tensor { data }
    }
}

fn main() {
    let tensor1 = Tensor::new(vec![1.0, 0.2]);
    let tensor2 = Tensor::new(vec![0.0, 1.2]);

    let sum_tensor = tensor1.clone() + tensor2.clone(); 
    let dot_product = torch::dot(&tensor1, &tensor2); 

    println!("{:?}", sum_tensor); // Output: Tensor { data: [11, 22
    println!("Dot Product: {}", dot_product); // Output: Dot Product: 300
}
