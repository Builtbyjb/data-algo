/// Time complexity O(N^2)
fn bubble_sort(arr: &mut Vec<i16>) {
    let len = arr.len();

    for i in 0..len {
        for j in 0..len - 1 - i {
            if arr[j] > arr[j + 1] {
                let tmp = arr[j];
                arr[j] = arr[j + 1];
                arr[j + 1] = tmp
            }
        }
    }
    return;
}

fn main() {
    let mut arr: Vec<i16> = vec![3, 6, 1, 2, 7, 9, 4, 5, 6, 8];
    bubble_sort(&mut arr);
    println!("Arr: {:?}", arr);
}

#[cfg(test)]
mod tests {
    use super::bubble_sort;

    #[test]
    fn sorted() {
        let mut arr = vec![3, 6, 1, 2, 7, 9, 4, 5, 6, 8];
        let sorted = vec![1, 2, 3, 4, 5, 6, 6, 7, 8, 9];
        bubble_sort(&mut arr);
        assert_eq!(arr, sorted)
    }

    #[test]
    fn sorted_two() {
        let mut arr = vec![3, 6, 1, 2, 7, 9, 4, 5, 6, 8, 11, 19, 2, 8, 14];
        let sorted = vec![1, 2, 2, 3, 4, 5, 6, 6, 7, 8, 8, 9, 11, 14, 19];
        bubble_sort(&mut arr);
        assert_eq!(arr, sorted)
    }
}
