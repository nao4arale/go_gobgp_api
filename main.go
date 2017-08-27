package main

import (
//  "bytes"
//  "encoding/json"
  "strconv"
  "io/ioutil"
  "strings"
  "time"
  "net/http"
  "github.com/gorilla/mux"
  jwt "github.com/dgrijalva/jwt-go"
 "github.com/auth0/go-jwt-middleware"
    "os"
    "log"
    "github.com/joho/godotenv"
    "fmt"
    "github.com/codeskyblue/go-sh"
    "github.com/mattn/go-shellwords"
    "github.com/mattn/go-scan"
)

func Env_load() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")

    }
}

var StatusHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
if checkAuth(r) == false {
        w.Header().Set("WWW-Authenticate", `Basic realm="MY REALM"`)
        w.WriteHeader(401)
        w.Write([]byte("401 Unauthorized\n"))
        return
    } else {
  w.Write([]byte("API is up and running\n"))
 }
})

var JwkStatusHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
 w.Write([]byte("JWK is up and running\n"))
})

//var TestHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
//w.Write([]byte("Test"))})

var ReceiveCommandHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
  //Validate request
  if r.Method != "POST" {
    w.WriteHeader(http.StatusBadRequest)
    return
  }

  if r.Header.Get("Content-Type") != "application/json" {
    w.WriteHeader(http.StatusBadRequest)
    return
  }
/*
  //To allocate slice for request body
  length, err := strconv.Atoi(r.Header.Get("Content-Length"))
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    return
  }

  body := make([]byte, length)
  length, err = r.Body.Read(body)
  if err != nil && err != io.EOF {
    w.WriteHeader(http.StatusInternalServerError)
    return
  }

  //parse json
  var jsonBody map[string]interface{}
  err = json.Unmarshal(body[:length], &jsonBody)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    return
  }
*/
        var js = strings.NewReader(execute(r))
        var s string

        if err := scan.ScanJSON(js, "/command/", &s); err != nil {
        }
    runCmdStr(s)
  w.WriteHeader(http.StatusOK)
})

func execute(resp *http.Request) string {
        // request bodyを文字列で取得
        // ioutil.ReadAllを使う
        b, err := ioutil.ReadAll(resp.Body)
        if err == nil {
                return string(b)
        }
        return ""
}


func runCmdStr(cmdstr string) error {
    // 文字列をコマンド、オプション単位でスライス化する
    c, err := shellwords.Parse(cmdstr)
    if err != nil {
        return err
    }
    switch len(c) {
    case 0:
        // 空の文字列が渡された場合
        return nil
    case 1:
        // コマンドのみを渡された場合
        fmt.Println(c[0])
        err = sh.Command(c[0]).Run()
    default:
        // コマンド+オプションを渡された場合
        // オプションは可変長でexec.Commandに渡す
	// https://golang.org/doc/faq#convert_slice_of_interface
        s := make([]interface{}, len(c))
        for i, v := range c {
        s[i] = v
        }
        err = sh.Command(c[0], s[1:]...).Run()
	}
    if err != nil {
        return err
    }
    return nil
}

  /* Set up a global string for our secret */
var mySigningKey = []byte(os.Getenv("SECRET"))

  /* Handlers */
var GetTokenHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
   if checkAuth(r) == false {
        w.Header().Set("WWW-Authenticate", `Basic realm="MY REALM"`)
        w.WriteHeader(401)
        w.Write([]byte("401 Unauthorized\n"))
        return
    } else {
    /* Create the token */
    token := jwt.New(jwt.SigningMethodHS256)

    /* Create a map to store our claims*/
    claims := token.Claims.(jwt.MapClaims)

    /* Set token claims */
    const (
	expiredTime int = 72
	)
    claims["admin"] = true
    claims["name"] = "nya hoke"
    claims["exp"] = time.Now().Add(time.Hour * time.Duration(expiredTime)).Unix()
    /* Sign the token with our secret */
    tokenString, _ := token.SignedString(mySigningKey)
    expiredTime_s := strconv.Itoa(expiredTime)
//    expiredTime = strconv.Itoa(expiredTime)
    /* Finally, write the token to the browser window */
    TokenData := `{"token":` + `"` + tokenString + `", ` + `"expired":"` + expiredTime_s +`"}`
    w.Write([]byte(TokenData+"\n"))

    }
  })

var jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
  ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
    return mySigningKey, nil
  },
  SigningMethod: jwt.SigningMethodHS256,
})

func checkAuth(r *http.Request) bool {
    Env_load()
    username, password, ok := r.BasicAuth()
    if ok == false {
        return false
    }
    return username == os.Getenv("USERNAME") && password == os.Getenv("PASSWORD")
}

func main() {

  r := mux.NewRouter()

  r.Handle("/api/token", GetTokenHandler).Methods("GET")
  r.Handle("/api/status", StatusHandler).Methods("GET")
  r.Handle("/api/jwkstatus",  jwtMiddleware.Handler(JwkStatusHandler)).Methods("GET")
//  r.Handle("/api/command", ReceiveCommandHandler).Methods("POST")
 r.Handle("/api/command", jwtMiddleware.Handler(ReceiveCommandHandler)).Methods("POST")
// r.Handle("test", TestHandler).Methods("GET")

  http.ListenAndServe(":3000", r)
}
