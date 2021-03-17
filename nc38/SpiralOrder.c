#include <stdlib.h>
#include <string.h>
#include <stdio.h>


/**
 * 
 * @param matrix int整型二维数组 
 * @param matrixRowLen int matrix数组行数
 * @param matrixColLen int* matrix数组列数
 * @return int整型一维数组
 * @return int* returnSize 返回数组行数
 */
int* spiralOrder(int** matrix, int matrixRowLen, int* matrixColLen, int* returnSize ) {
    // write code here

    if(matrixRowLen == 0) {
        return NULL;
    }

    int rows = matrixRowLen;
    int columns = *matrixColLen;
    int count = rows * columns;

    int sizeInt = sizeof(int);

    int *result = (int*)malloc(sizeInt * count);

    for(int i = 0, r = 0, c = 0, row = rows-1, col = columns-1; r <= row && c <= col;) {
        // int length = col - c + 1;
        // memcpy(result + i * sizeInt, matrix[r] + c * sizeInt, sizeInt * length);
        // i += length;
        // r++;
        for(int cc = c; cc <= col; cc++, i++) {
            result[i] = matrix[r][cc];
        }
        r++;

        for(int rr = r; rr <= row; rr++, i++) {
            result[i] = matrix[rr][col];
        }
        col--;

        if(r <= row) {
            for(int cc = col; cc >= c; cc--, i++) {
                result[i] = matrix[row][cc];
            }
        }
        row--;

        if(c <= col) {
            for(int rr = row; rr >= r; rr--, i++) {
                result[i] = matrix[rr][c];
            }
        }
        c++;
    }

    *returnSize = count;

    return result;
}


int main(int argc, char **argv) {
    int rows = 3;
    int columns = 3;
    int matrix[3][3] = {{1, 2, 3}, {4, 5, 6}, {7, 8, 9}};

    int ret = 0;
    int *result = spiralOrder(matrix, 3, &columns, &ret);
    for(int i = 0; i < ret; i++) {
        printf("%d ", result[i]);
    }
    printf("\n");
    free(result);
}