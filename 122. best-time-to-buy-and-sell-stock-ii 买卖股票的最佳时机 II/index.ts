/**
 * 
 * @param prices 
 * @desc 买卖分很多区间，区间的划分以这一天对比前一天价格算，只要降了，前一天就是上一个区间截止，这一天就是下一个区间开始
 */
function maxProfit(prices: number[]): number {
  let min = prices[0], profit = 0, len = prices.length
  for (let i = 0; i < len; i++) {
    if (prices[i] < prices[i - 1]) {
      profit += prices[i - 1] - min
      min = prices[i]
    }
  }
  if (prices[len - 1] > min) {
    profit += prices[len - 1] - min
  }
  return profit
};

console.log(maxProfit([7,1,5,3,6,4])) // 7
console.log(maxProfit([1,2,3,4,5])) // 4
console.log(maxProfit([7,6,4,3,1])) // 0