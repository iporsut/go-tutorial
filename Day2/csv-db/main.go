package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Employee struct {
	ID        int    `json:"employee_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

func main() {
	http.HandleFunc("/employees", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			ListEmployeeHandler(w, r)
		case "POST":
			CreateEmployeeHandler(w, r)
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	})

	http.HandleFunc("/employees/", func(w http.ResponseWriter, r *http.Request) {
		idStr := strings.TrimPrefix(r.URL.Path, "/employees/")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		switch r.Method {
		case "GET":
			FindEmployeeHandler(id, w, r)
		case "PUT", "POST":
			UpdateEmployeeHandler(id, w, r)
		case "DELETE":
			DeleteEmployeeHandler(id, w, r)
		}
	})

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func ListEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	empls, err := AllEmployee()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	body, err := json.Marshal(empls)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "%s", body)
}

func CreateEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	var emp Employee
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	err = json.Unmarshal(b, &emp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = CreateEmployee(emp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func FindEmployeeHandler(id int, w http.ResponseWriter, r *http.Request) {
	emp, err := FindEmployeeByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	body, err := json.Marshal(emp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "%s", body)
}

func UpdateEmployeeHandler(id int, w http.ResponseWriter, r *http.Request) {
	var emp Employee
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	err = json.Unmarshal(b, &emp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = UpdateEmployee(id, emp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func DeleteEmployeeHandler(id int, w http.ResponseWriter, r *http.Request) {
	_, err := FindEmployeeByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = DeleteEmployee(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func AllEmployee() ([]Employee, error) {
	var empls []Employee

	f, err := os.Open("employee.csv")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.Read()

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}

		age, err := strconv.Atoi(record[0])
		if err != nil {
			return nil, err
		}

		empls = append(empls, Employee{
			ID:        age,
			FirstName: strings.TrimSpace(record[1]),
			LastName:  strings.TrimSpace(record[2]),
			Email:     strings.TrimSpace(record[3]),
		})
	}

	return empls, nil
}

func FindEmployeeByID(id int) (*Employee, error) {
	emps, err := AllEmployee()
	if err != nil {
		return nil, err
	}

	for _, emp := range emps {
		if emp.ID == id {
			return &emp, nil
		}
	}

	return nil, fmt.Errorf("Not Found")
}

func CreateEmployee(emp Employee) error {
	f, err := os.OpenFile("employee.csv", os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	w := csv.NewWriter(f)

	err = w.Write([]string{
		strconv.Itoa(emp.ID),
		emp.FirstName,
		emp.LastName,
		emp.Email,
	})
	if err != nil {
		return err
	}

	return nil
}

func UpdateEmployee(id int, emp Employee) error {
	emps, err := AllEmployee()
	if err != nil {
		return err
	}

	idx := -1
	for i := range emps {
		if emps[i].ID == id {
			idx = i
			break
		}
	}
	if idx == -1 {
		return fmt.Errorf("Not Found")
	}

	emps[idx].FirstName = emp.FirstName
	emps[idx].LastName = emp.LastName
	emps[idx].Email = emp.Email
	err = os.Remove("employee.csv")
	if err != nil {
		return err
	}
	f, err := os.Create("employee.csv")
	if err != nil {
		return err
	}
	defer f.Close()

	w := csv.NewWriter(f)

	// Header
	err = w.Write([]string{"employee_id", "first_name", "last_name", "email"})
	if err != nil {
		return err
	}

	for _, emp := range emps {
		err = w.Write([]string{
			strconv.Itoa(emp.ID),
			emp.FirstName,
			emp.LastName,
			emp.Email,
		})
		fmt.Println(emp)
		if err != nil {
			return err
		}
	}

	w.Flush()
	if err := w.Error(); err != nil {
		return err
	}

	return nil
}

func DeleteEmployee(id int) error {
	emps, err := AllEmployee()
	if err != nil {
		return err
	}

	idx := -1
	for i := range emps {
		if emps[i].ID == id {
			idx = i
			break
		}
	}
	if idx == -1 {
		return fmt.Errorf("Not Found")
	}

	emps = append(emps[:idx], emps[idx+1:]...)

	err = os.Remove("employee.csv")
	if err != nil {
		return err
	}
	f, err := os.Create("employee.csv")
	if err != nil {
		return err
	}
	defer f.Close()

	w := csv.NewWriter(f)

	// Header
	w.Write([]string{"employee_id", "first_name", "last_name", "email"})
	for _, emp := range emps {
		err = w.Write([]string{
			strconv.Itoa(emp.ID),
			emp.FirstName,
			emp.LastName,
			emp.Email,
		})
		if err != nil {
			return err
		}
	}

	w.Flush()
	if err := w.Error(); err != nil {
		return err
	}

	return nil
}
