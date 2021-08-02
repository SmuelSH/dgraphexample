package main

import (
	"encoding/json"
	"fmt"
	"bufio"
	"log"
	"os"
	"strings"
)

type product struct{
	IdProduct string `json:"idProduct"`
 }
 
 type transaction struct{
	  Id          string  `json:"id"`
	  Buyerid     string `json:"buyerid"`
	  Ip          string `json:"ip"`
	  Device      string `json:"device"`
	  Products    []product `json:"products"`
 }
 

func producto(sprod string)[]product {
var prod    product
var prods []product

slimite := ","
sdata := strings.Split(sprod, slimite)
	for _, scol := range sdata {		
         prod.IdProduct=scol
         prods=append(prods,prod)		 
	}
    //fmt.Println(sprod) 
	return(prods)
}

func main() {
	
var sid string
var sbuyerid string
var sip string

var sdevice string 
var sproductoid string

var tran transaction
var trans []transaction

file, err:= os.Open("example.txt")
if err !=nil {
	log.Fatalf( "readLines: %s", err)

}
defer file.Close()
var  slines []string

scanner:= bufio.NewScanner(file) 
buf := make([]byte, 0, 64*1024)
scanner.Buffer(buf, 1024*1024)
for scanner.Scan() { 
	//fmt.Println(scanner.Text())
	slines = append(slines, scanner.Text()) 
}

if err := scanner.Err(); err != nil {
	fmt.Fprintln(os.Stderr, "reading standard input:", err)
}


stexto:=strings.Join(slines,"\n")
slimite :="\000" 
var iindice int=0


sdata := strings.Split(stexto, slimite)
	for _, scol := range sdata {				   
           if iindice<5  {            
				switch iindice {			
				case 0:
					sid=scol
				case 1:	
					sbuyerid =scol
				case 2:
					sip=scol
				case 3:
					sdevice=scol
				case 4:
				     sproductoid=strings.Trim(strings.Trim(scol,"("),")")
				}
	        
				if iindice==4 {				
				tran.Id=sid
				tran.Buyerid=sbuyerid
				tran.Ip=sip
				tran.Device= sdevice
				tran.Products=producto(sproductoid)				
				/*				
				if err  != nil {
					fmt.Fprintln(os.Stderr, "asign:", err)
				}*/

                //fmt.Println(sproductoid)	
				fmt.Println(producto(sproductoid)) 
				trans =append(trans, tran)			  

				}
			  iindice=iindice+1 
		   }else
		   {
			iindice=0            
		   }		 
	}
	fmt.Println(trans)
	jsontran,err:=json.Marshal(trans)
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

}

