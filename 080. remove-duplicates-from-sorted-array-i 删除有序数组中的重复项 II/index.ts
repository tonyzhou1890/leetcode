function removeDuplicates(nums: number[]): number {
  const len = nums.length
  let p = 1
  for (let i = 2; i < len; i++) {
    if (nums[i] === nums[p] && nums[p] === nums[p - 1]) {
      continue
    }
    nums[++p] = nums[i]
  }
  return p + 1
};

console.log(removeDuplicates([1,1,1,2,2,3]))
console.log(removeDuplicates([0,0,1,1,1,1,2,3,3]))