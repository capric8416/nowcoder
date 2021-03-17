/**
 * 反转字符串
 * @param str string字符串 
 * @return string字符串
 */
char* solve(char* str ) {
    // write code here

    char temp;
    for(int i = 0, j = strlen(str)-1; i < j; i++, j--) {
        temp = str[i];
        str[i] = str[j];
        str[j] = temp;
    }

    return str;
}