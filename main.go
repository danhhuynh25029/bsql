package main

import(
	"fmt"
	"bufio"
	"os"
	s "strings"
	r "resources/code"
)
var DB string = "";
func main(){
	var sql string = " ";
	for true{
		if len(sql) != 0{
			fmt.Print("tsql> ");
		}
		reader := bufio.NewReader(os.Stdin);
		sql, _ = reader.ReadString('\n');
		if sql != "\n"{
			sql = s.Trim(sql,"\n;");
		}
		if sql == "exit"{
			break;
		}
		query(&sql);
	}	
	
}
func query(sql *string){
	// fmt.Println(*sql);
	var arr []string = s.Split(*sql," ");
	arr[0] = s.ToUpper(arr[0]);
	switch arr[0]{
		case "USE":
			DB = r.Use(arr);
			fmt.Println(DB);
		case "CREATE":
			r.Create(arr,DB);
		case "SELECT":
			r.Select(arr,DB);
		case "DELETE":
			r.Delete(arr,DB);
		case "INSERT":
			r.Insert(arr,DB);
		case "UPDATE":
			r.Update(arr,DB);
		default:
			fmt.Println("query is not exists");
	}
}
