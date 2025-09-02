// Lab: Abstraction with Go (Car Example)
// --------------------------------------
// ตัวอย่างนี้สาธิตการใช้งาน Abstraction ผ่าน Interface ชื่อ Car
// โดยมีรถหลายประเภท (Sedan, Pickup, EV) ที่ implements เมธอดเดียวกัน
//
// รันโปรแกรม:
//   go run main.go
//
// ผู้เขียน: You 🧑‍💻
package main

import "fmt"

// Car คือ interface ที่เป็นนามธรรมของ "รถ"
// ทุก struct ที่จะถือว่าเป็นรถ ต้องมีเมธอด Drive() และ Info()
type Car interface {
	Drive()       // เมธอดสำหรับแสดงพฤติกรรมการขับรถ
	Info() string // เมธอดสำหรับคืนค่าข้อมูลรถ
}

// -------------------- Sedan --------------------
// Sedan คือ struct ที่แทนรถเก๋ง
type Sedan struct {
	Brand string
	Model string
	Color string
}

// implements เมธอด Drive() ของ Car
func (s Sedan) Drive() {
	fmt.Println("ขับรถเก๋ง ->", s.Brand, s.Model, "สี:", s.Color)
}

// implements เมธอด Info() ของ Car
func (s Sedan) Info() string {
	return fmt.Sprintf("Sedan(%s %s, %s)", s.Brand, s.Model, s.Color)
}

// -------------------- Pickup --------------------
// Pickup คือ struct ที่แทนรถกระบะ
type Pickup struct {
	Brand   string
	Payload int  // น้ำหนักบรรทุกสูงสุด (kg)
	FourWD  bool // มีระบบขับเคลื่อน 4 ล้อหรือไม่
}

func (p Pickup) Drive() {
	mode := "2WD"
	if p.FourWD {
		mode = "4WD"
	}
	fmt.Println("ขับรถกระบะ ->", p.Brand, "บรรทุกได้:", p.Payload, "kg", "| โหมด:", mode)
}

func (p Pickup) Info() string {
	return fmt.Sprintf("Pickup(%s, %dkg, FourWD:%t)", p.Brand, p.Payload, p.FourWD)
}

// -------------------- EV --------------------
// EV คือ struct ที่แทนรถยนต์ไฟฟ้า
type EV struct {
	Brand      string
	Battery    int  // ความจุแบตเตอรี่ (kWh)
	RangeKM    int  // ระยะทางที่วิ่งได้ต่อการชาร์จ (km)
	FastCharge bool // รองรับชาร์จเร็วหรือไม่
}

func (e EV) Drive() {
	fc := "ชาร์จปกติ"
	if e.FastCharge {
		fc = "รองรับชาร์จเร็ว"
	}
	fmt.Println("ขับรถไฟฟ้า ->", e.Brand, "| แบตเตอรี่:", e.Battery, "kWh | ระยะทาง:", e.RangeKM, "km |", fc)
}

func (e EV) Info() string {
	return fmt.Sprintf("EV(%s, %dkWh, %dkm, FastCharge:%t)", e.Brand, e.Battery, e.RangeKM, e.FastCharge)
}

// -------------------- main --------------------
func main() {
	// สร้าง object ของรถแต่ละประเภท
	sedan := Sedan{Brand: "Toyota", Model: "Camry", Color: "ดำ"}
	pickup := Pickup{Brand: "Isuzu", Payload: 1000, FourWD: true}
	ev := EV{Brand: "Tesla", Battery: 75, RangeKM: 500, FastCharge: true}

	// ใช้งาน abstraction ผ่าน interface Car
	cars := []Car{sedan, pickup, ev}

	// วนลูปเรียกใช้งาน โดยไม่สนใจว่ารถเป็นชนิดใด
	for _, car := range cars {
		fmt.Println(car.Info())
		car.Drive()
		fmt.Println("---")
	}
}
