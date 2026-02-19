fn linear_search(haystack: [i16; 9], needle: i16) -> bool {
    for i in haystack {
        if needle == i {
            return true;
        }
    }

    return false;
}

fn main() {
    let haystack: [i16; 9] = [1, 2, 3, 4, 5, 6, 7, 8, 9];
    let needle: i16 = 5;
    let result = linear_search(haystack, needle);
    println!("Is {} in the haystack? {}", needle, result);
}

#[cfg(test)]
mod tests {
    use super::linear_search;

    #[test]
    fn test_linear_search_found() {
        let haystack: [i16; 9] = [1, 2, 3, 4, 5, 6, 7, 8, 9];
        let needle: i16 = 2;

        assert_eq!(linear_search(haystack, needle), true);
    }

    #[test]
    fn test_linear_search_not_found() {
        let haystack: [i16; 9] = [1, 11, 3, 4, 5, 6, 7, 8, 9];
        let needle: i16 = 2;

        assert_eq!(linear_search(haystack, needle), false);
    }
}
