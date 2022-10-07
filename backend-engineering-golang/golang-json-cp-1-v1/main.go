package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var scale = map[string]float64{
	"A":  4.0,
	"AB": 3.5,
	"B":  3.0,
	"BC": 2.5,
	"C":  2.0,
	"CD": 1.5,
	"D":  1.0,
	"DE": 0.5,
	"E":  0.0,
}

type Study struct {
	StudyName   string `json:"study_name"`
	StudyCredit int    `json:"study_credit"`
	Grade       string `json:"grade"`
}
type Report struct {
	Id       string  `json:"id"`
	Name     string  `json:"name"`
	Date     string  `json:"date"`
	Semester int     `json:"semester"`
	Studies  []Study `json:"studies"`
}

// gunakan fungsi ini untuk mengambil data dari file json
// kembalian berupa struct 'Report' dan error
func ReadJSON(filename string) (Report, error) {
	file, err := os.Open(filename)
	if err != nil {
		return Report{}, err
	}

	var report Report
	err = json.NewDecoder(file).Decode(&report)
	if err != nil {
		return Report{}, err
	}
	return report, nil
}

func GradePoint(report Report) float64 {
	if len(report.Studies) == 0 {
		return 0.0
	}

	totalNilaiMK := 0.0
	totalCredit := 0.0

	for _, v := range report.Studies {
		totalNilaiMK += float64(v.StudyCredit) * scale[v.Grade]
		totalCredit += float64(v.StudyCredit)
	}

	return totalNilaiMK / totalCredit
}

func main() {
	// bisa digunakan untuk menguji test case
	report, err := ReadJSON("report.json")
	if err != nil {
		panic(err)
	}

	gradePoint := GradePoint(report)
	fmt.Println(gradePoint)
}
