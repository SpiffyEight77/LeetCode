class Solution {
public:
    int searchInsert(vector<int>& nums, int target)
    {
        for (int i = 0; i < nums.size(); i++) {
            if (target <= nums[i]) {
                return i;
            }
        }
        return nums.size();
    }
};

class Solution {
public:
    int searchInsert(vector<int>& nums, int target)
    {
        int left, mid, right;
        right = nums.size();
        left = 0;
        mid = left + (right - left) / 2;
        while (left < right) {
            if (target > nums[mid]) {
                left = mid + 1;
            } else {
                right = mid;
            }
        }
        return left;
    }
};
