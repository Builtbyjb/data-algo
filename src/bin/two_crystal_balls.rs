fn two_crystal_balls(breaks: Vec<bool>) -> isize {
    let jmp_amount = breaks.len().isqrt() as isize;
    let len = breaks.len() as isize;

    let mut i = jmp_amount;

    // We use the first crystal ball to see where it breaks
    while i < len {
        if breaks[i as usize] {
            break;
        }

        i += jmp_amount;
    }

    i -= jmp_amount;

    let mut j = 0;

    while j < jmp_amount && i < len - 1 {
        j += 1;
        i += 1;

        if breaks[i as usize] {
            return i;
        }
    }

    return -1;
}

fn main() {
    let breaks = vec![false, false, false, false, false, false, true, true, true];
    let result = two_crystal_balls(breaks); // Ans: 6
    println!("Result: {}", result);
}

#[cfg(test)]
mod tests {
    use super::two_crystal_balls;

    #[test]
    fn test_break_found() {
        let breaks = vec![false, false, false, false, false, false, true, true, true];
        assert_eq!(two_crystal_balls(breaks), 6)
    }

    #[test]
    fn test_break_found_two() {
        let breaks = vec![
            false, false, false, false, false, false, false, false, false, false, true, true, true,
        ];
        assert_eq!(two_crystal_balls(breaks), 10)
    }

    #[test]
    fn test_break_not_found() {
        let breaks = vec![
            false, false, false, false, false, false, false, false, false,
        ];
        assert_eq!(two_crystal_balls(breaks), -1)
    }

    #[test]
    fn test_break_not_found_two() {
        let breaks = vec![
            false, false, false, false, false, false, false, false, false, false, false, false,
            false,
        ];
        assert_eq!(two_crystal_balls(breaks), -1)
    }
}
