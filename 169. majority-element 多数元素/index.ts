/**
 * @param nums 
 * @desc 摩尔投票法。类似消消乐。因为众数超过一半，所以即使其他所有数与众数抵消，依然会有剩下，如果其他数先相互抵消，最后剩下的还是众数。
 */
function majorityElement(nums: number[]): number {
  let count = 0
  let maj = 0
  let len = nums.length
  for (let i = 0; i < len; i++) {
    if (count) {
      if (maj === nums[i]) {
        count++
      } else {
        count--
      }
    } else {
      maj = nums[i]
      count = 1
    }
  }
  return maj
};

console.log(majorityElement([3,2,3]))