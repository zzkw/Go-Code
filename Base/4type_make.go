package main

import (
	"fmt"
)

const (
	usixteenbitmax float64 = 65535
	kmh_multiple   float64 = 1.60934
)

type car struct {
	gas_pedal      uint16 //min 0  max 65535
	brake_p        uint16
	steering_wheel int16 //-32k ~ +32k
	top_speed_kmh  float64
}

func (c car) kmh() float64 {
	return float64(c.gas_pedal) * (c.top_speed_kmh / usixteenbitmax)
}
func (c *car) new_top_speed(newspeed float64) {
	c.top_speed_kmh = newspeed
}
func main() {
	a_car := car{65000, 0, 1264, 253.154}
	fmt.Println(a_car.gas_pedal)
	fmt.Println(a_car.kmh())
	a_car.new_top_speed(500)
	fmt.Println(a_car.kmh())
	s1 := make([]int, 3, 10)
	fmt.Printf("%v %p\n", s1, s1) //这里注意只能Printf。。。%v值 %p地址
}
