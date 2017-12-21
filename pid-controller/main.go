package main

import "fmt"

type PIDController struct {
	// proportional gain
	kp float64
	// integral gain
	ki       float64
	integral float64
	// derivative gain
	kd float64
	// previouse value
	value     float64
	frequency float64
}

func NewPIDController(kp float64, ki float64, kd float64) *PIDController {
	pidc := new(PIDController)
	pidc.kp = kp
	pidc.ki = ki
	pidc.integral = 0.0
	pidc.kd = kd
	pidc.frequency = 1.0
	return pidc
}

func (pidc *PIDController) Iteration(value float64) (result float64) {
	value = value - pidc.value
	pidc.integral += value
	derivative := (value - pidc.value) * pidc.frequency
	result = (pidc.kp * value)
	result += (pidc.ki * pidc.integral / pidc.frequency)
	result += (pidc.kd * derivative)
	pidc.value = value
	return result
}

func main() {
	controller := NewPIDController(1.2, 0.003, 1)
	o := 5.0
	for {
		o := controller.Iteration(o)
		fmt.Println(o)
	}
}
