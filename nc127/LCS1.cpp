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
    int size1 = str1.size();
    int size2 = str2.size();
    if (size1 == 0 || size2 == 0)
    {
        return 0;
    }

    // the start position of substring in str1
    int offset = -1;

    // the longest length of common substring
    int longest = 0;

    int indices[] = {0, 0};
    int sizes[] = {size1, size2};

    // shift strings to find the longest common substring
    for (int index = 0; index < 2; ++index)
    {
        indices[0] = 0;
        indices[1] = 0;

        // i is reference to the value in array
        int &i = indices[index];
        int size = sizes[index];

        // this is tricky to skip comparing strings both start with 0 for second loop
        i = index;
        for (; i < size; ++i)
        {
            int m = indices[0];
            int n = indices[1];
            int length = 0;

            // with following check to reduce some more comparisons
            if (size1 - m <= longest || size2 - n <= longest)
            {
                break;
            }

            while (m < size1 && n < size2)
            {
                if (str1[m] != str2[n])
                {
                    length = 0;
                }
                else
                {
                    ++length;
                    if (longest < length)
                    {
                        longest = length;
                        offset = m - longest + 1;
                    }
                }

                ++m;
                ++n;
            }
        }
    }

    if (offset == -1)
    {
        return "";
    }

    return str1.substr(offset, longest);
}

int main(int argc, char **argv)
{
    cout << LCS("ab", "a") << endl;
    cout << LCS("a", "b") << endl;
    cout << LCS("afgha", "abcde") << endl;
    cout << LCS("22222", "22222") << endl;
    cout << LCS("1AB2345CD", "12345EF") << endl;
    cout << LCS("2LQ74WK8Ld0x7d8FP8l61pD7Wsz1E9xOMp920hM948eGjL9Kb5KJt80", "U08U29zzuodz16CBZ8xfpmmn5SKD80smJbK83F2T37JRqYfE76vh6hrE451uFQ100ye9hog1Y52LDk0L52SuD948eGjLz0htzd5YF9J1Y6oI7562z4T2") << endl;

    getchar();
}
