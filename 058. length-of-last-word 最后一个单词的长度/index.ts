function lengthOfLastWord(s: string): number {
  let count = 0
  for (let i = s.length - 1; i >= 0; i--) {
    if (!count && s[i] === ' ') continue
    if (count && s[i] === ' ') break
    count++
  }
  return count
};

console.log(lengthOfLastWord("Hello World")) // 5
console.log(lengthOfLastWord("   fly me   to   the moon  ")) // 4
console.log(lengthOfLastWord("luffy is still joyboy")) // 6