// APIのサンプル AuthはJSONを受け取って返す
// エラー処理などはソース簡略化のため省略
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func main() {
	// portを環境変数から取得
	p := os.Getenv("HTTPD_PORT")
	if len(p) == 0 {
		p = "8080"
	}

	// ハンドラー設定
	http.HandleFunc("/Auth", Auth)

	// httpd開始
	log.Printf("httpd start port:%s", p)
	http.ListenAndServe(":"+p, nil)
}

// リクエスト
type ReqAuth struct {
	ID  string `json:"ID"`
	PWD string `json:"PWD"`
}

// レスポンス
type ResAuth struct {
	Status string `json:"status"`
}

func Auth(w http.ResponseWriter, r *http.Request) {
	//リクエスト処理
	req := &ReqAuth{}
	json.NewDecoder(r.Body).Decode(&req)
	log.Printf("Req ID:%s PWD:%s", req.ID, req.PWD)

	//本来は何か処理

	//レスポンス処理
	res := ResAuth{
		Status: "OK",
	}
	log.Printf("Res Status:%s", res.Status)
	bytes, _ := json.Marshal(res)
	w.Write(bytes)
}
