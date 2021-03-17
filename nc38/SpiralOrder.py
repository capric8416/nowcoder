class Solution:
    def spiralOrder(self, matrix):
        rows = len(matrix)
        if rows == 0:
            return []

        columns = len(matrix[0])

        count = rows * columns
        result = [None] * count

        i, r, c, row, col = 0, 0, 0, rows-1, columns-1
        while r <= row and c <= col:
            length = col - c + 1
            result[i: i+length] = matrix[r][c: c+length]
            i += length
            r += 1

            rr = r
            while rr <= row:
                result[i] = matrix[rr][col]
                rr, i = rr+1, i+1
            col -= 1

            if r <= row:
                length = col - c + 1
                result[i: i + length] = matrix[row][col-length+1: col+1][::-1]
                i += length
            row -= 1

            if c <= col:
                rr = row
                while rr >= r:
                    result[i] = matrix[rr][c]
                    rr, i = rr-1, i+1
            c += 1

        return result


if __name__ == '__main__':
    print(Solution().spiralOrder([[]]))

    print(Solution().spiralOrder([[1, 2, 3, 4]]))

    print(Solution().spiralOrder([[1], [2], [3], [4]]))

    print(
        Solution().spiralOrder([
            [1, 2, 3, 4, 5, 6, 7, 8, 9, 10],
            [11, 12, 13, 14, 15, 16, 17, 18, 19, 20],
            [21, 22, 23, 24, 25, 26, 27, 28, 29, 30],
            [31, 32, 33, 34, 35, 36, 37, 38, 39, 40],
            [41, 42, 43, 44, 45, 46, 47, 48, 49, 50],
            [51, 52, 53, 54, 55, 56, 57, 58, 59, 60],
            [61, 62, 63, 64, 65, 66, 67, 68, 69, 70],
            [71, 72, 73, 74, 75, 76, 77, 78, 79, 80],
            [81, 82, 83, 84, 85, 86, 87, 88, 89, 90],
            [91, 92, 93, 94, 95, 96, 97, 98, 99, 100],
            [101, 102, 103, 104, 105, 106, 107, 108, 109, 110]
        ])
    )
