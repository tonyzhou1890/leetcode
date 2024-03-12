function longestCommonPrefix(strs: string[]): string {
  const len = strs.length
  const firstStrLen = strs[0].length
  let prefix = ''

  for (let i = 0; i < firstStrLen; i++) {
    let flag = true
    for (let j = 1; j < len; j++) {
      if (strs[j][i] !== strs[0][i]) {
        flag = false
        break
      }
    }
    if (flag) {
      prefix += strs[0][i]
    } else {
      break
    }
  }
  return prefix
};

console.log(longestCommonPrefix(["flower","flow","flight"])) // 'fl'
console.log(longestCommonPrefix(["dog","racecar","car"])) // ''