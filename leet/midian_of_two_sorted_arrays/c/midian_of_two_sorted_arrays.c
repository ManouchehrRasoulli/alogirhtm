#include <stdio.h>
#include <stdlib.h>

double findMedianSortedArrays(int* nums1, int nums1Size, int* nums2, int nums2Size) {

    if (nums1Size > nums2Size) {
        return findMedianSortedArrays(nums2, nums2Size, nums1, nums1Size);
    }

    int total = nums1Size+nums2Size;

    if (nums1Size == 0) {
        if (total % 2 == 1) {
            return nums2[total/2];
        } else {
            return (nums2[total/2] + nums2[total/2 - 1]) / (double)2;
        }
    }

    int nums[total];

    int num1Counter = 0;
    int num2Counter = 0;
    int index = 0;

    for (int i = 0; i < total; i++) {
        if (nums1[num1Counter] < nums2[num2Counter]){
            nums[i] = nums1[num1Counter++];
        } else if (nums2[num2Counter] < nums1[num1Counter]) {
            nums[i] = nums2[num2Counter++];
        } else {
            nums[i] = nums1[num1Counter++];
        }

        if (num1Counter == nums1Size || num2Counter == nums2Size) {
            index = i+1;
            break;
        }
    }

    if (num1Counter == nums1Size) {
        for (int i = index; i < total; i++) {
            nums[i] = nums2[num2Counter++];
        }
    } else {
        for (int i = index; i < total; i++) {
            nums[i] = nums1[num1Counter++];
        }
    }

    if (total % 2 == 1) {
        return nums[total/2];
    } else {
        return (nums[total/2] + nums[total/2 - 1]) / (double)2;
    }
}

int main() {
    printf("ok\n");

    int nums1[] = {3};
    int nums2[] = {-2, -1};

    printf("%f\n", findMedianSortedArrays(nums1, 1, nums2, 2));

    return 0;
}

