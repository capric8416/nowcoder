#include <iostream>
#include <string>

using namespace std;

/**
 * longest common substring
 * @param str1 string字符串 the string
 * @param str2 string字符串 the string
 * @return string字符串
 */
string LCS(string str1, string str2)
{
    // write code here

    size_t m = str1.size();
    size_t n = str2.size();

    if (m == 0 || n == 0)
    {
        return "";
    }

    // new
    int **dp = new int *[m];
    for (int i = 0; i < m; i++)
    {
        dp[i] = new int[n];
    }

    int str2Index = 0;
    int maxCommonLength = 0;

    // init 1st row and column
    for (int i = 0; i < m; i++)
    {
        if (str1[i] == str2[0])
        {
            dp[i][0] = 1;
            maxCommonLength = 1;
        }
        else
        {
            dp[i][0] = 0;
        }
    }
    for (int j = 0; j < n; j++)
    {
        if (str2[j] == str1[0])
        {
            dp[0][j] = 1;
            maxCommonLength = 1;
        }
        else
        {
            dp[0][j] = 0;
        }
    }

    for (int i = 1; i < m; i++)
    {
        for (int j = 1; j < n; j++)
        {
            if (str1[i] == str2[j])
            {
                dp[i][j] = dp[i - 1][j - 1] + 1;
                if (maxCommonLength < dp[i][j])
                {
                    str2Index = j + 1;
                    maxCommonLength = dp[i][j];
                }
            }
        }
    }

    // delete
    for (int i = 0; i < m; i++)
    {
        delete[] dp[i];
    }
    delete[] dp;

    if (maxCommonLength == 0)
    {
        return "";
    }
    if (maxCommonLength == 1 && str2Index == 0)
    {
        return str2.substr(0, 1);
    }
    return str2.substr(str2Index - maxCommonLength, maxCommonLength);
}

int main(int argc, char **argv)
{
    // cout << LCS("ab", "a") << endl;
    // cout << LCS("a", "b") << endl;
    // cout << LCS("afgha", "abcde") << endl;
    cout << LCS("22222", "22222") << endl;
    // cout << LCS("1AB2345CD", "12345EF") << endl;
    // cout << LCS("2LQ74WK8Ld0x7d8FP8l61pD7Wsz1E9xOMp920hM948eGjL9Kb5KJt80", "U08U29zzuodz16CBZ8xfpmmn5SKD80smJbK83F2T37JRqYfE76vh6hrE451uFQ100ye9hog1Y52LDk0L52SuD948eGjLz0htzd5YF9J1Y6oI7562z4T2") << endl;
}
