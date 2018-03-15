package v1

import (
	"proto/crm"
	"net/http"
	"time"
	"github.com/gorilla/mux"
	"fmt"
	"encoding/json"
	"golang.org/x/net/context"
	"github.com/micro/go-micro/client"
	"proto/gm"
)

type Restful struct {
	version string
	crm     crm_api.CRMServiceClient
	gm      gm_api.GameServiceClient
}

var restful Restful

func SetupHandler(r *mux.Router, c client.Client) {
	restful = Restful{
		version: "v1",
		crm:     crm_api.NewCRMServiceClient("crmService", c),
		gm:      gm_api.NewGameServiceClient("GameService", c),
	}

	r.HandleFunc("/"+restful.version+"/kv", restful.getAllKV)
	r.HandleFunc("/"+restful.version+"/signup", restful.signUp)
	r.HandleFunc("/"+restful.version+"/startgame", restful.startGame)
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

	//phone := r.FormValue("phone")
	//if len(phone) == 0 {
	//	w.WriteHeader(http.StatusForbidden)
	//	fmt.Fprintln(w, "参数'phone'不能为空")
	//	return
	//}
	//
	//gender := r.FormValue("gender")
	//if len(gender) == 0 {
	//	w.WriteHeader(http.StatusForbidden)
	//	fmt.Fprintln(w, "参数'gender'不能为空")
	//	return
	//}

	ret, err := rest.crm.Signup(context.TODO(), &crm_api.SignupReq{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "[gateway] 无法执行注册账号:%s\n", err)
		return
	}

	m := map[string]string{
		"token": ret.Token,
		"id":    ret.ID,
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

func (rest *Restful) startGame(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "不能使用'POST'之外的请求")
		return
	}

	token := r.FormValue("token")
	if len(token) == 0 {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintln(w, "参数'token'不能为空")
		return
	}

	name := r.FormValue("name") // 参数可以为空

	_, err := rest.gm.StartGame(context.TODO(), &gm_api.StartGameReq{Token: token, Name: name})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ret := map[string]string{
			"code":    "1",
			"message": fmt.Sprintf("%s", err),
			"now":     fmt.Sprintf("%d", time.Now().Unix()),
		}
		rsp, _ := json.Marshal(ret)
		fmt.Fprintln(w, string(rsp))
		return
	}

	m := map[string]string{
		"now": fmt.Sprintf("%d", time.Now().Unix()),
	}

	rsp, err := json.Marshal(m)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "json.Marshal()解析出错!!!")
		return
	}

	fmt.Fprintln(w, string(rsp))
}
