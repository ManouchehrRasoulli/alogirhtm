#include <stdio.h>

int firstMissingPositive(int* nums, int numsSize) {
    
    for (int i = 0; i < numsSize; i++) {
		while (nums[i] > 0 && nums[i] < numsSize && nums[i] != nums[nums[i]-1]) {
            int x = nums[nums[i] - 1];
            nums[nums[i] - 1] = nums[i];
            nums[i] = x;
		}
	}

	for (int i = 0; i < numsSize; i++) {
		if (nums[i] != i+1) {
			return i + 1;
		}
	}


    return numsSize+1;
}

int main() {
    puts("ok !");

    int nums[] = {-1,4,2,1,9,10};
    printf("%i\n", firstMissingPositive(nums, 6));

    return 0;
}

