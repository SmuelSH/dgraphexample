package importendpoint

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type products struct {
	Uid string `json:"uid"`
	//IdProduct string `json:"Transaction.idProducto"`
}

type buyerTr struct {
	Uid string `json:"uid"`
}

type Transaction struct {
	IdTran   string     `json:"Transaction.idTran"`
	IdBuyer  string     `json:"Transaction.idBuyer"`
	Buyer    buyerTr    `json:"Transaction.Buyer"`
	Ip       string     `json:"Transaction.ip"`
	Device   string     `json:"Transaction.device"`
	Products []products `json:"Transaction.Products"`
}

type Product struct {
	Uid        string `json:"uid"`
	IdProducto string `json:"Product.idProduct"`
	Name       string `json:"Product.name"`
	Price      int    `json:"Product.price"`
}

type Buyer struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type BuyerNew struct {
	Uid     string `json:"uid"`
	IdBuyer string `json:"Buyer.idBuyer"`
	Name    string `json:"Buyer.name"`
	Age     int    `json:"Buyer.age"`
}

type Shop struct {
	Transaction []Transaction
	Product     []Product
	Buyer       []BuyerNew
}

func buyer(sUIdBuyer string) buyerTr {
	var buyers buyerTr

	buyers.Uid = "_:" + sUIdBuyer
	return (buyers)
}

func producto(sprod string) []products {
	var prod products
	var prods []products

	slimite := ","
	sdata := strings.Split(sprod, slimite)
	for _, scol := range sdata {
		prod.Uid = "_:" + scol
		//prod.IdProduct=scol
		prods = append(prods, prod)
	}
	//fmt.Println(sprod)
	return (prods)
}

func importTransac() []Transaction {
	var sid string
	var sidbuyer string
	var sip string

	var sdevice string
	var sproductoid string

	var tran Transaction
	var trans []Transaction

	file, err := os.Open("transactions.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)

	}
	defer file.Close()
	var slines []string

	scanner := bufio.NewScanner(file)
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 1024*1024)
	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		slines = append(slines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	stexto := strings.Join(slines, "\n")
	slimite := "\000"
	var iindice int = 0

	var i int = 0
	sdata := strings.Split(stexto, slimite)
	for _, scol := range sdata {
		//if i<=100{
		if iindice < 5 {
			switch iindice {
			case 0:
				sid = scol
			case 1:
				sidbuyer = scol
			case 2:
				sip = scol
			case 3:
				sdevice = scol
			case 4:
				sproductoid = strings.Trim(strings.Trim(scol, "("), ")")
			}

			if iindice == 4 {
				tran.IdTran = sid
				tran.IdBuyer = sidbuyer
				tran.Buyer = buyer(sidbuyer)
				tran.Ip = sip
				tran.Device = sdevice
				tran.Products = producto(sproductoid)
				//fmt.Println(producto(sproductoid))
				trans = append(trans, tran)

			}
			iindice = iindice + 1
		} else {
			iindice = 0
		}
		i++
		//}
	}
	//fmt.Println(trans)
	jsontran, err := json.Marshal(trans)
	if err != nil {
		fmt.Println("Error:.....")
		os.Exit(1)
	}
	//fmt.Println(string(jsontran))

	json_file, err := os.Create("transactions.json")
	if err != nil {
		fmt.Println(err)
	}
	defer json_file.Close()

	json_file.Write(jsontran)
	json_file.Close()

	return (trans)
}

func importProduct() []Product {
	var sid string
	var sname string
	var iprice int
	var prod Product
	var prods []Product
	var e error

	file, err := os.Open("products.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)

	}
	defer file.Close()
	var slines []string
	var saux string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		saux = scanner.Text()

		if strings.Index(saux, "\"") > -1 {
			saux = strings.Replace(saux, "'", ",", 1)
			saux = strings.Replace(saux, "\"", "", 1)
			ichar := strings.Count(saux, "'")
			saux = strings.Replace(saux, "'", "º", ichar-1)
			saux = strings.Replace(saux, "\"", "", 1)
			saux = strings.Replace(saux, "'", ",", 1)
			saux = strings.Replace(saux, "º", "'", ichar-1)
		} else {
			saux = strings.Replace(saux, "'", ",", -1)
		}

		slines = append(slines, saux)
		//println(saux)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	stexto := strings.Join(slines, ",")
	slimite := ","
	var iindice int = 0

	sdata := strings.Split(stexto, slimite)
	for _, scol := range sdata {
		switch iindice {
		case 0:
			sid = scol
			iindice = iindice + 1
		case 1:
			sname = scol
			iindice = iindice + 1
		case 2:
			iprice, e = strconv.Atoi(scol)
			if e != nil {
				fmt.Println(e.Error())
				os.Exit(1)
			}
			prod.Uid = "_:" + sid
			prod.IdProducto = sid
			prod.Name = sname
			prod.Price = iprice
			prods = append(prods, prod)
			iindice = 0
		}
	}

	jsontran, err := json.Marshal(prods)
	if err != nil {
		fmt.Println("Error:.....")
		os.Exit(1)
	}

	json_file, err := os.Create("products.json")
	if err != nil {
		fmt.Println(err)
	}
	defer json_file.Close()

	json_file.Write(jsontran)
	json_file.Close()

	return (prods)
}

func importBuyer() []BuyerNew {

	jsonFile, err := os.Open("buyers.json")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	var buyers []Buyer
	var buyersnew []BuyerNew
	var buyernew BuyerNew

	json.Unmarshal(byteValue, &buyers)

	for i := 0; i < len(buyers); i++ {

		buyernew.Uid = "_:" + buyers[i].Id
		buyernew.IdBuyer = buyers[i].Id
		buyernew.Name = buyers[i].Name

		buyernew.Age = buyers[i].Age
		buyersnew = append(buyersnew, buyernew)

		jsontran, err := json.Marshal(buyersnew)
		if err != nil {
			fmt.Println("Error:.....")
			os.Exit(1)
		}

		json_file, err := os.Create("buyersNew.json")
		if err != nil {
			fmt.Println(err)
		}
		defer json_file.Close()

		json_file.Write(jsontran)
		json_file.Close()
	}

	return (buyersnew)
}

func GenerateShop() Shop {
	var shop Shop

	shop.Buyer = importBuyer()
	shop.Transaction = importTransac()
	shop.Product = importProduct()

	jsontran, err := json.Marshal(shop)
	if err != nil {
		fmt.Println("Error:.....")
		os.Exit(1)
	}

	json_file, err := os.Create("shop.json")
	if err != nil {
		fmt.Println(err)
	}

	defer json_file.Close()

	json_file.Write(jsontran)
	json_file.Close()

	return shop
}
