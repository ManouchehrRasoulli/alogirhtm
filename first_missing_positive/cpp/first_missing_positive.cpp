#include <iostream>
#include <vector>

using namespace std;

int firstMissingPositive(vector<int>& nums) {
    int n = nums.size();

    for (int i = 0; i < n; i++) {
        while (nums[i] > 0 && nums[i] < n && nums[i] != nums[nums[i] - 1]) {
            int x = nums[nums[i] - 1];
            nums[nums[i] - 1] = nums[i];
            nums[i] = x;
        }
    }

    for (int i = 0; i < n; i++) {
        if (nums[nums[i]- 1] != i+1) {
            return i+1;
        }
    }


    return n+1;
}

int main() {
    vector nums {2, 1, 0};
    cout << "nums contains #" << nums.size() << " items\n";

    int missing = firstMissingPositive(nums);

    cout << "missing one is " << missing << "\n";

    return 0;
}
