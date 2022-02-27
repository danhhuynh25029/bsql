package main

import(
	"fmt"
	"bufio"
	// "io/ioutil"
	"os"
	s "strings"
)
func main(){
	var sql string = " ";
	for true{
		if sql == "exit\n"{
			break;
		}
		if len(sql) != 0{
			fmt.Print("tsql> ");
		}
		reader := bufio.NewReader(os.Stdin);
		sql, _ = reader.ReadString('\n');
		sql = s.Trim(sql,"\n")
		query(&sql);
	}	
	
}
func query(sql *string){
	switch sql{
		case 
	}
}
func use(){

}
func create(){

}
func select(){

}

func delect(){

}

func insert(){

}