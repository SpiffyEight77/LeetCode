class Solution {
public:
    int pivotIndex(vector<int>& nums) {
        int left, right;
        left = right = 0;
        
        for (int i = 0; i < nums.size(); i++) {
            right+=nums[i];
        }
        for (int i = 0; i < nums.size(); i++) {
            if (left * 2 == right - nums[i]) {
                return i;
            }
            left += nums[i];
        }
        return -1;
    }
};
