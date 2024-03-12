/**
 * @param ratings 
 * @desc 分别从左右计算每个位置需要的糖果数量，然后取多的 
 */
function candy(ratings: number[]): number {
  const len = ratings.length
  let leftArr = new Array(len)
  let rightArr = new Array(len)
  leftArr[0] = 1
  // 从左侧计算
  for (let i = 1; i < len; i++) {
    if (ratings[i] > ratings[i - 1]) {
      leftArr[i] = leftArr[i - 1] + 1
    } else {
      leftArr[i] = 1
    }
  }
  let count = Math.max(1, leftArr[len - 1])
  let curr = 1
  // 从右侧计算
  for (let i = len - 2; i >= 0; i--) {
    if (ratings[i] > ratings[i + 1]) {
      curr += 1
    } else {
      curr = 1
    }
    count += Math.max(curr, leftArr[i])
  }
  return count
};

console.log(candy([1,0,2])) // 5
console.log(candy([1,2,2])) // 4