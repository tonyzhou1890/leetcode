function convert(s: string, numRows: number): string {
  if (numRows <= 1) return s
  const arr = new Array(numRows).fill('')
  const len = s.length

  for (let i = 0, r = 0, step = 1; i < len; i++) {
    arr[r] += s[i]
    if (r === numRows - 1 && step === 1) {
      step = -1
    }
    if (r === 0 && step === -1) {
      step = 1
    }
    r += step
  }

  return arr.join('')
};

console.log(convert("PAYPALISHIRING", 3) === 'PAHNAPLSIIGYIR')
console.log(convert("PAYPALISHIRING", 4) === 'PINALSIGYAHRPI')