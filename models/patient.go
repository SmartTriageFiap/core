package models

type Patient struct {
	Salt     string `json:"salt"`
	Cpf      string `json:"cpf"`
	FullName string `json:"full_name"`
	Phone    string `json:"phone"`
}

type PatientData struct {
	Patient    Patient `json:"patient"`
	EmployeeId string  `json:"employee_id"`
	UnityId    string  `json:"unity_id"`
}
