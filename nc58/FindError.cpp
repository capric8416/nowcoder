/**
 * struct TreeNode {
 *	int val;
 *	struct TreeNode *left;
 *	struct TreeNode *right;
 * };
 */

class Solution
{
public:
    int a = 0;
    int b = 0;
    int i = 0;
    TreeNode *pre = nullptr;

    void inorder(TreeNode *root)
    {
        if (!root)
        {
            return;
        }

        inorder(root->left);

        //确定a找到了，找错误节点b
        if (pre != nullptr && pre->val > root->val && i == 1)
        {
            b = root->val;
        }
        //找寻错误节点a
        if (pre != nullptr && pre->val > root->val && i == 0)
        {
            a = pre->val;
            i++;
        }

        pre = root;

        inorder(root->right);
    }

    /**
     * 
     * @param root TreeNode类 the root
     * @return int整型vector
     */
    vector<int> findError(TreeNode *root)
    {
        // write code here

        inorder(root);
        vector<int> res;
        res.push_back(b);
        res.push_back(a);
        return res;
    }
};