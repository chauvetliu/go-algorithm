package main

import (
	"fmt"
	"strconv"
)

//思路：
//1、每个整数都可以归纳为三种情况：1位数字；2位数字；3位数字，然后递归即可。
//2、统计切分次数，且只能切分3次，所以第四次切分时，剩余的字符串必为空。

func restoreIpAddress(s string) (res []string) {
	res = restore(0, "", s, res)
	fmt.Println("restoreIpAddress", res)
	return res
}

// count record split times, ip record ip, s record remaining string
func restore(count int, ip, s string, res []string) []string {

	if count == 4 {
		if s == "" {
			res = append(res, ip[:len(ip)-1]) //len(ip)-1是为了删掉末尾多余的“.”
			return res
		}
	} else if count < 4 { // 切分次数小于4时
		//每个整数位数总共有三种情况: 1位；2位；3位
		if len(s) > 0 { //1位数
			res = restore(count+1, ip+string(s[0])+".", s[1:], res)
		}
		if len(s) > 1 && string(s[0]) != "0" { //2位数，且第一位非0
			res = restore(count+1, ip+string(s[:2])+".", s[2:], res)
		}
		if len(s) > 2 && string(s[0]) != "0" { //三位数，第一位非0，且必须小于256
			number, _ := strconv.Atoi(s[0:3]) //
			if number < 256 {
				// fmt.Println("string(s[:3])", s)
				res = restore(count+1, ip+string(s[:3])+".", s[3:], res)
			}
		}
	}
	return res
}

func main() {
	restoreIpAddress("1921681101")
}
