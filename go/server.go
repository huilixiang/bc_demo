package main

import (
	"fmt"
	"html/template"
	//"io/ioutil"
	"net/http"
	//"path/filepath"
	"math/rand"
	"strconv"
	"strings"
)

var baseDir = "/home/vagrant/mxcc/html"
var certMap map[string]string
var chainID = "mxcc"
var pList = make([]Product, 0, 3)

type loginPage struct {
	Title    string
	Username string
	Warning  string
	Body     []byte
}

type mechantPage struct {
	Title       string
	MechantName string
	Balance     string
	Body        []byte
	PList       []Product
}

type userPage struct {
	Title      string
	Username   string
	Balance    string
	RebateLink string
}

type Product struct {
	PID       string
	PName     string
	PPrice    string
	PBack     string
	PStrategy string
}

func myInit() {
	certMap = make(map[string]string)
	certMap["m_001"] = "123"
	certMap["u_001"] = "123"
	certMap["u_002"] = "123"
	certMap["u_003"] = "123"
	certMap["u_004"] = "123"
	certMap["u_005"] = "123"
	certMap["u_006"] = "123"
	certMap["u_007"] = "123"
	certMap["u_008"] = "123"
	rand.Seed(100)
	for i := 1; i < 4; i++ {
		si := strconv.Itoa(i)
		price := rand.Float64() * float64(i)
		cashback := fmt.Sprintf("%.2f", price*0.1)
		prd := Product{"p_00" + si, "商品" + si, fmt.Sprintf("%.2f", price), cashback, "Head4"}
		pList = append(pList, prd)
	}

}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	fmt.Printf("login info, username:%s, password:%s\n", username, password)
	v, ok := certMap[username]
	if ok && v == password {
		if strings.HasPrefix(username, "m") {
			mt := &mechantPage{MechantName: username, Balance: "100", PList: pList}
			t, err := template.ParseFiles(baseDir + "/mechant.html")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			err = t.Execute(w, mt)
			if err != nil {
				fmt.Println("mechant error......")
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		} else {

		}
	} else {
		warning := "用户名或密码错误"
		loginHtml := baseDir + "/login.html"
		page := &loginPage{Title: "登陆页面", Username: username, Warning: warning}
		t, err := template.ParseFiles(loginHtml)
		if err != nil {
			fmt.Println("template.ParseFiles err")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = t.Execute(w, page)
		if err != nil {
			fmt.Println("template execute error: %s\n", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func deposit(w http.ResponseWriter, r *http.Request) {
	merchant := r.FormValue("")
	fmt.Println(merchant)
}

func main() {
	myInit()
	InitChain(chainID)
	http.HandleFunc("/login.go", loginHandler)
	http.ListenAndServe(":8081", nil)
}
