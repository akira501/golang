package main

import (
	"fmt"
)

func main() {
	ListStudent := []SinhVien{
		{"Nguyễn Thị Mộng Mơ", 9, 10, 0},
		{"Bùi Như Lạc", 6, 5, 0},
		{"Đỗ Trần Nước Sôi", 7, 8, 0},
		{"Trần Như Nhộng", 8, 8, 0},
		{"Đào Nguyên Hạt", 4, 6, 0},
	}

	n := len(ListStudent)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if ListStudent[j].TinhDiem() < ListStudent[j+1].TinhDiem() {
				temp := ListStudent[j]
				ListStudent[j] = ListStudent[j+1]
				ListStudent[j+1] = temp
			}
		}
	}

	for k := 0; k < len(ListStudent); k++ {
		fmt.Println(" Sinh Vien ", ListStudent[k].FullName, " đứng thứ ", k+1, " đạt điểm tổng kết ", ListStudent[k].DiemTongKet)
	}

}

type SinhVien struct {
	FullName                            string
	DiemGiuaKy, DiemCuoiKy, DiemTongKet int
}

func (s *SinhVien) TinhDiem() int {
	s.DiemTongKet = ((s.DiemCuoiKy * 2) + s.DiemGiuaKy) / 3
	return s.DiemTongKet
}
