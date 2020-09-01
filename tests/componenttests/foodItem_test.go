package tests

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/garcialuis/Nutriport/api/controllers"
	"github.com/garcialuis/Nutriport/client/client"
	"github.com/garcialuis/Nutriport/sdk/client/fooditem"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

var server = controllers.Server{}

func TestMain(m *testing.M) {
	fmt.Println("Testing main T test first")
	Database()
	StartServer()
	os.Exit(m.Run())
}

func Database() {
	var err error

	TestDbDriver := os.Getenv("DB_POSTGRES_DRIVER")

	if TestDbDriver == "postgres" {
		DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", os.Getenv("TEST_DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("TEST_DB_USER"), "nutriport", os.Getenv("TEST_DB_PASSWORD"))
		server.DB, err = gorm.Open(TestDbDriver, DBURL)

		if err != nil {
			fmt.Printf("Cannot connect to %s database\n", TestDbDriver)
			log.Fatal("This is the error: ", err)
		} else {
			fmt.Printf("We are connected to the %s database\n", TestDbDriver)
		}
	}
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

	cfg := client.DefaultTransportConfig().WithHost("localhost:8085")
	c := client.NewHTTPClientWithConfig(nil, cfg)

	homeOk, err := c.Home.Home(nil)

	if err != nil {
		log.Fatalf("Failed test due to: %v", err)
	}

	fmt.Printf("homeOk: %v\n", homeOk)
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

	foodItemClient := fooditem.NewClientService()

	foodItems := foodItemClient.GetAllFoodItems()

	for _, foodItem := range foodItems {
		fmt.Println(foodItem)
	}
}
