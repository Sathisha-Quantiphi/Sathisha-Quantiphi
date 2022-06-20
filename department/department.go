package department

import (
	"context"
	"sync"
	"time"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
)

// departmentInter: Private Interface
type departmentInter interface {
	AddEmployeeDepartment()
	UpdateEmployeeDepartment()
	GetEmployeeDepartment()
	GetAllEmployeeDepartments()
}

func GetDepartment() *Department {
	return &Department{}
}

// Interface method implementation
func (d *Department) AddEmployeeDepartment(dept []Department) {
	var w sync.WaitGroup
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return
	}
	collection := client.Database("employee").Collection("empdepartment")

	if len(dept) > 0 {
		for i := 0; i < len(dept); i++ {
			w.Add(1)
			go addingEmpDepartment(&dept[i], &w, collection, ctx)
		}
		w.Wait()
	}
}

func addingEmpDepartment(dept *Department, wg *sync.WaitGroup, coll *mongo.Collection, ctx context.Context) {
	_, err := coll.InsertOne(ctx, dept)
	if err != nil {
		return
	}
	wg.Done()
}

func (d *Department) UpdateEmployeeDepartment(dept Department) {
	var dep Department
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return
	}
	coll := client.Database("employee").Collection("empdepartment")
	update := bson.D{
		{"$set", bson.D{
			{"depatname", dept.DeptName},
			{"deptid", dept.DeptID},
		}}}
	err = coll.FindOneAndUpdate(ctx, bson.M{"empId": dept.EmpID}, update).Decode(&dep)
	if err != nil {
		return
	}
	return
}

func (d *Department) GetAllEmployeeDepartments() []Department {
	var empdept []Department
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return []Department{}
	}
	collection := client.Database("employee").Collection("empdepartment")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return []Department{}
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var d Department
		cursor.Decode(&d)
		empdept = append(empdept, d)
	}
	return empdept
}

func (d *Department) GetEmployeeDepartment(empId string) Department {
	var empdept Department
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return Department{}
	}
	collection := client.Database("employee").Collection("empdepartment")
	data := collection.FindOne(ctx, bson.M{"empId": empId})
	data.Decode(&empdept)
	return empdept
}
