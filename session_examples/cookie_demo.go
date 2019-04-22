package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"github.com/gomodule/redigo/redis"
	"time"
)

var redis_cache redis.Conn

var users  = map[string]string{
	"user1":"password1",
	"user2":"password2",
}

type Creds struct{
	Password string `json:"password"`
	Username string `json:"username"`
}

func init(){
	conn, err:= redis.DialURL("redis://localhost")
	if err != nil{
		panic(err)
	}
	redis_cache = conn
}

func signin(w http.ResponseWriter, r *http.Request){
	var creds Creds
	err:= json.NewDecoder(r.Body).Decode(&creds)
	if err!=nil{
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("Error in signin while Decoding",err)
		return
	}

	mypass, status:= users[creds.Username]
	if !status || mypass != creds.Password{
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
        
	//Generating the cookie ID
	sessionToken:= uuid.New().String()

	_, err= redis_cache.Do("SETEX", sessionToken, 120, creds.Username)
	if err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name: "session_token",
		Value: sessionToken,
		Expires: time.Now().Add(120 * time.Second),
	})

}


func welcome(w http.ResponseWriter, r *http.Request){
	cookieObtained, err:= r.Cookie("session_token")
	if err!=nil{
		if err == http.ErrNoCookie{
    		    w.WriteHeader(http.StatusUnauthorized)
		    fmt.Println("Cookie obtained in welcome",err)
		    return
	       }
	       
	        w.WriteHeader(http.StatusBadRequest)
		return
	}

	sessionToken:= cookieObtained.Value

	response, err:= redis_cache.Do("GET", sessionToken)
	if err !=nil{
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Error with GET redis response", err)
		return
	}
	if response == nil{
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.Write([]byte(fmt.Sprintf("Welcome %s", response)))
}


func refresh(w http.ResponseWriter, r *http.Request){
	getToken, err:= r.Cookie("session_token")
	if err!=nil{
		if err == http.ErrNoCookie{
                    w.WriteHeader(http.StatusUnauthorized)
                    fmt.Println("Refresh error",err)
                    return
               }

                w.WriteHeader(http.StatusBadRequest)
                return
	}
	sessionToken:= getToken.Value

	response, err:= redis_cache.Do("GET", sessionToken)
	if err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Error while getting response for Refresh",err)
		return
	}
	if response == nil{
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	newSessionToken:=uuid.New().String()
	_, err= redis_cache.Do("SETEX", newSessionToken, 120, fmt.Sprintf("%s", response))
	if err!= nil{
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Error while setting token in refresh",err)
		return
	}

	_, err = redis_cache.Do("DEL", sessionToken)
	if err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Error while deleting the token in refresh",err)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name: "session_token",
		Value: newSessionToken,
		Expires: time.Now().Add(120*time.Second),
	})
}

func main(){
	http.HandleFunc("/signin", signin)
	http.HandleFunc("/welcome", welcome)
	http.HandleFunc("/refresh", refresh)
	http.ListenAndServe(":8008", nil)
}
