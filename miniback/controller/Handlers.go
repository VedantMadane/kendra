package controller

import (
	"gominiback/backend/connections"
	"gominiback/backend/data"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetStudentData(c *gin.Context) {

	client, ctx := connections.ConnectDB()
	database := client.Database("SchoolData")
	collection := database.Collection("student_Registration")
	var data data.RegistrationData
	c.Bind(&data)
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
	// ConnectDB.Collection.InsertOne(ctx, data)
	collection.InsertOne(ctx, data)
	log.Println("Success!")

}

func MySelectedData(c *gin.Context) {
	client, ctx := connections.ConnectDB()
	defer client.Disconnect(ctx)
	database := client.Database("SchoolData")
	collection := database.Collection("student_Registration")

	findCursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var data []bson.M
	if err = findCursor.All(ctx, &data); err != nil {
		log.Fatal(err)
	}
	c.Bind(&data)
	c.JSON(http.StatusOK, gin.H{
		"got": data,
	})

	// defer findCursor.Close(ctx)
	// for findCursor.Next(ctx) {
	// 	var data []bson.M
	// 	if err := findCursor.Decode(&data); err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	c.Bind(&data)
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"got": data,
	// 	})
	// }
}

func UpdateStudent(c *gin.Context) {
	client, ctx := connections.ConnectDB()
	defer client.Disconnect(ctx)
	database := client.Database("SchoolData")
	collection := database.Collection("student_Registration")

	update, err := collection.UpdateOne(ctx, bson.M{"mothername": "Shila"},
		bson.M{"$set": bson.M{
			"mothername": "Lila",
		},
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"update": update,
	})
}

func DeleteStudent(c *gin.Context) {
	client, ctx := connections.ConnectDB()
	defer client.Disconnect(ctx)
	database := client.Database("SchoolData")
	collection := database.Collection("student_Registration")

	delete, err := collection.DeleteOne(ctx, bson.M{"mothername": "Lila"})
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"update": delete,
	})
}
