package code

import(
	"fmt"
	"io/ioutil"
	"encoding/json"
	s "strings"
)
type data struct{
	Datas []string
}
// insert into table name,age value
func Insert(arr []string,db string){
	
	if len(db) == 0{
		fmt.Println("query error");
	}else{
		index := -1;
		file,_ := ioutil.ReadFile("./data/" + db);
		DB := database{};
		_ = json.Unmarshal([]byte(file),&DB);
		// fmt.Println(DB);
		for i := 0 ; i < len(DB.Tables) ; i++{
			if arr[2] == DB.Tables[i].Name{
				index = i;
				break;
			}
		}
		if index == -1 {
			fmt.Println("table is not exists");
		}else{
			fields := s.Split(arr[3],",");
			content := s.Split(arr[5],",");
			for i := 0 ; i < len(fields) ; i++{
				if fields[i] != DB.Tables[index].Columns[i]{
					fmt.Println("field is not exists");
					return;
				} 
			}
			d := data{};
			tmp := []string{};
			for i := 0 ; i < len(DB.Tables[index].Columns) ; i++{
				if (i + 1) > len(content){
					tmp = append(tmp,"");
				}else{
					tmp = append(tmp,content[i]);
				}
			}
			d.Datas = tmp;
			
			list_data := DB.Tables[index].Rows;
			list_data = append(list_data,d);
			DB.Tables[index].Rows = list_data;
			
			datas,_ := json.MarshalIndent(DB,""," ");
			err := ioutil.WriteFile("./data/" + db,datas,0644);
			if err != nil{
				fmt.Println("query error");
			}else{
				fmt.Println("query ok");
			}
			// fmt.Println(tmp);
			
			// fmt.Println(fields);
			// fmt.Println(content);
		}
	}
}
