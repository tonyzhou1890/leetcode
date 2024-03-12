function canCompleteCircuit(gas: number[], cost: number[]): number {
  const len = gas.length
  let balance = 0
  let debt = 0
  let start = 0

  for (let i = 0; i < len; i++) {
    balance += gas[i] - cost[i]
    debt += gas[i] - cost[i]
    if (balance < 0) {
      balance = 0
      start = i + 1
    }
  }

  if (debt < 0) {
    return -1
  }

  return start
};

console.log(canCompleteCircuit([1,2,3,4,5], [3,4,5,1,2])) // 3
console.log(canCompleteCircuit([2,3,4], [3,4,3])) // -1