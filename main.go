package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Đọc hai số từ bàn phím
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Nhập số thứ nhất: ")
	num1Str, _ := reader.ReadString('\n')
	fmt.Print("Nhập số thứ hai: ")
	num2Str, _ := reader.ReadString('\n')

	// Chuyển chuỗi đọc được từ bàn phím thành số
	num1, _ := strconv.ParseInt(num1Str, 10, 64)
	num2, _ := strconv.ParseInt(num2Str, 10, 64)

	// Tính trung bình cộng và in kết quả ra màn hình
	avg := (num1 + num2) / 2
	fmt.Println("Trung bình cộng của hai số là:", avg)
}
