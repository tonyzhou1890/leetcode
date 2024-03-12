function intToRoman(num: number): string {
  const table: [number, string][] = [
    [1000, 'M'],
    [900, 'CM'],
    [500, 'D'],
    [400, 'CD'],
    [100, 'C'],
    [90, 'XC'],
    [50, 'L'],
    [40, 'XL'],
    [10, 'X'],
    [9, 'IX'],
    [5, 'V'],
    [4, 'IV'],
    [1, 'I']
  ]

  let s = ''
  let i = 0
  while (num) {
    if (num >= table[i][0]) {
      s += table[i][1]
      num -= table[i][0]
    } else {
      i++
    }
  }

  return s
};

// I             1
// V             5
// X             10
// L             50
// C             100
// D             500
// M             1000

console.log(intToRoman(3)) // "III"
console.log(intToRoman(4)) // "IV"
console.log(intToRoman(58)) // "LVIII"
console.log(intToRoman(1994)) // "MCMXCIV"