package handlers

import (
	// "context"
	"context"
	"encoding/csv"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"technical_test_Go/backend/models"
	"time"

	"github.com/dgraph-io/dgo"
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

// DataFilter to check for repeated data already in the database
func DataFilter(dgraphClient *dgo.Dgraph, dataBuyers []models.Buyer, dataProducts []models.Product, dataTransactions []models.Transaction) ([]models.Buyer, []models.Product, []models.Transaction, error) {
	uniqueBuyers, err := uniqueBuyersSelector(dgraphClient, dataBuyers)
	if err != nil {
		return nil, nil, nil, err
	}
	uniqueProducts, err := uniqueProductsSelector(dgraphClient, dataProducts)
	if err != nil {
		return nil, nil, nil, err
	}
	uniqueTransactions, err := uniqueTransactionsSelector(dgraphClient, dataTransactions)
	if err != nil {
		return nil, nil, nil, err
	}

	return uniqueBuyers, uniqueProducts, uniqueTransactions, nil
}

func uniqueBuyersSelector(dgraphClient *dgo.Dgraph, dataBuyersEndpoint []models.Buyer) ([]models.Buyer, error) {
	ctx := context.Background()
	q := `{
			allBuyers(func: type(Buyer)){
				id
				name
				age
			}
		}`

	txn := dgraphClient.NewTxn()

	defer txn.Discard(ctx)

	responseBuyersDB, err := txn.Query(ctx, q)
	if err != nil {
		return nil, err
	}

	type buyerResponse struct {
		AllBuyers []models.Buyer `json:"allBuyers,omitempty"`
	}

	var dataBuyersResponse *buyerResponse

	err = json.Unmarshal([]byte(responseBuyersDB.Json), &dataBuyersResponse)
	if err != nil {
		return nil, err
	}

	dataBuyersDB := dataBuyersResponse.AllBuyers

	var finalBuyers []models.Buyer

	for _, newBuyer := range dataBuyersEndpoint {
		flag := false
		for _, dbBuyer := range dataBuyersDB {
			if newBuyer.ID == dbBuyer.ID {
				flag = true
				break
			}
		}
		if !flag {
			finalBuyers = append(finalBuyers, newBuyer)
		}
	}

	return finalBuyers, nil
}

func uniqueProductsSelector(dgraphClient *dgo.Dgraph, dataProductsEndpoint []models.Product) ([]models.Product, error) {
	ctx := context.Background()
	q := `{
			allProducts(func: type(Product)){
				id
				name
				price
			}
		}`

	txn := dgraphClient.NewTxn()

	defer txn.Discard(ctx)

	res, err := txn.Query(ctx, q)
	if err != nil {
		return nil, err
	}

	type productResponse struct {
		AllProducts []models.Product `json:"allProducts,omitempty"`
	}

	var dataProductsResponse *productResponse

	err = json.Unmarshal(res.Json, &dataProductsResponse)
	if err != nil {
		return nil, err
	}

	dataProductsDB := dataProductsResponse.AllProducts

	var finalProducts []models.Product

	for _, newProduct := range dataProductsEndpoint {
		flag := false
		for _, dbProduct := range dataProductsDB {
			if newProduct.ID == dbProduct.ID {
				flag = true
				break
			}
		}
		if !flag {
			finalProducts = append(finalProducts, newProduct)
		}
	}

	return finalProducts, nil
}

func uniqueTransactionsSelector(dgraphClient *dgo.Dgraph, dataTransactionsEndpoint []models.Transaction) ([]models.Transaction, error) {
	ctx := context.Background()
	q := `{
			allTransactions(func: type(Transaction)){
				id
			}
		}`

	txn := dgraphClient.NewTxn()

	defer txn.Discard(ctx)

	res, err := txn.Query(ctx, q)
	if err != nil {
		return nil, err
	}

	type TransactionResponse struct {
		AllTransactions []models.Transaction `json:"allTransactions,omitempty"`
	}

	var dataTransactionsResponse *TransactionResponse

	err = json.Unmarshal(res.Json, &dataTransactionsResponse)
	if err != nil {
		return nil, err
	}

	dataTransactionsDB := dataTransactionsResponse.AllTransactions

	var finalTransactions []models.Transaction

	for _, newTransaction := range dataTransactionsEndpoint {
		flag := false
		for _, dbTransaction := range dataTransactionsDB {
			if newTransaction.ID == dbTransaction.ID {
				flag = true
				break
			}
		}
		if !flag {
			finalTransactions = append(finalTransactions, newTransaction)
		}
	}

	return finalTransactions, nil
}
