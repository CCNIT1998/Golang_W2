package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

/* https://stackoverflow.com/questions/45303326/how-to-parse-non-standard-time-format-from-json
"name":"Dee Leng",
"email":"dleng0@cocolog-nifty.com",
"job":"developer",
"gender":"Female",
"city":"London",
"salary":9662,
"birthdate":"2007-09-30" */
type Person struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Job      string `json:"job"`
	City     string `json:"city"`
	Salary   int    `json:"salary"`
	Birthday string `json:"birthdate"`
}

func (p *Person) String() string {
	return fmt.Sprintf("name: %s, email: %s, job: %s, city: %s, salary: %d, birthday: %s",
		p.Name, p.Email, p.Job, p.City, p.Salary, p.Birthday)
}

func main() {
	// Open our jsonFile
	jsonFile, err := os.Open("personsmall.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened person.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var people []Person

	json.Unmarshal(byteValue, &people)

	/*
		for i := 0; i < 10; i++ {
			fmt.Println(&people[i])
		}
	*/

	// 2.1 Gom tất cả những người trong cùng một thành phố lại
	peopleByCity := GroupPeopleByCity(people)
	for key, value := range peopleByCity {
		fmt.Println(key)
		for _, person := range value {
			fmt.Println("  ", (&person).Name)
		}
	}

	// 2.2 Nhóm các nghề nghiệp và đếm số người làm
	fmt.Println("_________ 2.2 Nhóm các nghề nghiệp và đếm số người làm _________")
	peopleByJob := GroupPeopleByJob(people)
	fmt.Println(peopleByJob)
	for key, value := range peopleByJob {
		fmt.Printf("%s : %d\n", key, value)
	}

	// 2.3 Tìm 5 nghề có nhiều người làm nhất, đếm từ cao xuống thấp
	fmt.Println("_________ 2.3 Tìm 5 nghề có nhiều người làm nhất, đếm từ cao xuống thấp _________")
	T5JBN := Top5JobsByNumer(peopleByJob)
	fmt.Println(T5JBN)
	for _,value := range T5JBN{
		fmt.Printf("%s : %d\n", value.Name, value.Number)
	}
	// 2.4 Tìm 5 thành phố có nhiều người trong danh sách ở nhất, đếm từ cao xuống thấp
	fmt.Println("_________ 2.4 Tìm 5 thành phố có nhiều người trong danh sách ở nhất, đếm từ cao xuống thấp	_________")
	T5CBN := Top5CitiesByNumber(people)
	fmt.Println(T5CBN)
	for _,value := range T5CBN{
		fmt.Printf("%s : %d\n", value.Name, value.Number)
	}

	// 2.5 Trong mỗi thành phố, hãy tìm ra nghề nào được làm nhiều nhất
	fmt.Println("_________ 2.5 Trong mỗi thành phố, hãy tìm ra nghề nào được làm nhiều nhất	_________")
	topJobInEachCity := TopJobByNumberInEachCity(people)
	fmt.Println(topJobInEachCity)
	for key, value := range topJobInEachCity {
		fmt.Println("========> City", key)
		for k, v := range value {
			fmt.Printf("%s : %d", k, v)
			fmt.Println("")
		}
	}
	
	// 2.6 Ứng với một nghề, hãy tính mức lương trung bình
	fmt.Println("_________ 2.6 Ứng với một nghề, hãy tính mức lương trung bình	_________")
	ASBJ := AverageSalaryByJob(people)
	fmt.Println(ASBJ)
	for key, value := range ASBJ {
		fmt.Printf("%s : %d", key, value)
		fmt.Println("")
	}

	// 2.7 Năm thành phố có mức lương trung bình cao nhất
	fmt.Println("_________ 2.7 Năm thành phố có mức lương trung bình cao nhất _________")
	FCHTAS := FiveCitiesHasTopAverageSalary(people)
	fmt.Println(FCHTAS)
	for _, value := range FCHTAS {
		fmt.Printf("%s : %d\n", value.Name, value.Number)
	}

	// 2.8 Năm thành phố có mức lương trung bình của developer cao nhất
	fmt.Println("_________ 2.8 Năm thành phố có mức lương trung bình của developer cao nhất _________")
	FCHTSFD := FiveCitiesHasTopSalaryForDeveloper(people)
	fmt.Println(FCHTSFD)
	for _, value := range FCHTSFD {
		fmt.Printf("%s : %d", value.Name, value.Number)
		fmt.Println("")
	}

	// 2.9 Tuổi trung bình từng nghề nghiệp
	fmt.Println("_________ 2.9 Tuổi trung bình từng nghề nghiệp _________")
	AAPJ := AverageAgePerJob(people)
	fmt.Println(AAPJ)
	for key, value := range AAPJ {
		fmt.Printf("%s : %.2f", key, value)
		fmt.Println("")
	}

	// 2.10 Tuổi trung bình ở từng thành phố
	fmt.Println("_________ 2.10 Tuổi trung bình ở từng thành phố _________")
	AAPC := AverageAgePerCity(people)
	fmt.Println(AAPJ)
	for key, value := range AAPC {
		fmt.Printf("%s : %.2f", key, value)
		fmt.Println("")
	}
}
