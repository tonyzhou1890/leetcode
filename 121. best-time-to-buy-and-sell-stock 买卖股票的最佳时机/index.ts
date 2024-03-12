/**
 * 
 * @param prices 
 * @desc 总是在过去某天最低至买入，并计算今天卖出可以赚多少
 */
function maxProfit(prices: number[]): number {
  let min = prices[0]
  let profit = 0
  let len = prices.length
  for (let i = 1; i < len; i++) {
    if (prices[i] - min > profit) {
      profit = prices[i] - min
    }
    if (prices[i] < min) {
      min = prices[i]
    }
  }
  return profit
};

console.log(maxProfit([7,1,5,3,6,4])) // 5