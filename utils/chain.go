package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"ginVue/utils/http"
	"github.com/spf13/viper"
	"log"
)

var (
	hostname string
	port     string
	ledger   string
	name     string
)

func ChainInit() {
	hostname = viper.GetString("chain.hostname")
	port = viper.GetString("chain.port")
	ledger = viper.GetString("chain.ledger")
	name = viper.GetString("chain.name")
}

type QueryResp struct {
	State   int    `json:"state"`
	Message string `json:"message"`
	Data    struct {
		Result string `json:"result"`
	} `json:"data"`
}
type QueryReq struct {
	Version  string        `json:"version"`
	Category string        `json:"category"`
	Method   string        `json:"method"`
	Args     interface{}    `json:"args"`
	Name     string        `json:"name"`
}

func Invoke(method string, arg interface{}) error {
	chainUrl := fmt.Sprintf("https://%s:%s/v2/tx/sc/invoke?ledger=%s", hostname, port, ledger)
	args := arg
	QR := QueryReq{
		Category: "wvm",
		Name:     name,
		Method:   method,
		Args:     args,
	}
	body, _ := json.Marshal(QR)

	fmt.Println("打印Body")
	fmt.Println(string(body))

	fmt.Println()
	_, res, err := http.NewClient().Post(chainUrl, string(body))
	if err != nil {
		fmt.Errorf("POST url[%v] failed, err:%v", chainUrl, err)
		return err
	}
	var resp QueryResp
	json.Unmarshal(res, &resp)
	fmt.Printf("%+v\n", resp)
	if resp.State != 200 {
		return errors.New("调用失败！")
	} else {
		return nil
	}
}

func Query(method string, arg interface{}) (error, string) {
	chainUrl := fmt.Sprintf("https://%s:%s/v2/tx/sc/query?ledger=%s", hostname, port, ledger)
	args := arg
	QR := QueryReq{
		Category: "wvm",
		Name:     name,
		Method:   method,
		Args:     args,
	}
	body, _ := json.Marshal(QR)
	log.Println("发送请求body为")
	log.Println(string(body))
	_, res, err := http.NewClient().Post(chainUrl, string(body))
	if err != nil {
		fmt.Errorf("POST url[%v] failed, err:%v", chainUrl, err)
		return err, ""
	}
	var resp QueryResp
	json.Unmarshal(res, &resp)
	//fmt.Printf("%+v\n", resp)
	if resp.State != 200 {
		return errors.New("调用失败-----"+resp.Message), ""
	} else {
		return nil, resp.Data.Result
	}
}
