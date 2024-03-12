/**
 * @param numbers 
 * @param target 
 * @desc 双指针，两端收拢。如果和大于目标值，右侧往左，如果和小于目标值，左侧往右
 */
function twoSum(numbers: number[], target: number): number[] {
  let i = 0
  let j = numbers.length - 1
  let sum = 0
  while (i < j) {
    sum = numbers[i] + numbers[j]
    if (sum > target) {
      j--
    } else if (sum < target) {
      i++
    } else {
      return [i + 1, j + 1]
    }
  }
  return []
};

console.log(twoSum([2,7,11,15], 9)) // [1, 2]
console.log(twoSum([2,3,4], 6)) // [1, 3]
console.log(twoSum([-1,0], -1)) // [1, 2]