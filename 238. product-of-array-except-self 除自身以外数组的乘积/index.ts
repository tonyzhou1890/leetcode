/**
 * @param nums 
 * @desc 用 answer 存储左侧前缀之积，然后从右往左乘积
 */
function productExceptSelf(nums: number[]): number[] {
  const len = nums.length
  const answers = new Array(nums.length).fill(1)
  for (let i = 1; i < len; i++) {
    answers[i] = answers[i - 1] * nums[i - 1]
  }
  let right = 1
  for (let i = len - 2; i >= 0; i--) {
    right *= nums[i + 1]
    answers[i] = answers[i] * right
  }
  return answers
};

console.log(productExceptSelf([1,2,3,4])) // [24,12,8,6]
console.log(productExceptSelf([-1,1,0,-3,3])) // [0,0,9,0,0]