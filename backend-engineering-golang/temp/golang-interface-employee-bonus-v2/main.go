package main

type Employee interface {
	GetBonus() float64
}

type Junior struct {
	Name                     string
	BaseSalary, WorkingMonth int
}

type Senior struct {
	Name                     string
	BaseSalary, WorkingMonth int
	PerformanceRate          float64
}

func (j Junior) GetBonus() float64 {
	var prorta float64 = 1
	if j.WorkingMonth > 12 {
		prorta = 1
	} else {
		prorta = float64(j.WorkingMonth) / 12
	}

	return float64(j.BaseSalary) * prorta
}

type Manager struct {
	Name                              string
	BaseSalary, WorkingMonth          int
	PerformanceRate, BonusManagerRate float64
}

func (s Senior) GetBonus() float64 {
	var prorta float64 = 1
	if s.WorkingMonth > 12 {
		prorta = 1
	} else {
		prorta = float64(s.WorkingMonth) / 12
	}
	return float64(2*s.BaseSalary)*prorta + (s.PerformanceRate * float64(s.BaseSalary))
}

func (m Manager) GetBonus() float64 {
	var prorta float64 = 1
	if m.WorkingMonth > 12 {
		prorta = 1
	} else {
		prorta = float64(m.WorkingMonth) / 12
	}
	return float64(2*m.BaseSalary)*prorta + (m.PerformanceRate * float64(m.BaseSalary)) + (m.BonusManagerRate * float64(m.BaseSalary))
}

func EmployeeBonus(employee Employee) float64 {
	var bonus = employee.GetBonus()
	return bonus
}

func TotalEmployeeBonus(employees []Employee) float64 {
	var total float64
	for _, v := range employees {
		total = total + v.GetBonus()
	}
	return total
}
