package address

type Address struct {
	EmpID        string `json:"empId"`
	AddressLine1 string `json:"addressline1"`
	AddressLine2 string `json:"addressline1"`
	Street_Name  string `json:"streetname"`
	City_Name    string `json:"cityname"`
	State_Name   string `json:"statename"`
	Pincode      string `json:"pincode"`
}
