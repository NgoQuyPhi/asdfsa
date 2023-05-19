package main

import "fmt"

type MonHoc struct {
	Subject_ID   int
	Subject_Name string
	SinhVien
}

type SinhVien struct {
	ID          int
	Name        string
	PhoneNumber string
	Address     string
}

type I interface {
	describe()
}

func (a SinhVien) describe() {
	fmt.Print("ID: ", a.ID, "\n")
	fmt.Print("FullName: " + a.Name + "\n")
	fmt.Print("PhoneNumber " + a.PhoneNumber + "\n")
	fmt.Print("Address: " + a.Address + "\n")

}

func (a *SinhVien) InsertSinhVienInformation() {
	fmt.Println("Your ID:")
	fmt.Scan(&a.ID)
	fmt.Println("Your Name: ")
	fmt.Scan(&a.Name)
	fmt.Println("Your PhoneNumber:")
	fmt.Scan(&a.PhoneNumber)
	fmt.Println("Your Address:")
	fmt.Scan(&a.Address)
}

func CreateAList(sinhvien []SinhVien) {
	var NumberOfStudents int
	fmt.Println("Nhap vao so sinh vien trong danh sach:")
	fmt.Scan(&NumberOfStudents)
	var Student = make([]SinhVien, NumberOfStudents)
	//Nhap thong tin cho danh sach sinh vien
	for i, _ := range Student {
		Student[i].InsertSinhVienInformation()
	}
	sinhvien = append(sinhvien, Student...)
}

func Menu_Action(a []SinhVien, b *bool) {
	fmt.Println("__________________MENU_____________________")
	fmt.Println("|1:Them 1 so sinh vien vao danh sach      |")
	fmt.Println("|2:Xoa 1 so Sinh vien khoi danh sach      |")
	fmt.Println("|3:Thay Doi thong tin cua sinh vien       |")
	fmt.Println("|4:Hien Thi Danh Sach Sinh Vien           |")
	fmt.Println("|5:Thoat                                  |")
	fmt.Println("|_________________________________________|")
	var Action int8
	fmt.Println("Nhap lua chon cua ban")
	fmt.Scan(&Action)
	switch Action {
	case 1:
		var NumberOfStudents int
		fmt.Println("Nhap vao so sinh vien trong danh sach:")
		fmt.Scan(&NumberOfStudents)
		var Student = make([]SinhVien, NumberOfStudents)
		//Nhap thong tin cho danh sach sinh vien
		for i, _ := range Student {
			Student[i].InsertSinhVienInformation()
		}
		a = append(a, Student...)
		*b = false
		break
	case 2:
		fmt.Printf("Nhap ID Cua Sinh Vien Can Xoa")
		var temp int
		fmt.Scan(&temp)
		for i, _ := range a {
			if a[i].ID == temp {
				a = append(a[:i], a[i+1:]...)
			}
		}
		*b = false
		break
	case 3:
		fmt.Printf("Nhap ID Cua Sinh Vien Can Thay Doi Thong Tin")
		var temp2 int
		fmt.Scan((&temp2))
		for i, _ := range a {
			if a[i].ID == temp2 {
				a[i].InsertSinhVienInformation()
			}
		}
		*b = false
		break
	case 4:
		for i, _ := range a {
			a[i].describe()
		}
		*b = false
		break
	case 5:
		*b = true
		break
	}

}

func main() {
	var List []SinhVien
	t := &List
	var check bool = false
	for check == false {
		Menu_Action(*t, &check)
	}

}
