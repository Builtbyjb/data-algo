
def linear_search(haystack: list[int], needle: int) -> bool:
    for i in haystack:
        if i == needle: return True
    return False

if __name__ == "__main__":
    haystack: list[int] = [1, 2, 3, 4, 5]
    result = linear_search(haystack, 5)
    print(result)
        
