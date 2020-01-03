class Solution {
public:
    int removeDuplicates(vector<int>& nums) {
        int i = !nums.empty();
        for (int n : nums)
            if (n > nums[i-1])
                nums[i++] = n;
        return i;
    }
};