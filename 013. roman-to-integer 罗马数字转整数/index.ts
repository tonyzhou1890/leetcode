function romanToInt(s: string): number {
  const table: {
    [x: string]: number
  } = {
    'I': 1,
    'V': 5,
    'X': 10,
    'L': 50,
    'C': 100,
    'D': 500,
    'M': 1000
  }
  let num = 0
  const len = s.length
  for (let i = 0; i < len; i++) {
    if (i < len - 1 && table[s[i]] < table[s[i + 1]]) {
      num += table[s[i + 1]] - table[s[i]]
      i++
    } else {
      num += table[s[i]]
    }
  }
  return num
};

console.log(romanToInt("III")) // 3
console.log(romanToInt("IV")) // 4
console.log(romanToInt("IX")) // 9
console.log(romanToInt("LVIII")) // 58
console.log(romanToInt("MCMXCIV")) // 1994

// I             1
// V             5
// X             10
// L             50
// C             100
// D             500
// M             1000