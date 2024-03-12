function reverseWords(s: string): string {
  let word = ''
  let res = ''
  for (let i = s.length - 1; i >= 0; i--) {
    if (!word && s[i] === ' ') {
      continue
    }
    if (word && s[i] === ' ') {
      res += word + ' '
      word = ''
    }
    if (s[i] !== ' ') {
      word = s[i] + word
    }
  }
  if (word) {
    res += word
  }
  return res.trim()
};

console.log(reverseWords("the sky is blue")) // "blue is sky the"
console.log(reverseWords("  hello world  ")) // "world hello"
console.log(reverseWords("a good   example")) // "example good a"