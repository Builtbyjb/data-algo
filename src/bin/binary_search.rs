/// NOTE: Primary conditions for binary search
/// - Value equals target
/// - Value is larger than target
/// - Value is smaller than target
fn binary_search(haystack: [i16; 9], needle: i16) -> bool {
    let mut low = 0;
    let mut high = haystack.len();

    while low < high {
        let mid = low + ((high - low) / 2);
        let value = haystack[mid as usize];

        if value == needle {
            return true;
        } else if value > needle {
            high = mid
        } else {
            low = mid + 1
        }
    }

    return false;
}

fn main() {
    let haystack = [1, 2, 3, 4, 5, 6, 7, 8, 9];
    let needle = 5;
    let found = binary_search(haystack, needle);
    println!("{}", found);
}

#[cfg(test)]
mod tests {
    use super::binary_search;

    #[test]
    fn test_found() {
        let haystack = [1, 2, 3, 4, 5, 6, 7, 8, 9];
        let needle = 5;

        assert_eq!(binary_search(haystack, needle), true)
    }

    #[test]
    fn test_not_found() {
        let haystack = [1, 2, 3, 4, 6, 7, 8, 9, 11];
        let needle = 5;

        assert_eq!(binary_search(haystack, needle), false)
    }
}
