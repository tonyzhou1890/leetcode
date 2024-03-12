function isPalindrome(s: string): boolean {
  let i = 0
  let j = s.length - 1
  let flag = true

  while (i < j) {
    if (!valid(s[i])) {
      i++
    } else if (!valid(s[j])) {
      j--
    } else if (s[i].toLocaleLowerCase() !== s[j].toLocaleLowerCase()) {
      return false
    } else {
      i++
      j--
    }
  }

  return flag

  function valid(c: string) {
    return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
  }
};

console.log(isPalindrome("A man, a plan, a canal: Panama")) // true
console.log(isPalindrome("race a car")) // false
console.log(isPalindrome(" ")) // true
console.log(isPalindrome("0P")) // false