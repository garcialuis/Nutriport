package models

type Person struct {
	Weight         float64 `json:"weight,omitempty"`
	Height         float64 `json:"height,omitempty"`
	Gender         uint    `json:"gender,omitempty"`
	Age            uint    `json:"age,omitempty"`
	BMI            float64 `json:"BMI,omitempty"`
	BMIDescription string  `json:"BMIDescription,omitempty"`
	TEE            float64 `json:"TEE,omitempty"`
	ActivityLevel  string  `json:"activityLevel,omitempty"`
}
