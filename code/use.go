package code

import(
	"fmt"
	"io/ioutil"
)
func Use(arr []string) string{
	var tmp bool = false;
	arr[1] = arr[1] + ".json";
	files,_ := ioutil.ReadDir("./data");
	for _,file := range files{
		if file.Name() == arr[1]{
			tmp = true;
		}
	}
	if tmp == false{
		fmt.Println("databse is not exists");
		return "";
	}else{
		fmt.Println("query ok!");
		return arr[1];
	}
}