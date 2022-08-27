# golang_valid_parentheses

Given a string `s` containing just the characters `'('`, `')'`, `'{'`, `'}'`, `'['` and `']'`, determine if the input string is valid.

An input string is valid if:

1. Open brackets must be closed by the same type of brackets.
2. Open brackets must be closed in the correct order.
3. Every close bracket has a corresponding open bracket of the same type.

## Examples

**Example 1:**

```
Input: s = "()"
Output: true

```

**Example 2:**

```
Input: s = "()[]{}"
Output: true

```

**Example 3:**

```
Input: s = "(]"
Output: false

```

**Constraints:**

- `1 <= s.length <= $10^4$`
- `s` consists of parentheses only `'()[]{}'`.

## 解析

給定一個字串 s, 這個字串只會由 '(', ')', '[', ']','{','}' 六種字元所組成

定義字串 s 是 valid 如果符合以下條件：

1  對於每個 '(', '[', '{' 三種字元，找得到在其右方找到對應的 ')',']','}' 來做配對

2  每個 ')',']','}' 三種字元， 找在其左方找到對應的 '(', '[', '{' 來做配對

3  每個字元只能用來做一次配對

4 每個配對位置順需必須正確

在這種情況下假設字串 s 是合法的, 可以發現對於每種字元做計數

可以發現每個位置 的 '(', '[', '{'  個數一定要大於等於 ')',']','}'

因為要能配對成功 ')',']','}' 必須在左方找到 ')',']','}'

且因為配對順序必須正確代表 後出現的字元必須要先配對

所以可以採用 stack 來蒐集 '(', '[', '{'， 當遇到 ')',']','}' 就 pop 出 stack 最上方字元做比對

![](https://i.imgur.com/F82ATak.png)

如果發現 stack 最上方字元 比對不到正確的配對 則返回false

![](https://i.imgur.com/sLFQ5g6.png)

如果到最後都比對正確代表符合 valid 且 stack is empty 則代表可以完全 match 返回 true

作法如下圖

![](https://i.imgur.com/DaSt27O.png)


## 程式碼
```go
package sol

func isValid(s string) bool {
	if len(s) <= 1 {
		return false
	}
	stack := []rune{}
	popMap := map[rune]rune{')': '(', ']': '[', '}': '{'}
	for _, val := range s {
		if popVal, exist := popMap[val]; !exist {
			stack = append(stack, val)
		} else {
			if len(stack) == 0 {
				return false
			}
			topVal := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if topVal != popVal {
				return false
			}
		}
	}
	return len(stack) == 0
}

```
## 困難點

1. 需要理解比對的方式具有後進先出的概念

## Solve Point

- [x]  初始化一個 stack := []byte{}
- [x]  從 pos := 0..len(s)-1 做以下運算
- [x]  當遇到 s[pos] = '(', '[', '{' 時，把 s[pos] push 到 stack 之內
- [x]  當遇到 s[pos] = ')',']','}' 時 做以下運算
- [x]  如果 stack size == 0 則代表沒有足夠配對的字元,返回 false
- [x]  把 stack pop 出一個字元 檢查這個字元是不是 s[pos] 的對應配對
- [x]  如果不是 則返回 false
- [x]  當走完所有 pos ，如果 stack size == 0 則代表都比對完了返回 true, 否則 false