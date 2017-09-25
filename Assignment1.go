package main

import (
"fmt"
"net/http"
"os"
"encoding/json"
"io/ioutil"

)

type Data struct {
Project 		string 	`json:"repository"`
Owner			string 	`json:"owner"`
TopContributor 	string 	`json:"topContributor"`
Contributors	int		`json:"commits"`
Languages		[]string`json:"languages"`


}
//struct to find the repository name and the owner
type Repository struct{

	Fullname string `json:"full_name"`
	Owner struct {
		Username string `json:"login"`
	}	`json:"owner"`
}

//struct to find the top contributor 
type Contributors struct{
	Login 			string 	`json:"login"`
	Contributions	int		`json:"contributions"`
}
/*
TRIED TO MAKE THIS TO TEST IT BUT URL WAS UDEFINED
func getUrl(x int) string {
	switch x {
	case 1:
		url := "https://api.github.com/repos/golang/go"
	case 2:
		url := "https://api.github.com/repos/golang/go/contributors"
	case 3:
		url := "https://api.github.com/repos/golang/go/languages"
	}
	return url	
}
*/

//Handler function
func serveRest(w http.ResponseWriter, r *http.Request) {
	//urls for repositoies 
	url1 := "https://api.github.com/repos/golang/go"
	url2 := "https://api.github.com/repos/golang/go/contributors"
	url3 :=	"https://api.github.com/repos/golang/go/languages"

//making map to get languages
	v := make(map[string]int)

//Repository
repo, err1 := http.Get(url1)
if err1 != nil {
	panic(err1)
} 
defer repo.Body.Close()

//Contributor
cont, err2 := http.Get(url2)
if err2 != nil {
	panic(err2)
} 
defer cont.Body.Close()

//Languages
lang, err3 := http.Get(url3)
if err3 != nil {
	panic(err3)
} 
defer lang.Body.Close()

//Body for repository
body1, err1 := ioutil.ReadAll(repo.Body)
if err1 != nil {
	panic(err1)
}

//Body for contributor
body2, err2 := ioutil.ReadAll(cont.Body)
if err2 != nil {
	panic(err2)
}

//Body for languages
body3, err3 := ioutil.ReadAll(lang.Body)
if err3 != nil {
	panic(err3)
}

//variable to get the repository owner and name
var p Repository

//variable to get the top contributor and the numer of contributions
var c []Contributors


//Unmarshall repository
err1 = json.Unmarshal(body1, &p)
if err1 != nil {
	panic(err1)
}

//Unmarshall contributor
err2 = json.Unmarshal(body2, &c)
if err2 != nil {
	panic(err2)
}

//Unmarshall languages
err3 = json.Unmarshal(body3, &v)
if err3 != nil {
	panic(err3)
}

//making git so that there will be no space between / and start of p.Fullname
git := ("github.com/" + p.Fullname)

//Printing out the owner and the name of the repository
fmt.Fprintln(w, "Owner:", p.Owner.Username, "\nRepository: ", git)

//Printing out the Top commiter and their number of contributions by acsessing the first item in the array 
fmt.Fprintln(w, "Top Commiter:", c[0].Login, "\nNumber of contributions:", c[0].Contributions)
fmt.Fprintln(w, "\nLanguages: ")

//Printing out all the different languages used in the repository
for key := range v {
	fmt.Fprint(w, key," ")
}


}

func main() {
	
	// setting up the server
	http.HandleFunc("/",serveRest)
	fmt.Println("listening...")
	err := http.ListenAndServe(":"+os.Getenv("PORT"),nil)
	if err != nil{
		panic(err)
	}
	
}