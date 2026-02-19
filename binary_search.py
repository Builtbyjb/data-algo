import math

def binary_search(haystack: list[int], number: int) -> bool:
    """ NOTE: Only works on a sorted array """
    low = 0
    high = len(haystack)

    while low < high:
        mid = math.floor( low + ((high-low) /2 ))
        value = haystack[mid]

        if value == number:
            return True
        elif value > number:
            high = mid
        elif value < number:
            low = mid + 1

    return False
        

if __name__ == "__main__":
    haystack: list[int] = [3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14 , 15, 16, 17]
    result = binary_search(haystack, 7)
    print(result)
