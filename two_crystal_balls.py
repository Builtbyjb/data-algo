import math

def two_crystal_balls(breaks: list[bool]) -> int:
    jmp_amount: int = math.floor(math.sqrt((len(breaks))))

    i = jmp_amount
    while i < len(breaks):
        if breaks[i]: break
        i += jmp_amount

    i -= jmp_amount

    j = 0
    while j < jmp_amount and i < len(breaks):
        if breaks[i]: return i
        j += 1
        i += 1

    return -1


if __name__ == "__main__":
    breaks: list[bool] = [False, False, False, False, False, False, True, True]
    result = two_crystal_balls(breaks)
    print(result) # Ans: 6
