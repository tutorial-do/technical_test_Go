package handlers

import (
	// "context"
	"encoding/csv"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"technical_test_Go/backend/models"
	"time"
	// "github.com/dgraph-io/dgo"
)

const (
	// URL is the given endpoint that will be consulted
	URL       = "https://kqxty15mpg.execute-api.us-east-1.amazonaws.com/"
	layoutISO = "2006-01-02"
)

func timeConversor(date string) string {
	if date == "" {
		currentDate := time.Now().Unix()
		currentDateString := strconv.FormatInt(currentDate, 10)
		return currentDateString
	}
	t, _ := time.Parse(layoutISO, date)
	dateUnix := t.Unix()

	dateUnixString := strconv.FormatInt(dateUnix, 10)

	return dateUnixString
}

// FetchData is a function to that gets and retrieves all data from the endpoints
func FetchData(date string) (buyersList []models.Buyer, productsList []models.Product, transactionsList []models.Transaction) {

	dateTime := timeConversor(date)

	buyers, err := fetchBuyersData(dateTime)
	if err != nil {
		log.Fatalln(err)
	}

	products, err := fetchProductsData(dateTime)
	if err != nil {
		log.Fatalln(err)
	}

	transactions, err := fetchTransactionsData(dateTime)
	if err != nil {
		log.Fatalln(err)
	}

	return buyers, products, transactions
}

func fetchBuyersData(dateTime string) ([]models.Buyer, error) {
	response, err := http.Get(URL + "buyers" + "?date=" + dateTime)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	type buyersList []models.Buyer

	var allBuyers buyersList
	var buyers []models.Buyer

	err = json.NewDecoder(response.Body).Decode(&allBuyers)
	if err != nil {
		log.Fatalln(err)
	}

	for _, element := range allBuyers {
		elementID := element.ID
		elementName := element.Name
		elementAge := element.Age

		buyer, err := models.NewBuyer(elementID, elementName, elementAge)

		if err != nil {
			log.Fatal(err)
		}
		buyers = append(buyers, *buyer)
	}

	return buyers, nil
}

func fetchProductsData(dateTime string) ([]models.Product, error) {

	response, err := http.Get(URL + "products" + "?date=" + dateTime)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	r := csv.NewReader(response.Body)
	r.Comma = '\''
	records, _ := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var allProducts []models.Product

	for _, element := range records {
		elementID := string(element[0])
		elementName := string(element[1])
		elementPrice, _ := strconv.Atoi(element[2])

		product, err := models.NewProduct(elementID, elementName, elementPrice)

		if err != nil {
			log.Fatal(err)
		}
		allProducts = append(allProducts, *product)
	}

	return allProducts, nil
}

func fetchTransactionsData(dateTime string) ([]models.Transaction, error) {

	response, err := http.Get(URL + "transactions" + "?date=" + dateTime)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var allTransactions []models.Transaction

	records := strings.Split(string(b), "#")

	for index, element := range records {

		if index == 0 {
			continue
		}

		subelements := strings.Split(element, "\x00")

		id := subelements[0]
		buyerID := subelements[1]
		ip := subelements[2]
		device := subelements[3]
		productsIds := strings.Split(subelements[4][1:len(subelements[4])-1], ",")

		transaction, err := models.NewTransaction(id, buyerID, ip, device, productsIds)
		if err != nil {
			log.Fatal(err)
		}

		allTransactions = append(allTransactions, *transaction)
	}

	return allTransactions, nil
}

// Strainer to check for repeated data already in the database
// func Strainer(dgraphClient *dgo.Dgraph, dataBuyers []models.Buyer, dataProducts []models.Product, dataTransactions []models.Transaction) {
// 	UniqueBuyers, err := strainerBuyers()
// 	UniqueProducts, err := strainerProducts()
// 	UniqueTransactions, err := strainerTransactions()
// }

// func strainerBuyers(dgraphClient *dgo.Dgraph, dataBuyers []models.Buyer) ([]models.Buyer, error) {
// 	ctx := context.Background()
// 	q := `{
// 			allBuyers(func: type(Buyer)){
// 				id
// 				name
// 				age
// 			}
// 		}`

// 	txn := dgraphClient.NewTxn()

// 	defer txn.Discard(ctx)

// 	var

// 	res, err := txn.Query(ctx, q)
// 	if err != nil {
// 		return ,
// 	}
// 	res.allBuyers

// }
// func strainerProducts(dgraphClient *dgo.Dgraph, dataProducts []models.Product) ([]models.Product, error) {
// }
// func strainerTransactions(dgraphClient *dgo.Dgraph, dataTransactions []models.Transaction) ([]models.Transaction, error) {
// }
