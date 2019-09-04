/*
 * @lc app=leetcode.cn id=67 lang=golang
 *
 * [67] 二进制求和
 */
func addBinary(a string, b string) string {
    if len(a) == 0 {
		return b
	}
	if len(b) == 0 {
		return a
	}

	i := 1
	needCarry := false
	ret := ""

	for len(a) - i >= 0 || len(b) - i >= 0 {
		aLoc := len(a) - i
		bLoc := len(b) - i

		aByte := "0"
		bByte := "0"

		if aLoc >= 0 {
			aByte = a[aLoc: aLoc+1]
		}
		if bLoc >= 0 {
			bByte = b[bLoc: bLoc+1]
		}

		aNum := 1
		bNum := 1

		if aByte == "0" {
			aNum = 0
		}
		if bByte == "0" {
			bNum = 0
		}

		num := aNum + bNum
		if needCarry {
			num += 1
		}

		if num <= 1 {
			needCarry = false
		} else {
			needCarry = true
		}

		curNum := num % 2
		if curNum == 0 {
			ret = "0" + ret
		} else if curNum == 1 {
			ret = "1" + ret
		}

		i++
	}
	
	if needCarry {
		ret = "1" + ret
	}
	return ret
}

