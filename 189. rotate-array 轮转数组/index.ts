/**
 Do not return anything, modify nums in-place instead.
 */
/**
 * 
 * @param nums 
 * @param k 
 * @desc 反转数组
 */
function rotate(nums: number[], k: number): void {
  const len = nums.length
  const offset = k % len
  // 第一步，反转全部
  for (let i = 0, half = len / 2; i < half; i++) {
    swap(nums, i, len - i - 1)
  }
  // 第二步，根据分割点反转左右两部分
  for (let i = 0, half = offset / 2; i < half; i++) {
    swap(nums, i, offset - i - 1)
  }
  for (let i = offset, half = (len - offset) / 2 + offset; i < half; i++) {
    swap(nums, i, len - (i - offset) - 1)
  }
};

function swap(nums: number[], a: number, b: number) {
  const temp = nums[a]
  nums[a] = nums[b]
  nums[b] = temp
}

const nums1 = [1,2,3,4,5,6,7]
rotate(nums1, 3)
console.log(nums1)