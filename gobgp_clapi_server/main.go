package main

import (
  "strconv"
  "io/ioutil"
  "strings"
  "time"
  "net/http"
  "fmt"
  "os"
  "log"
  "github.com/gorilla/mux"
  jwt "github.com/dgrijalva/jwt-go"
  "github.com/auth0/go-jwt-middleware"
  "github.com/joho/godotenv"
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
/* This HTTP Handler Thank you for...               */
/* https://auth0.com/blog/authentication-in-golang/ */

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

var JwtStatusHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("JWT is up and running\n"))
})

/* Debug Hundler
var TestHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
w.Write([]byte("Test"))})
*/

var ReceiveCommandHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
  //Validate request
	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
/*
	if r.Header.Get("Content-Type") != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
*/
	w.Header().Set("Content-type", "application/json")

	var js = strings.NewReader(execute(r))
	var s string
		if err := scan.ScanJSON(js, "/command/", &s); err != nil {
			}

	runCmdStr(s)

	w.WriteHeader(http.StatusOK)

})

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
    w.Header().Set("Content-type", "application/json")

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

func execute(resp *http.Request) string {
        // Getting request body as string.
        b, err := ioutil.ReadAll(resp.Body)
        if err == nil {
                return string(b)
        }
        return ""
}


func runCmdStr(cmdstr string) error {
    // Receiving post string which is shell command, And Slice Parse. 
    c, err := shellwords.Parse(cmdstr)
    if err != nil {
        return err
    }
    switch len(c) {
    case 0:
        // nil
        return nil
	fmt.Println("No Command?")
    case 1:
        // no option
	// for Debug
        fmt.Println(c[0])
        err = sh.Command(c[0]).Run()
	fmt.Println(" ### done. ###")
    default:
	/*                                                             */
        /* one or more option(It's main).                              */
        /* shellwords.Parse() -> string[] -> interface{} ->sh.Command()*/
	/* Thank you for...                                            */
	/* http://qiita.com/tanksuzuki/items/9205ff70c57c4c03b5e5      */
        /* https://golang.org/doc/faq#convert_slice_of_interface       */
	/*							       */
	fmt.Print(c[:])
	s := make([]interface{}, len(c))
		for i, v := range c {
		s[i] = v
		}
        err = sh.Command(c[0], s[1:]...).Run()
	fmt.Println(" ### done. ###")
	}
    if err != nil {
        return err
    }
    return nil
}


func main() {

  r := mux.NewRouter()

  r.Handle("/api/token", GetTokenHandler).Methods("GET")
  r.Handle("/api/status", StatusHandler).Methods("GET")
  r.Handle("/api/jwtstatus",  jwtMiddleware.Handler(JwtStatusHandler)).Methods("GET")
//  r.Handle("/api/command", ReceiveCommandHandler).Methods("POST")
 r.Handle("/api/command", jwtMiddleware.Handler(ReceiveCommandHandler)).Methods("POST")
// r.Handle("test", TestHandler).Methods("GET")

  http.ListenAndServe(":3000", r)
}

