# 287. Find the Duplicate Number

Given an array of integers nums containing n + 1 integers where each integer is in the range [1, n] inclusive.

There is only one repeated number in nums, return this repeated number.

You must solve the problem without modifying the array nums and uses only constant extra space.

### Example 1:

```
Input: nums = [1,3,4,2,2]
Output: 2
```

### Example 2:

```
Input: nums = [3,1,3,4,2]
Output: 3
```

### Constraints:

* 1 <= n <= 105
* nums.length == n + 1
* 1 <= nums[i] <= n
* All the integers in nums appear only once except for precisely one integer which appears two or more times.

### Follow up:

* How can we prove that at least one duplicate number must exist in nums?
* Can you solve the problem in linear runtime complexity?

### Translate:

> 287. 寻找重复数

给定一个包含n + 1 个整数的数组nums ，其数字都在[1, n]范围内（包括 1 和 n），可知至少存在一个重复的整数。

假设 nums 只有 一个重复的整数 ，返回这个重复的数 。

你设计的解决方案必须 不修改 数组 nums 且只用常量级 O(1) 的额外空间。

示例 1：

输入：nums = [1,3,4,2,2]
输出：2 示例 2：

输入：nums = [3,1,3,4,2]
输出：3

提示：

1 <= n <= 105 nums.length == n + 1 1 <= nums[i] <= n nums 中 只有一个整数 出现 两次或多次 ，其余整数均只出现 一次

进阶：

如何证明 nums 中至少存在一个重复的数字? 你可以设计一个线性级时间复杂度 O(n) 的解决方案吗？