package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
)


func GenerateRsa() (string,error){
    privatekey, err := rsa.GenerateKey(rand.Reader, 2048)
    if err != nil {
        return "",err
    
    }
    publickey := &privatekey.PublicKey

    // dump public key to file
    publicKeyBytes, err := x509.MarshalPKIXPublicKey(publickey)
    if err != nil {
        return "",err
        
    }
    json,err:= json.Marshal(publicKeyBytes)
    if err != nil {
        return "",err
    }
    st:=string(json)
    return st, nil
}