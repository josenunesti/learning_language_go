package main

import (
	"encoding/json"
	"fmt"
)

// Também poderia ser criado um slice de struct ex: "pessoa []struct" assim não seria necessário declarar um slice de pessoas conforme o block-code main
// Também seria possivel adicionar as tags na struct para mapear os campos do json com os do struct
type pessoa struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func main() {
	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
	fmt.Println(s)

	var pessoas []pessoa

	err := json.Unmarshal([]byte(s), &pessoas)

	if err != nil {
		fmt.Println(err)
	}
	
	for _, v := range pessoas {
		fmt.Println(v)
	}
}
