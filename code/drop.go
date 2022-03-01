package code

import(
	"fmt"
	s "strings"
	"os"
)
func Drop(arr []string,db string){
	arr[1] = s.ToUpper(arr[1]);
	switch arr[1]{
	case "TABLE":
		if len(db) == 0{
			fmt.Println("query error");
		}else{
			dropTable(arr[2],db);
		}
	case "DATABASE":
		dropDatabase(arr[2] + ".json");
	}
}
func dropTable(name string,db string){

}
func dropDatabase(name string){
	err := os.Remove("./data/" + name);
	if err != nil{
		fmt.Println("query error");
	}else{
		fmt.Println("query ok!");
	}
}