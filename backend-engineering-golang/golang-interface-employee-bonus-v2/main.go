package main

type Employee interface {
	GetBonus() float64
}

type Junior struct {
	Name         string
	BaseSalary   int
	WorkingMonth int
}

func prorata(month int) float64 {
	if month > 12 {
		return 1
	}
	return float64(month) / 12
}

func (j Junior) GetBonus() float64 {
	prorata := prorata(j.WorkingMonth)
	return float64(j.BaseSalary) * prorata
}

type Senior struct {
	Name            string
	BaseSalary      int
	WorkingMonth    int
	PerformanceRate float64
}

func (s Senior) GetBonus() float64 {
	prorata := prorata(s.WorkingMonth)
	return float64(2*s.BaseSalary)*prorata + (s.PerformanceRate * float64(s.BaseSalary))
}

type Manager struct {
	Name             string
	BaseSalary       int
	WorkingMonth     int
	PerformanceRate  float64
	BonusManagerRate float64
}

func (m Manager) GetBonus() float64 {
	prorata := prorata(m.WorkingMonth)
	return float64(2*m.BaseSalary)*prorata + (m.PerformanceRate * float64(m.BaseSalary)) + (m.BonusManagerRate * float64(m.BaseSalary))
}

func EmployeeBonus(employee Employee) float64 {
	return employee.GetBonus()
}

func TotalEmployeeBonus(employees []Employee) float64 {
	var total float64
	for _, v := range employees {
		total += v.GetBonus()
	}
	return total
}

func main() {

}
