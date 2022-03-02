package code

import(
	"fmt"
	"io/ioutil"
	"encoding/json"
)
func Select(arr []string,db string){
	if len(db) == 0{
		fmt.Println("query error");
	}else{
		file,_ := ioutil.ReadFile("./data/" + db);
		DB := database{};
		_ = json.Unmarshal([]byte(file),&DB);
		index := -1;
		for i := 0 ; i < len(DB.Tables) ; i++{
			if arr[3] == DB.Tables[i].Name{
				index = i;
				break;
			}
		}
		if index == -1{
			fmt.Println("Table is not exists");
		}else{
			for i := 0 ; i < len(DB.Tables[index].Fields) ; i++{
				fmt.Print(DB.Tables[index].Fields[i].Name);
			}
			fmt.Println();
		    for i := 0 ; i < len(DB.Tables[index].Rows) ; i++{
		        for j := 0 ; j < len(DB.Tables[index].Rows[i].Datas) ; j++{
					fmt.Print(DB.Tables[index].Rows[i].Datas[j]);
				}   
				fmt.Println();
		    }
		}
	}
}
