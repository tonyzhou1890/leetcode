/**
 * @param citations 
 * @desc 套百度的公式 要确定一个人的h指数非常容易，到SCI网站，查出某个人发表的所有SCI论文，让其按被引次数从高到低排列，往下核对，直到某篇论文的序号大于该论文被引次数
 */
function hIndex(citations: number[]): number {
  citations.sort((a, b) => b - a)
  let len = citations.length
  let h = 0
  for (let i = 0; i < len; i++) {
    if ((i + 1) > citations[i]) {
      break
    }
    h = i + 1
  }
  return h
};

console.log(hIndex([3,0,6,1,5])) // 3
console.log(hIndex([1,3,1])) // 1
console.log(hIndex([1])) // 1