function strStr(haystack: string, needle: string): number {
  let hi = needle.length - 1
  let ni = hi
  while (hi < haystack.length && ni < needle.length) {
    // 相等，从后往前继续比较
    if (haystack[hi] === needle[ni]) {
      if (ni === 0) {
        return hi
      }
      hi--
      ni--
    } else {
      ni--
      // 不等，从当前 ni 位置往前搜索 needle 匹配的字符，然后移动字符串
      while (ni >= 0) {
        if (haystack[hi] === needle[ni]) {
          break
        }
        ni--
      }
      hi += needle.length - ni - 1
      ni = needle.length - 1
    }
  }
  return -1
};

console.log(strStr("sadbutsad", 'sad') === 0)
console.log(strStr("leetcode", 'leeto') === -1)