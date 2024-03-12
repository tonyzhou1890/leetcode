/**
 * @param nums 
 * @desc 检查每一个位置可以到达的最远距离能不能连起来以及覆盖最后一个元素
 */
function canJump(nums: number[]): boolean {
  let destination = 0
  const len = nums.length
  for (let i = 0; i < len; i++) {
    if (destination >= i) {
      if (i + nums[i] > destination) {
        destination = i + nums[i]
      }
    } else {
      return false
    }
  }
  return true
};

console.log(canJump([2,3,1,1,4])) // true
console.log(canJump([3,2,1,0,4])) // false