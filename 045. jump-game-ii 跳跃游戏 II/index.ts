/**
 * @param nums
 * @desc 从当前元素可以到达的范围内，记录范围内再次跳跃最远距离，依次跳跃
 */
function jump(nums: number[]): number {
  const len = nums.length
  let destination = 0
  let nextDestination = 0
  let count = 0
  if (len === 1) return count
  for (let i = 0; i < len; i++) {
    if (nums[i] + i > nextDestination) {
      nextDestination = nums[i] + i
      if (nextDestination >= len - 1) {
        // 提前计数
        count++
        break
      }
    }
    if (destination === i) {
      destination = nextDestination
      count++
    }
  }
  return count
};

console.log(jump([2,3,1,1,4])) // 2
console.log(jump([2,3,0,1,4])) // 2
console.log(jump([0])) // 0
console.log(jump([1,2])) // 1
console.log(jump([2,3,1,1,4])) // 2