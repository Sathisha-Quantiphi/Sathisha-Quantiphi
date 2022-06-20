package salary

import (
	"context"
	"sync"
	"time"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
)

// salaryInter: Private Interface
type salaryInter interface {
	AddEmployeeSalary()
	UpdateEmployeeSalary()
	GetEmployeeSalary()
	GetAllEmployeesSalary()
}

func GetSalary() *Salary {
	return &Salary{}
}

// Interface method implementation
func (d *Salary) AddEmployeeSalary(salry []Salary) {
	var w sync.WaitGroup
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return
	}
	collection := client.Database("employee").Collection("empsalary")

	if len(salry) > 0 {
		for i := 0; i < len(salry); i++ {
			w.Add(1)
			go addingEmpSalary(&salry[i], &w, collection, ctx)
		}
		w.Wait()
	}
}

func addingEmpSalary(slry *Salary, wg *sync.WaitGroup, coll *mongo.Collection, ctx context.Context) {
	_, err := coll.InsertOne(ctx, slry)
	if err != nil {
		return
	}
	wg.Done()
}

func (d *Salary) UpdateEmployeeSalary(salry Salary) {
	var s Salary
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return
	}
	coll := client.Database("employee").Collection("empsalary")
	update := bson.D{
		{"$set", bson.D{
			{"bsalary", salry.BSalary},
			{"crosspay", salry.Cross_Pay},
		}}}
	err = coll.FindOneAndUpdate(ctx, bson.M{"empId": salry.EmpId}, update).Decode(&s)
	if err != nil {
		return
	}
	return
}

func (d *Salary) GetEmployeeSalary(empId string) Salary {
	var empslry Salary
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return Salary{}
	}
	collection := client.Database("employee").Collection("empsalary")
	data := collection.FindOne(ctx, bson.M{"empId": empId})
	data.Decode(&empslry)
	return empslry
}

func (d *Salary) GetAllEmployeesSalary() []Salary {
	var empslry []Salary
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return []Salary{}
	}
	collection := client.Database("employee").Collection("empsalary")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return []Salary{}
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var s Salary
		cursor.Decode(&s)
		empslry = append(empslry, s)
	}
	return empslry
}
