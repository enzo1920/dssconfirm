package main

import (
  "fmt"
  "log"
  "os"
  "flag"
  "strings"
  "net/http"
  "time"
  "encoding/json"
)

//структура конфига 
type Configuration struct {
	Url_auth string `json:"url_auth"`
        Url_req string  `json:"url_req"`
        Client_id string  `json:"client_id"`
	Log_file_name string  `json:"log_file_name"`
}




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
func CriptoAuth (username string,url string, client_id string)(authtoken string) {
  method := "POST"

  payload := strings.NewReader("grant_type=password&username="+username+"&client_id="+client_id+"&resource=urn%3Acryptopro%3Adss%3Asignserver%3Aa1ss&password=")

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
<<<<<<< HEAD
 
/*  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    fmt.Println(err)
    return
  }
=======
>>>>>>> d197204f645f71773d6cc823280f3a0157d2c1f5
  fmt.Println(string(body))
*/

  auth_struct := &Auth{}
  err = json.NewDecoder(res.Body).Decode(auth_struct)
  if err != nil {
    log.Println(err)
  }
  log.Printf("token: %s  \n",auth_struct.AToken)
  authtoken = auth_struct.AToken


<<<<<<< HEAD
 defer res.Body.Close()
=======
<<<<<<< HEAD
 defer res.Body.Close()
=======
 //defer res.Body.Close()
>>>>>>> 7817ccb05ba9036137ba40370db74864350ae0b3
>>>>>>> d197204f645f71773d6cc823280f3a0157d2c1f5


  return authtoken
}

