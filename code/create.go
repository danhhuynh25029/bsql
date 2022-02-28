package code

import(
	"fmt"
	// "encoding/json"
	"io/ioutil"
	s "strings"
	"os"
	"encoding/json"
)
type 
func Create(arr []string,db string){
	if len(db) == 0{
		fmt.Println("query error!");
	}else{
		arr[1] = s.ToUpper(arr[1])
			switch arr[1]{
			case "TABLE":
				createTable(arr,db);
			case "DATABASE":
				createDatabase(arr[2]+".json");
			}
	}
}
func createTable(arr []string,db string){
	content,_ :=  ioutil.ReadFile("./data/" + db);
	for _,name := range json{
		if arr[2] == name{
			fmt.Println("table is exists");
			return;
		}
		fmt.Println(name);
	}
	// _ := ioutil.WriteFile("./data/"+db,"hello",0644);	
	// fmt.Println(json);
	// defer json.Close();
}

func createDatabase(name string){
	files,_ := ioutil.ReadDir("./data");
	for _,file := range files{
		if file.Name() == name{
			fmt.Println("database is exists");
			return;
		}
	}
	file,err := os.Create("./data/" + name);
	if err != nil{
		fmt.Println("err");
	}else{
		fmt.Println("query ok!");
	}
	defer file.Close();
}