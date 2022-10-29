package main

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"keys"
	"log"
)

type Boss struct {
	Name string `bson:"Name"`
	Des  string `bson:"Description"`
}

//Two custom structure variables for insertion and replacement
var bee Boss = Boss{"bee", "medium"}
var mantis Boss = Boss{"Mantis Lord", "medium"}

func checkerr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func main() {
	//build connect
	uri := keys.Mongouri
	ctx := context.TODO()
	dbname := "hollow"
	collname := "boss"
	if len(uri) == 0 {
		logrus.Fatal("uri settings corrupt!")
	}
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		logrus.Error("Build connect failed")
	}
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			logrus.Error("disconnet error:", err)
		}
	}()
	if err = client.Ping(ctx, nil); err != nil {
		logrus.Error("Connect failed", err)
	} else {
		fmt.Println("Pong!")
	}
	//绑定集合
	coll := client.Database(dbname).Collection(collname)
	/*查看数据库*/
	{
		database, err := client.ListDatabaseNames(ctx, bson.M{})
		if err != nil {
			logrus.Info("listDataBase failed", err)
		}
		fmt.Println(database)
	}

	/****插入数据****/
	/*插入单个数据*/
	{
		Insertone, err := coll.InsertOne(ctx, bee)
		if err != nil {
			logrus.Error("Insert data failed", err)
		}
		fmt.Println("Insert a single document,here's its ID:", Insertone.InsertedID)
	}
	/*插入多个数据*/
	//{
	//	bees := []interface{}{bee, bee, bee, bee, bee, bee}
	//	insertMany, err := coll.InsertMany(ctx, bees)
	//	if err != nil {
	//		logrus.Error("Insert data failed", err)
	//	}
	//	fmt.Println("InsertMany Information:", insertMany.InsertedIDs)
	//}
	//删除操作

	/****删除数据****/
	/*删除单个数据*/
	//{
	//	filter := bson.D{{"Name", "bee"}}
	//	result, err := coll.DeleteOne(ctx, filter)
	//	if err != nil {
	//		logrus.Error("Delete data failed", err)
	//	}
	//	fmt.Println("Delete data successfully", result)
	//}
	/*删除多个数据*/
	//{
	//	filter := bson.D{{"Name", "Bee"}}
	//	result, err := coll.DeleteMany(ctx, filter)
	//	if err != nil {
	//		logrus.Error("Delete data failed", err)
	//	}
	//	fmt.Printf("Delete %v data successfully", result.DeletedCount)
	//}

	/****更新数据****/
	/*更新单个数据*/
	//{
	//
	//	filter := bson.D{{"Name", "bee"}}
	//	opts := options.Update().SetUpsert(true)
	//	update := bson.D{
	//		{"$set",
	//			bson.D{{"Name", "Mantis Lord"}},
	//		},
	//	}
	//	result, err := coll.UpdateOne(ctx, filter, update, opts)
	//	if err != nil {
	//		logrus.Error("Update data failed", err)
	//	}
	//	if result.MatchedCount != 0 {
	//		fmt.Println("Data update successfully!\n", result.MatchedCount, result.ModifiedCount)
	//	}
	//	if result.UpsertedCount != 0 {
	//		fmt.Println("Insert a new date successfully!\n", result.UpsertedCount)
	//	}
	//}
	/*更新多个数据*/
	//{
	//	filter := bson.D{{"Name", "bee"}}
	//	opts := options.Update().SetUpsert(true)
	//	update := bson.D{
	//		{"$set",
	//			bson.D{{"Name", "Bee"}},
	//		},
	//	}
	//	result, err := coll.UpdateMany(ctx, filter, update, opts)
	//	if err != nil {
	//		logrus.Error("Update data failed", err)
	//	}
	//	if result.MatchedCount != 0 {
	//		fmt.Println("Data update successfully!\n", result.MatchedCount, result.ModifiedCount)
	//	}
	//	if result.UpsertedCount != 0 {
	//		fmt.Println("Insert a new date successfully!\n", result.UpsertedCount)
	//	}
	//}

	/****查询数据****/
	/*查询单个数据*/
	//{
	//
	//	var result Boss
	//	filter := bson.D{{"Name", "Bee"}}
	//	err = coll.FindOne(ctx, filter).Decode(&result)
	//	if err == mongo.ErrNoDocuments {
	//		fmt.Printf("No document was found with the filter %s", filter)
	//		return
	//	}
	//	if err != nil {
	//		logrus.Info("find data failed")
	//	}
	//	jsonData, err := json.MarshalIndent(result, "", "    ")
	//	logrus.Error("Marshal data failed", err)
	//	fmt.Printf("%s\n", jsonData)
	//
	//}
	/*查询多个数据*/
	//{
	//	var results []Boss
	//	filter := bson.D{{"Name", "Bee"}}
	//	findoptions := options.Find()
	//	findoptions.SetLimit(10)
	//	cur, err := coll.Find(ctx, filter, findoptions)
	//	if err != nil {
	//		logrus.Info("find data failed")
	//	}
	//	defer func() {
	//		err := cur.Close(ctx)
	//		checkerr(err)
	//	}()
	//	K := 0
	//	for cur.Next(ctx) {
	//		var result Boss
	//		err := cur.Decode(&result)
	//		if err != nil {
	//			logrus.Error("Decode data failed")
	//		}
	//		K++
	//		results = append(results, result)
	//	}
	//	fmt.Println(K)
	//	fmt.Println(results)
	//}
}
