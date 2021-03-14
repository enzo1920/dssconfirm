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
     Exp_in   int `json:"expires_in"`
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

//export CryptoAuth
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
	
  if res != nil {
        defer res.Body.Close()
  }
	
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

  auth_struct := &Auth{}
  err = json.NewDecoder(res.Body).Decode(auth_struct)
  if err != nil {
    fmt.Println(err)
  }
  fmt.Printf("token: %s  \n",auth_struct.AToken)
  authtoken = auth_struct.AToken


 //defer res.Body.Close()


  return authtoken
}

//export StartReq
func StartReq(authtoken string, url string, client_id string)(refid string, iserr bool) {
  t :=time.Now()
  method := "POST"

  //payload := strings.NewReader(`{`+""+`"Resource" : "urn:cryptopro:dss:signserver:skbkonturss",`+""+`"ClientId" : "`+client_id +`" ,`+""+`"ConfirmationScope" : "checkprofile"`+""+`}`)
  
  payload := strings.NewReader(`{`+""+`"Resource" : "urn:cryptopro:dss:signserver:skbkonturss",`+""+`"ClientId" : "`+client_id +`",`+""+`"ConfirmationScope" : "checkprofile",`+""+`"ConfirmationParams": {`+""+`"CpTime":"`+t.Format("17.01.2018 17:54:02") +`"`+""+`}`+""+``+""+` }`)
  
  //fmt.Printf("payload: %s  \n", payload)
  client := &http.Client {
  }
  req, err := http.NewRequest(method, url, payload)
	
  if req != nil {
        defer req.Body.Close()
  }
	
  if err != nil {
    fmt.Println(err)
    return
  }
  req.Header.Add("Content-Type", "application/json")
  req.Header.Add("Authorization", "Bearer "+authtoken)
  req.Header.Add("Cookie", "ASP.NET_SessionId=epmyzjtgbumjzu1scopqmesy")

//  fmt.Println(req)

  res, err := client.Do(req)
	
  if res != nil {
        defer res.Body.Close()
  }
	
  if err != nil {
    fmt.Println(err)
    return
  }


//  defer res.Body.Close()

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

//export ResponseCheck
func ResponseCheck(authtoken string, url string, refid string)(isfinal bool, iserr bool) {

  method := "POST"

  payload := strings.NewReader(`{`+""+`"Resource" : "urn:cryptopro:dss:signserver:skbkonturss",`+""+`"ChallengeResponse" : {`+""+`
    "TextChallengeResponse" : [ {`+""+`"RefId" :"`+""+refid+""+`"} ]`+" "+`}`+" "+`}`)
  
//  fmt.Printf("payload: %s  \n", payload)
  client := &http.Client {
  }
  req, err := http.NewRequest(method, url, payload)
	
  if req != nil {
        defer req.Body.Close()
  }
	
  if err != nil {
    fmt.Println(err)
    return
  }
  req.Header.Add("Content-Type", "application/json")
  req.Header.Add("Authorization", "Bearer "+authtoken)
  req.Header.Add("Cookie", "ASP.NET_SessionId=aql2nw3cfewfewrx5bu4zaqr")

  res, err := client.Do(req)
	
  if res != nil {
        defer res.Body.Close()
  }
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

  //defer res.Body.Close()

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

//exit code 
  exit_code := 0
///URLS
url_auth := "https://"
url_req := "https://"
client_id := "client1111"
//аргументы для запуска
  var msisdn string
  var final bool
  var errf bool
//  var refid string
//  var iserr bool
// flag declaration
  flag.StringVar(&msisdn,"m","","Specify iccid.")

  flag.Parse()
  if len(os.Args) == 1 {
     fmt.Printf("Usage: \n")
     fmt.Printf("./dssconfirm -m iccid \n")
     exit_code = -1 
     os.Exit(exit_code)
  } else{
       token := CriptoAuth(msisdn,url_auth,client_id)
       if len(token)==0{
          fmt.Printf("Token len: %v, token: %v \n", len(token), token)
          exit_code = 1 
          os.Exit(exit_code)
       }
       // подождем 3 секунды перед запросом
       time.Sleep(3 * time.Second)
       refid, err := StartReq(token,url_req,client_id)
       fmt.Printf("i've got ref: %v, error %v \n", refid, err)
       time.Sleep(3 * time.Second)
       for {
            time.Sleep(2 * time.Second)
            final, errf = ResponseCheck(token,url_req,refid)
            fmt.Printf("i've got Final: %v , Error %v \n", final, errf)
            if final==true{
               break
               os.Exit(exit_code)
            }
       }
       //fmt.Printf("i've got final: %v , error %v \n", final, errf)
  }


}
