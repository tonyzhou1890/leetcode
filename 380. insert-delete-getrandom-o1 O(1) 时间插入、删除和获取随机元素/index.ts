class RandomizedSet {
  constructor() {
    
  }

  store: { [x: number]: number } = {}
  arr: any[] = []
  len = 0

  insert(val: number): boolean {
    if (this.store[val] !== undefined) {
      return false
    }
    this.arr[this.len] = val
    this.len++
    this.store[val] = this.len - 1

    return true
  }

  remove(val: number): boolean {
    if (this.store[val] === undefined) return false

    // 将数组最后一个元素移动到要删除的元素位置，保持连续性
    const lastVal = this.arr[this.len - 1]
    this.arr[this.store[val]] = lastVal
    this.store[lastVal] = this.store[val]
    this.len--
    delete this.store[val]

    return true
  }

  getRandom(): number {
    return this.arr[(Math.random() * this.len) >> 0]
  }
}

// const randomizedSet = new RandomizedSet()
// randomizedSet.insert(1); // 向集合中插入 1 。返回 true 表示 1 被成功地插入。
// randomizedSet.remove(2); // 返回 false ，表示集合中不存在 2 。
// randomizedSet.insert(2); // 向集合中插入 2 。返回 true 。集合现在包含 [1,2] 。
// console.log(randomizedSet.getRandom()); // getRandom 应随机返回 1 或 2 。
// randomizedSet.remove(1); // 从集合中移除 1 ，返回 true 。集合现在包含 [2] 。
// randomizedSet.insert(2); // 2 已在集合中，所以返回 false 。
// console.log(randomizedSet.getRandom()); // 由于 2 是集合中唯一的数字，getRandom 总是返回 2 。

const randomizedSet = new RandomizedSet()
randomizedSet.insert(0);
randomizedSet.insert(1);
randomizedSet.remove(0);
randomizedSet.insert(2);
randomizedSet.remove(1);
console.log(randomizedSet.getRandom()); // 2