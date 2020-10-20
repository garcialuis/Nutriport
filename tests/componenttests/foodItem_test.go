package tests

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/garcialuis/Nutriport/api/controllers"
	"github.com/garcialuis/Nutriport/api/seed"
	"github.com/garcialuis/Nutriport/client/client"
	nutriportclient "github.com/garcialuis/Nutriport/sdk/client"
	"github.com/garcialuis/Nutriport/sdk/models"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"gopkg.in/go-playground/assert.v1"
)

var server = controllers.Server{}

func TestMain(m *testing.M) {
	Database()
	StartServer()
	os.Exit(m.Run())
}

func Database() {
	var err error

	TestDbDriver := os.Getenv("DB_POSTGRES_DRIVER")

	if TestDbDriver == "postgres" {
		DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", os.Getenv("TEST_DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("TEST_DB_USER"), os.Getenv("NUTRIPORT_TEST_DB"), os.Getenv("TEST_DB_PASSWORD"))
		server.DB, err = gorm.Open(TestDbDriver, DBURL)

		if err != nil {
			fmt.Printf("Cannot connect to %s database\n", TestDbDriver)
			log.Fatal("This is the error: ", err)
		} else {
			fmt.Printf("We are connected to the %s database\n", TestDbDriver)
		}
	}

	fmt.Println("Seeding test db:")
	seed.Load(server.DB)
}

func StartServer() {
	server.Router = mux.NewRouter()
	server.InitializeRoutes()
	fmt.Println("Starting up server:")

	serviceRunning := make(chan struct{})
	serviceDone := make(chan struct{})

	go func() {
		close(serviceRunning)
		server.Run(":8085")
		defer close(serviceDone)
	}()
}

func TestNutriportClient(t *testing.T) {

	time.Sleep(2 * time.Second)

	cfg := client.DefaultTransportConfig().WithHost("localhost:8085")
	c := client.NewHTTPClientWithConfig(nil, cfg)

	homeOk, err := c.Home.Home(nil)

	if err != nil {
		log.Fatalf("Failed test due to: %v", err)
	}

	fmt.Printf("homeOk: %v\n", homeOk)
}

func TestCreateFoodItem(t *testing.T) {

	time.Sleep(2 * time.Second)

	nutriportClient := nutriportclient.NewClient()

	itemName := "Cucumber"
	var cupQtty float32 = 1
	var gWt float32 = 141.74
	var oWt float32 = 5

	foodItemToCreate := models.FoodItem{
		Name:          itemName,
		CarbLevelID:   2,
		FoodVarietyID: 1,
		FoodGroupID:   2,
		CupQuantity:   cupQtty,
		GramWeight:    gWt,
		OnceWeight:    oWt,
	}

	newFoodItem := nutriportClient.CreateFoodItem(foodItemToCreate)

	fmt.Println("NEW FOOD ITEM CREATED USING CLIENT: ")
	fmt.Println(newFoodItem)

}

func TestGetFoodItems(t *testing.T) {

	// TODO: Seed test database with expected records

	cfg := client.DefaultTransportConfig().WithHost("localhost:8085")
	c := client.NewHTTPClientWithConfig(nil, cfg)

	foodItems, err := c.FoodItem.GetAllFoodItems(nil)

	if err != nil {
		log.Fatalf("Failed to retrieve food items due to: %v", err)
	}

	foodItemsPayload := foodItems.GetPayload()
	fmt.Printf("Name: %#v, FoodGroup: %v\n", *foodItemsPayload[0].Name, foodItemsPayload[0].FoodGroup.FoodGroupName)
}

func TestGetAllFoodItems(t *testing.T) {

	nutriportClient := nutriportclient.NewClient()

	foodItems := nutriportClient.GetAllFoodItems()

	for _, foodItem := range foodItems {
		fmt.Println(foodItem)
	}
}

func TestDeleteFoodItem(t *testing.T) {

	nutriportClient := nutriportclient.NewClient()

	foodItemName := "Cucumber"
	affectedRecords := nutriportClient.DeleteFoodItem(foodItemName)
	fmt.Println("Affected Records: ", affectedRecords)

	assert.Equal(t, 1, affectedRecords)
}
