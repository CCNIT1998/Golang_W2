package main

import (
	"sort"
	"strings"
	"time"
	"github.com/bearbin/go-age"
	"strconv"
)

type job struct {
	Name  string
	Number int
}

type city struct {
	Name    string
	Number int
}

// 2.1 Gom tất cả những người trong cùng một thành phố lại

func GroupPeopleByCity(p []Person) (result map[string][]Person) {
	result = make(map[string][]Person)
	for _, person := range p {
		result[person.City] = append(result[person.City], person)
	}

	return result
}

// 2.2 Nhóm các nghề nghiệp và đếm số người làm

func GroupPeopleByJob(p []Person) (result map[string]int) {
	result = make(map[string]int)
	for _, product := range p {
		result[product.Job]++
	}
	return result
}

// 2.3 Tìm 5 nghề có nhiều người làm nhất, đếm từ cao xuống thấp

func Top5JobsByNumer(m map[string]int) (result []job) {
	for key, value := range m {
		result = append(result, job{key, value})
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].Number > result[j].Number
	})
	result = result[0:5]
	return result
}

// 2.4 Tìm 5 thành phố có nhiều người trong danh sách ở nhất, đếm từ cao xuống thấp

func Top5CitiesByNumber(p []Person) (result []city) {
	temp := make(map[string]int)
	for _, product := range p {
		temp[product.City]++
	}
	for key, value := range temp {
		result = append(result, city{key, value})
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].Number > result[j].Number
	})
	result = result[0:5]
	return result
}

// 2.5 Trong mỗi thành phố, hãy tìm ra nghề nào được làm nhiều nhất

func CountJob(items []string) map[string]int {
	result := make(map[string]int)
	for _, itm := range items {
		result[itm]++
	}
	return result
}

func jobNumberMax(input map[string]int) (result map[string]int) {
	result = make(map[string]int)
	max := 0
	for _, value := range input {
		if value > max {
			max = value
		}
	}
	for key, value := range input {
		if value == max {
			result[key] = value
		}
	}
	return result
}

func TopJobByNumberInEachCity(p []Person) (result map[string]map[string]int) {
	txtCity := make(map[string][]string)
	proccessTxtCity := make(map[string]map[string]int)
	for _, person := range p {
		txtCity[person.City] = append(txtCity[person.City], person.Job)
	}
	for key, value := range txtCity {
		proccessTxtCity[key] = CountJob(value)
	}

	result = make(map[string]map[string]int)
	for key, value := range proccessTxtCity {
		result[key] = jobNumberMax(value)
	}
	return result
}

// 2.6 Ứng với một nghề, hãy tính mức lương trung bình

func AverageSalaryByJob(p []Person) (result map[string]int) {
	result = make(map[string]int)
	totalJob := GroupPeopleByJob(p)
	totalSalary := make(map[string]int)
	for _, person := range p {
		totalSalary[person.Job] += person.Salary
	}
	for key := range totalJob {
		result[key] = totalSalary[key] / totalJob[key]
	}
	return result
}

// 2.7 Năm thành phố có mức lương trung bình cao nhất

func FiveCitiesHasTopAverageSalary(p []Person) (result []job){
	data := AverageSalaryByJob(p)
	for key, value := range data {
		result = append(result, job{key, value})
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].Number > result[j].Number
	})
	result = result[0:5]
	return result
}

// 2.8 Năm thành phố có mức lương trung bình của developer cao nhất
func FiveCitiesHasTopSalaryForDeveloper(p []Person) (result []city){ 
	totalDeveloperEachCity := make(map[string]int)
	totalSalaryDeveloperByCityEachCity := make(map[string]int)
	for _, person := range p {
		if person.Job == "developer" {
			totalDeveloperEachCity[person.City]++
			totalSalaryDeveloperByCityEachCity[person.City]+= person.Salary
		}
	}
	averageSalaryDeveloper := make(map[string]int)
	for key := range totalDeveloperEachCity {
		averageSalaryDeveloper[key] = totalSalaryDeveloperByCityEachCity[key] / totalDeveloperEachCity[key]
	}
	// averageSalaryDeveloper
	var txtResult []city
	for key, value := range averageSalaryDeveloper {
		txtResult = append(txtResult, city{key, value})
		// fmt.Println(key, value)
	}
	sort.Slice(txtResult, func(i, j int) bool {
		return txtResult[i].Number > txtResult[j].Number
	})
	temp := len(txtResult)
	if temp<=5{
		result = txtResult
	}else{
		result = txtResult[0:5]
	}
	return result
}

// 2.9 Tuổi trung bình từng nghề nghiệp
func getAge(year, month, day int) time.Time {
    Age := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
    return Age
}

func AverageAgePerJob(p []Person) (result map[string]float64){ 
	totalJobEachCity := make(map[string]float64)
	totalAgeEachJob := make(map[string]float64)

	for _, person := range p {
		date := person.Birthday
		txt := strings.Split(date, "-")
		txtY, _ := strconv.Atoi(txt[0])
		txtM, _ := strconv.Atoi(txt[1])
		txtD, _ := strconv.Atoi(txt[2])

		agePerson := getAge(txtY, txtM,txtD)
    	// fmt.Printf("Age is %d\n", age.Age(agePerson))
		totalJobEachCity[person.Job]++
		totalAgeEachJob[person.Job]+= float64(age.Age(agePerson))
	}
	result = make(map[string]float64)
	for key := range totalJobEachCity {
		result[key] = totalAgeEachJob[key] / totalJobEachCity[key]
	}
	return result
}


// 2.10 Tuổi trung bình ở từng thành phố

func AverageAgePerCity(p []Person) (result map[string]float64){ 
	totalNumberEachCity := make(map[string]float64)
	totalAgeEachCity := make(map[string]float64)

	for _, person := range p {
		date := person.Birthday
		txt := strings.Split(date, "-")
		txtY, _ := strconv.Atoi(txt[0])
		txtM, _ := strconv.Atoi(txt[1])
		txtD, _ := strconv.Atoi(txt[2])

		agePerson := getAge(txtY, txtM,txtD)
    	// fmt.Printf("Age is %d\n", age.Age(agePerson))
		totalNumberEachCity[person.City]++
		totalAgeEachCity[person.City]+= float64(age.Age(agePerson))
	}
	result = make(map[string]float64)
	for key := range totalNumberEachCity {
		result[key] = totalAgeEachCity[key] / totalNumberEachCity[key]
	}
	return result
}