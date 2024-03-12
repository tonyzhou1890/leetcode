/**
 * 双指针
 * @param height 
 * @desc i 位置能接多少雨水是由左侧最高的柱子和右侧最高的柱子中的较小者决定的，所以，用 leftIdx 和 leftMax 表示左侧当前位置和 leftIdx 左侧最高的柱子，rightIdx 和 rightMax 表示右侧位置和 rightIdx 右侧最高的柱子。当 leftMax < rightMax，leftIdx 的接水量就确定了，当 rightMax <= leftMax，rightIdx 的接水了量就确定了。
 */
function trap(height: number[]): number {
  const len = height.length
  let leftIdx = 0
  let leftMax = height[leftIdx]
  let rightIdx = len - 1
  let rightMax = height[rightIdx]
  let count = 0
  while (leftIdx < rightIdx) {
    leftMax = Math.max(leftMax, height[leftIdx])
    rightMax = Math.max(rightMax, height[rightIdx])
    if (leftMax < rightMax) {
      count += leftMax - height[leftIdx]
      leftIdx++
    } else {
      count += rightMax - height[rightIdx]
      rightIdx--
    }
  }
  return count
};

console.log(trap([0,1,0,2,1,0,1,3,2,1,2,1])) // 6
console.log(trap([4,2,0,3,2,5])) // 9