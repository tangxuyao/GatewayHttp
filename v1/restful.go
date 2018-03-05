package v1

import (
	"proto"
	"net/http"
	"time"
	"github.com/gorilla/mux"
	micro "github.com/micro/go-micro/client"
	"fmt"
	"encoding/json"
)

type Restful struct {
	version    string
	cliAccount account_api.AccountServiceClient
}

var restful Restful

func SetupHandler(r *mux.Router, cli micro.Client) {
	restful = Restful{
		version:    "v1",
		cliAccount: account_api.NewAccountServiceClient("account", cli),
	}

	r.HandleFunc("/"+restful.version+"/kv", restful.getAllKV)
	r.HandleFunc("/"+restful.version+"/signup", restful.signUp)
	r.HandleFunc("/"+restful.version+"/elk", restful.testElastic)

	//TODO：添加新的处理函数
}

func (rest *Restful) getAllKV(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	m := map[string]string{
		"version":  "1",
		"platform": "https",
	}

	rsp, err := json.Marshal(m)
	if err != nil {
		fmt.Fprintln(w, "json.Marshal()解析出错!!!")
		return
	}

	time.Sleep(time.Second * 2)
	fmt.Fprintln(w, string(rsp))
}

/**
account manager methods
 */
func (rest *Restful) signUp(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "不能使用'POST'之外的请求")
		return
	}

	phone := r.FormValue("phone")
	if len(phone) == 0 {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintln(w, "参数'phone'不能为空")
		return
	}

	gender := r.FormValue("gender")
	if len(gender) == 0 {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintln(w, "参数'gender'不能为空")
		return
	}

	//rest.cli_account.Signup(context.TODO(), &account_api.SignupReq{Name:"phone"})

	m := map[string]string{
		"token": "i'am god",
		"now":   fmt.Sprintf("%d", time.Now().Unix()),
	}

	rsp, err := json.Marshal(m)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "json.Marshal()解析出错!!!")
		return
	}

	fmt.Fprintln(w, string(rsp))
}

type Tweet struct {
	User     string    `json:"user"`
	PostDate time.Time `json:"postDate"`
	Message  string    `json:"message"`
}

func (rest *Restful) testElastic(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		name := r.FormValue("name")
		if len(name) == 0 {
			w.WriteHeader(http.StatusForbidden)
			fmt.Fprintln(w, "参数'name'不能为空")
			return
		}

		fmt.Fprintln(w, "fuck")

	default:
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "不能使用'POST'之外的请求")
	}
}
