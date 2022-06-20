package main

import (
	"employeemanagement/employee"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("I am inside Main")
	// addres := address.GetAddress()
	// dept := department.GetDepartment()
	// slry := salary.GetSalary()

	r := mux.NewRouter()

	// Add
	r.HandleFunc("/addemployeedetail", AddEmployeeHandler).Methods("POST")
	r.HandleFunc("/addemployeeaddress", AddAddressHandler).Methods("POST")
	r.HandleFunc("/addemployeedepartment", AddDepartmentHandler).Methods("POST")
	r.HandleFunc("/addemployeesalary", AddSalaryHandler).Methods("POST")

	//Update
	r.HandleFunc("/updateemployeedetail", UpdateEmployeeHandler).Methods("POST")
	r.HandleFunc("/updateemployeeaddress", UpdateEmplyeeAddressHandler).Methods("POST")
	r.HandleFunc("/updateemployeedepartment", UpdateEmployeeDepartmentHandler).Methods("POST")
	r.HandleFunc("/updateemployeesalary", UpdateEmployeeSalaryHandler).Methods("POST")

	//GetOne
	r.HandleFunc("/getemployeedetail", GetEmployeeHandler).Methods("GET")
	r.HandleFunc("/getemployeeaddress", GetEmplyeeAddressHandler).Methods("GET")
	r.HandleFunc("/getemployeedepartment", GetEmployeeDepartmentHandler).Methods("GET")
	r.HandleFunc("/getemployeesalary", GetEmplyeeSalaryHandler).Methods("GET")

	//GellAll
	r.HandleFunc("/getallemployeedetail", GetAllEmployeeHandler).Methods("GET")
	r.HandleFunc("/getallemployeeaddress", GetAllEmployeeAddressHandler).Methods("GET")
	r.HandleFunc("/getallemployeedepartment", GetAllEmployeeDepartmentHandler).Methods("GET")
	r.HandleFunc("/getallemployeesalary", GetAllEmployeeSalaryHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}

func AddEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("I am inside AddEmployeeHandler")
	emp := employee.GetEmployee()
	var empinput []employee.Employee
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&empinput); err != nil {
		fmt.Println("error in decode")
		return
	}
	defer r.Body.Close()
	emp.AddEmployee(empinput, TotalEmployee)
}

func AddAddressHandler(w http.ResponseWriter, r *http.Request) {
	reqSt := struct {
		EmpID        string
		AddressLine1 string
		AddressLine2 string
		Street_Name  string
		City_Name    string
		State_Name   string
		Pincode      string
	}{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&reqSt); err != nil {
		fmt.Println("error in decode")
		return
	}
	defer r.Body.Close()
}

func AddDepartmentHandler(w http.ResponseWriter, r *http.Request) {
	reqSt := struct {
		EmpID    string
		DeptName string
		DeptID   string
	}{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&reqSt); err != nil {
		fmt.Println("error in decode")
		return
	}
	defer r.Body.Close()
}

func AddSalaryHandler(w http.ResponseWriter, r *http.Request) {
	reqSt := struct {
		EmpID     string
		BSalary   string
		Cross_Pay string
	}{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&reqSt); err != nil {
		fmt.Println("error in decode")
		return
	}
	defer r.Body.Close()
}

// Update Handler

func UpdateEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	reqSt := struct {
		Name        string
		EmpID       string
		Age         string
		Sex         string
		Designation string
		EmailID     string
	}{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&reqSt); err != nil {
		fmt.Println("error in decode")
		return
	}
	defer r.Body.Close()
}

func UpdateEmplyeeAddressHandler(w http.ResponseWriter, r *http.Request) {
	reqSt := struct {
		EmpID        string
		AddressLine1 string
		AddressLine2 string
		Street_Name  string
		City_Name    string
		State_Name   string
		Pincode      string
	}{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&reqSt); err != nil {
		fmt.Println("error in decode")
		return
	}
	defer r.Body.Close()
}

func UpdateEmployeeDepartmentHandler(w http.ResponseWriter, r *http.Request) {
	reqSt := struct {
		EmpID    string
		DeptName string
		DeptID   string
	}{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&reqSt); err != nil {
		fmt.Println("error in decode")
		return
	}
	defer r.Body.Close()
}

func UpdateEmployeeSalaryHandler(w http.ResponseWriter, r *http.Request) {
	reqSt := struct {
		EmpID     string
		BSalary   string
		Cross_Pay string
	}{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&reqSt); err != nil {
		fmt.Println("error in decode")
		return
	}
	defer r.Body.Close()
}

// Get Handler

func GetEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	empId := vars["empid"]
	fmt.Println(empId)
}

func GetEmplyeeAddressHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	empId := vars["empid"]
	fmt.Println(empId)
}

func GetEmployeeDepartmentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	empId := vars["empid"]
	fmt.Println(empId)
}

func GetEmplyeeSalaryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	empId := vars["empid"]
	fmt.Println(empId)
	// API will be "/getemployee/{empid}"
}

// Get All Handler

func GetAllEmployeeHandler(w http.ResponseWriter, r *http.Request) {
}

func GetAllEmployeeAddressHandler(w http.ResponseWriter, r *http.Request) {
}

func GetAllEmployeeDepartmentHandler(w http.ResponseWriter, r *http.Request) {
}

func GetAllEmployeeSalaryHandler(w http.ResponseWriter, r *http.Request) {
	// API will be "/getemployee/{empid}"
}
