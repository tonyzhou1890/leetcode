/**
 * 
 * @param height 
 * @desc 双指针，一左一右，水量由短板决定，所以每次移动短板，因为移动长板时短板不变，但底边变短，所以水量肯定变少。
 */
function maxArea(height: number[]): number {
  let i = 0
  let j = height.length - 1
  let max = 0
  while (i < j) {
    max = Math.max(max, (j - i) * Math.min(height[i], height[j]))
    if (height[i] < height[j]) {
      i++
    } else {
      j--
    }
  }
  return max
};

console.log(maxArea([1,8,6,2,5,4,8,3,7]) === 49)
console.log(maxArea([1,1]) === 1)