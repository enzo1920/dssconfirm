package main

import (
  "fmt"
  "os"
  "flag"
  "strings"
  "net/http"
  "time"
  "encoding/json"
)

//  "io/ioutil"
//структура для получения первоночального токена
type Auth struct{
     AToken   string `json:"access_token"`
     Exp_in   uint64 `json:"expires_in"`
     TType    string `json:"token_type"`

}
//структура ответа
type Response struct {
        Challenge struct {
                Title struct {
                        Value string `json:"Value"`
                } `json:"Title"`
                TextChallenge []struct {
                        Label              string `json:"Label"`
                        ExpiresIn          int    `json:"ExpiresIn"`
                        CreatedAt          int    `json:"CreatedAt"`
                        ExpiresInSpecified bool   `json:"ExpiresInSpecified"`
                        IsHidden           bool   `json:"IsHidden"`
                        AuthnMethod        string `json:"AuthnMethod"`
                        RefID              string `json:"RefID"`
                        Title              string `json:"Title"`
                } `json:"TextChallenge"`
                ContextData struct {
                        RefID string `json:"RefID"`
                } `json:"ContextData"`
        } `json:"Challenge"`
        IsFinal bool `json:"IsFinal"`
        IsError bool `json:"IsError"`
}

type ResponseFinal struct {
	AccessToken string `json:"AccessToken"`
	ExpiresIn   int    `json:"ExpiresIn"`
	IsFinal     bool   `json:"IsFinal"`
	IsError     bool   `json:"IsError"`
}

func CriptoAuth (phone string,url string, client_id string)(authtoken string) {
  method := "POST"

  payload := strings.NewReader("grant_type=password&username="+phone+"&client_id="+client_id+"&resource=urn%3Acryptopro%3Adss%3Asignserver%3Askbkonturss&password=")

  client := &http.Client {
  }
  req, err := http.NewRequest(method, url, payload)

  if err != nil {
    fmt.Println(err)
    return
  }
  req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
  req.Header.Add("Cookie", "ASP.NET_SessionId=rojqot4xjumpq3afanivohwm")

  res, err := client.Do(req)
  if err != nil {
    fmt.Println(err)
    return
  }
  defer res.Body.Close()

  auth_struct := &Auth{}
  err = json.NewDecoder(res.Body).Decode(auth_struct)
  if err != nil {
    fmt.Println(err)
  }
  fmt.Printf("token: %s  \n",auth_struct.AToken)
  authtoken = auth_struct.AToken
  return authtoken
}

func StartReq(authtoken string, url string, client_id string)(refid string, iserr bool) {

  method := "POST"

  payload := strings.NewReader(`{`+""+`"Resource" : "urn:cryptopro:dss:signserver:skbkonturss",`+""+`"ClientId" : "`+client_id +`" ,`+""+`"ConfirmationScope" : "checkprofile"`+""+`}`)
  //fmt.Printf("payload: %s  \n", payload)
  client := &http.Client {
  }
  req, err := http.NewRequest(method, url, payload)

  if err != nil {
    fmt.Println(err)
    return
  }
  req.Header.Add("Content-Type", "application/json")
  req.Header.Add("Authorization", "Bearer "+authtoken)
  req.Header.Add("Cookie", "ASP.NET_SessionId=epmyzjtgbumjzu1scopqmesy")

//  fmt.Println(req)

  res, err := client.Do(req)
  if err != nil {
    fmt.Println(err)
    return
  }


  defer res.Body.Close()

/*  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(string(body))

*/
  response_struct := &Response{}
  err = json.NewDecoder(res.Body).Decode(response_struct)
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println(response_struct)
  fmt.Println(response_struct.Challenge.ContextData.RefID)
  fmt.Println(response_struct.IsError)
  refid = response_struct.Challenge.ContextData.RefID
  iserr = response_struct.IsError
  return refid,iserr
}
//функция проверки запроса

func ResponseCheck(authtoken string, url string, refid string)(isfinal bool, iserr bool) {

  method := "POST"

  payload := strings.NewReader(`{`+""+`"Resource" : "urn:cryptopro:dss:signserver:skbkonturss",`+""+`"ChallengeResponse" : {`+""+`
    "TextChallengeResponse" : [ {`+""+`"RefId" :"`+""+refid+""+`"} ]`+" "+`}`+" "+`}`)
  
//  fmt.Printf("payload: %s  \n", payload)
  client := &http.Client {
  }
  req, err := http.NewRequest(method, url, payload)

  if err != nil {
    fmt.Println(err)
    return
  }
  req.Header.Add("Content-Type", "application/json")
  req.Header.Add("Authorization", "Bearer "+authtoken)
  req.Header.Add("Cookie", "ASP.NET_SessionId=aql2nw3cfewfewrx5bu4zaqr")

  res, err := client.Do(req)
  if err != nil {
    fmt.Println(err)
    return
  }

/*  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(string(body))
*/

  defer res.Body.Close()

  responsefinal_struct := &Response{}
  err = json.NewDecoder(res.Body).Decode(responsefinal_struct)
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println(responsefinal_struct.IsFinal)
  isfinal = responsefinal_struct.IsFinal
  iserr = responsefinal_struct.IsError
  return isfinal,iserr
}


func main() {

///URLS
url_auth := "https://"
url_req := "https://"
client_id := "client_id"
//аргументы для запуска
  var msisdn string
//  var refid string
//  var iserr bool
// flag declaration
  flag.StringVar(&msisdn,"m","","Specify msisdn.")

  flag.Parse()
  if len(os.Args) == 1 {
     fmt.Printf("Usage: \n")
     fmt.Printf("./dssconfirm -m msisdn \n")
     os.Exit(0)
  } else{
       token := CriptoAuth(msisdn,url_auth,client_id)
       // подождем 3 секунды перед запросом
       time.Sleep(5 * time.Second)
       refid, err := StartReq(token,url_req,client_id)
       fmt.Printf("i've got ref: %v, error %v \n", refid, err)
       time.Sleep(5 * time.Second)
       final, errf := ResponseCheck(token,url_req,refid)
       fmt.Printf("i've got final: %v , error %v \n", final, errf)
  }


}