//export StartReq
func StartReq(authtoken string, url string, client_id string)(refid string, iserr bool) {

  method := "POST"

  payload := strings.NewReader(`{`+""+`"Resource" : "urn:cryptopro:dss:signserver:a1ss",`+""+`"ClientId" : "`+client_id +`" ,`+""+`"ConfirmationScope" : "checkprofile"`+""+`}`)
  //fmt.Printf("payload: %s  \n", payload)
  client := &http.Client {
  }
  req, err := http.NewRequest(method, url, payload)
	
  if req != nil {
        defer req.Body.Close()
  }
	
  if err != nil {
    log.Println(err)
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
    log.Println(err)
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
  log.Println("resp struct:", response_struct)
  log.Println("REFID: ",response_struct.Challenge.ContextData.RefID)
  log.Println("IsError refid: ",response_struct.IsError)
  refid = response_struct.Challenge.ContextData.RefID
  iserr = response_struct.IsError
  return refid,iserr
}
//функция проверки запроса

//export ResponseCheck
func ResponseCheck(authtoken string, url string, refid string)(isfinal bool, iserr bool) {

  method := "POST"

  payload := strings.NewReader(`{`+""+`"Resource" : "urn:cryptopro:dss:signserver:a1ss",`+""+`"ChallengeResponse" : {`+""+`
    "TextChallengeResponse" : [ {`+""+`"RefId" :"`+""+refid+""+`"} ]`+" "+`}`+" "+`}`)
  
//  fmt.Printf("payload: %s  \n", payload)
  client := &http.Client {
  }
  req, err := http.NewRequest(method, url, payload)
	
  if req != nil {
        defer req.Body.Close()
  }
	
  if err != nil {
    log.Println("err: ",err)
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
    log.Println("err:",err)
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
    log.Println("err:",err)
  }
  log.Println("is final: ",responsefinal_struct.IsFinal)
  isfinal = responsefinal_struct.IsFinal
  iserr = responsefinal_struct.IsError
  return isfinal,iserr
}




//func config_reader(cfg_file string)([]string){
func Config_reader(cfg_file string) Configuration {

	//c := flag.String("c", cfg_file, "Specify the configuration file.")
	//flag.Parse()
	file, err := os.Open(cfg_file)
	if err != nil {
		fmt.Println("can't open config file: ", err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	Config := Configuration{}
	err = decoder.Decode(&Config)
	if err != nil {
		fmt.Println("can't decode config JSON: ", err)
	}

	return Config
}


func main() {

//exit code 
  exit_code := 0
<<<<<<< HEAD
//аргументы для запуска
  var username string
  var final bool
  var errf bool


//************************* read config ******************************************//
   cfg := Config_reader("../config/dss.conf")



//*********************** URLS from config fiel **********************************//
   url_auth := cfg.Url_auth 
   url_req := cfg.Url_req
   client_id := cfg.Client_id
   //log file create 
   //logging
   log_dir := "./log"
   if _, err := os.Stat(log_dir); os.IsNotExist(err) {
		os.Mkdir(log_dir, 0644)
   }
   file, err := os.OpenFile("./log/"+cfg.Log_file_name, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
   if err != nil {
		log.Fatal(err)
   }
   defer file.Close()
   log.SetOutput(file)
   log.Println("Logging to a file dssconfirm!")


// start programm dss

=======
///URLS
<<<<<<< HEAD
//url_auth := "https://stenddss.cryptopro.ru/skbkonturidp/oauth/token"
//url_req := "https://stenddss.cryptopro.ru/skbkonturidp/confirmation"
//client_id := ""

url_auth := "https://stenddss.cryptopro.ru/a1idp/oauth/token"
url_req := "https://stenddss.cryptopro.ru/a1idp/confirmation"

client_id := ""


//аргументы для запуска
  var username string
=======
url_auth := "https://"
url_req := "https://"
client_id := "client"
//аргументы для запуска
  var msisdn string
>>>>>>> 7817ccb05ba9036137ba40370db74864350ae0b3
  var final bool
  var errf bool
//  var refid string
//  var iserr bool
// flag declaration
>>>>>>> d197204f645f71773d6cc823280f3a0157d2c1f5
  flag.StringVar(&username,"m","","Specify profile name.")

  flag.Parse()
  if len(os.Args) == 1 {
     fmt.Printf("Usage: \n")
<<<<<<< HEAD
     fmt.Printf("./dssconfirm -m profile_name \n")
     exit_code = -1 
     os.Exit(exit_code)
  } else{
//step 1 get token

       token := CriptoAuth(username,url_auth,client_id)
       if len(token)==0{
          log.Printf(" step1. Token len: %v, token: %v \n", len(token), token)
=======
<<<<<<< HEAD
     fmt.Printf("./dssconfirm -m profile_name \n")
     exit_code = -1 
     os.Exit(exit_code)
  } else{
       token := CriptoAuth(username,url_auth,client_id)
=======
     fmt.Printf("./dssconfirm -m msisdn \n")
     exit_code = -1 
     os.Exit(exit_code)
  } else{
       token := CriptoAuth(msisdn,url_auth,client_id)
>>>>>>> 7817ccb05ba9036137ba40370db74864350ae0b3
       if len(token)==0{
          fmt.Printf("Token len: %v, token: %v \n", len(token), token)
>>>>>>> d197204f645f71773d6cc823280f3a0157d2c1f5
          exit_code = 1 
          os.Exit(exit_code)
       }
       // подождем 3 секунды перед запросом
       time.Sleep(3 * time.Second)
<<<<<<< HEAD
//step 2 get refid

       refid, err := StartReq(token,url_req,client_id)
       if err ==true{
          log.Printf("step2. i've got error=true in startreq: %v, token: %v \n", len(token), token)
          exit_code = 2 
          os.Exit(exit_code)
       }
       log.Printf("step2. I've got ref: %v, error %v \n", refid, err)
       time.Sleep(3 * time.Second)
       //counter если достигли 30 выходим
       counter := 0
       for {
            time.Sleep(2 * time.Second)
//step 3 response check

            final, errf = ResponseCheck(token,url_req,refid)
            log.Printf("i've got Final: %v , Error %v \n", final, errf)
            if final==true &&errf == false{
               break
               os.Exit(exit_code)
            }
            if counter == 30{
               exit_code = 3 
               break
               os.Exit(exit_code)
            }
            counter++
=======
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
>>>>>>> d197204f645f71773d6cc823280f3a0157d2c1f5
       }
       //fmt.Printf("i've got final: %v , error %v \n", final, errf)
  }


}
