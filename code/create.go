package code

import(
	"fmt"
	// "encoding/json"
	"io/ioutil"
	s "strings"
	"os"
	"encoding/json"
)
type database struct{
	Tables []table
}
type table struct{
	Name string 
	Fields []field
	Columns []string
	Rows []data
}
type field struct{
	Name string
	Type string
}
func Create(arr []string,db string){
	
		arr[1] = s.ToUpper(arr[1])
			switch arr[1]{
				case "TABLE":
					if len(db) == 0 {
						fmt.Println("query error!");
					}else{
						createTable(arr,db);
					}
				case "DATABASE":
					createDatabase(arr[2]+".json");
			}
}
func createTable(arr []string,db string){
	file,_ :=  ioutil.ReadFile("./data/" + db);
	data := database{}
	_ = json.Unmarshal([]byte(file),&data);
	// fmt.Println(data.Tables);
	for i := 0 ; i < len(data.Tables) ; i++{
		// fmt.Println(data.Tables[i].Name);
		if data.Tables[i].Name == arr[2]{
			fmt.Println("table is exists");
			return;
		}
	}

	
	// var list_table []table;

	tb := table{};
	tb.Name = arr[2];
	var list_field []field;
	var tmp1 []string;
	for i := 3 ; i < len(arr) ; i++{
		tmp := s.Split(arr[i],":");
		f  := field{Name:tmp[0],Type:tmp[1]};
		tmp1 = append(tmp1,tmp[0]);
		list_field = append(list_field,f);
	}
	tb.Fields = list_field;
	tb.Columns = tmp1;
	// list_table = append(list_table,tb);
	data.Tables = append(data.Tables,tb);
	content,_ := json.MarshalIndent(data,""," ");
	err := ioutil.WriteFile("./data/" + db,content,0644);
	if err != nil{
		fmt.Println("query error");
	}else{
		fmt.Println("query ok");
	}
	// fmt.Println(content);

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
	data := database{};
	content,_ := json.Marshal(data);
	_ = ioutil.WriteFile("./data/" + name,content,0644);
}