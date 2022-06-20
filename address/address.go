package address

import (
	"context"
	"sync"
	"time"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
)

// addressInter : Private Interface
type addressInter interface {
	AddEmployeeAddress()
	UpdateEmployeeAddress()
	GetEmployeeAddress()
	GetAllEmployeesAddress()
}

func GetAddress() *Address {
	return &Address{}
}

// Interface method implementation
func (a *Address) AddEmployeeAddress(add []Address) {
	var w sync.WaitGroup
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return
	}
	collection := client.Database("employee").Collection("empaddress")

	if len(add) > 0 {
		for i := 0; i < len(add); i++ {
			w.Add(1)
			go addingEmpAddress(&add[i], &w, collection, ctx)
		}
		w.Wait()
	}
}

func addingEmpAddress(emp *Address, wg *sync.WaitGroup, coll *mongo.Collection, ctx context.Context) {
	_, err := coll.InsertOne(ctx, emp)
	if err != nil {
		return
	}
	wg.Done()
}

func (a *Address) UpdateEmployeeAddress(add Address) {
	var addr Address
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return
	}
	coll := client.Database("employee").Collection("empaddress")
	update := bson.D{
		{"$set", bson.D{
			{"addressline1", add.AddressLine1},
			{"addressLine2", add.AddressLine2},
			{"streetname", add.Street_Name},
			{"cityname", add.City_Name},
			{"statename", add.State_Name},
			{"pincode", add.Pincode},
		}}}
	err = coll.FindOneAndUpdate(ctx, bson.M{"empId": add.EmpID}, update).Decode(&addr)
	if err != nil {
		return
	}
	return
}

func (a *Address) GetEmployeeAddress(empId string) Address {
	var empadd Address
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return Address{}
	}
	collection := client.Database("employee").Collection("empaddress")
	data := collection.FindOne(ctx, bson.M{"empId": empId})
	data.Decode(&empadd)
	return empadd
}

func (a *Address) GetAllEmployeesAddress() []Address {
	var empadd []Address
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return []Address{}
	}
	collection := client.Database("employee").Collection("empaddress")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return []Address{}
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var a Address
		cursor.Decode(&a)
		empadd = append(empadd, a)
	}
	return empadd
}
