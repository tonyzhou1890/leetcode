/**
 * 双指针
 * @param s 
 * @param t 
 */
function isSubsequence(s: string, t: string): boolean {
  if (!s) return true
  for (let si = 0, ti = 0, slen = s.length, tlen = t.length; ti < tlen; ti++) {
    if (s[si] === t[ti]) {
      si++
      if (si >= slen) {
        return true
      }
    }
  }
  return false
};

console.log(isSubsequence("abc", "ahbgdc") === true) // true
console.log(isSubsequence("axc", "ahbgdc") === true) // false