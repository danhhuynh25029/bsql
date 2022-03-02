package code

import(
	"fmt"
	s "strings"
	"os"
	"io/ioutil"
	"encoding/json"
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
	index := -1
	file,_ := ioutil.ReadFile("./data/" + db );
	DB := database{};
	_ = json.Unmarshal([]byte(file),&DB);
	for i := 0 ; i < len(DB.Tables) ; i++{
		if name == DB.Tables[i].Name {
			index = i;
			break;
		}
	}
	if index == -1{
		fmt.Println("table is not exists");
	}else{
		tables := []table{};
		for i := 0 ; i < len(DB.Tables) ; i++{
			if i != index{
				tables = append(tables,DB.Tables[i]);
			}
		}
		DB.Tables = tables;
		data,_ := json.MarshalIndent(DB,""," ");
		err := ioutil.WriteFile("./data/" + db,data,0644);
		if err != nil{
			fmt.Println("query error");
		}else{
			fmt.Println("query ok!");
		}
	}
}
func dropDatabase(name string){
	err := os.Remove("./data/" + name);
	if err != nil{
		fmt.Println("query error");
	}else{
		fmt.Println("query ok!");
	}
}
