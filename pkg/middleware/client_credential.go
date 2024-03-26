package middleware

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func ClientAuthenticated() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			clientID := r.Header.Get(HEADER_CLIENT_ID)
			if clientID == "" {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}
			signature := r.Header.Get(HEADER_CLIENT_SIGNATURE)
			if signature == "" {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}
			timestamp := r.Header.Get(HEADER_TIMESTAMP)
			if timestamp == "" {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}

			var requestBody bytes.Buffer
			_, err := requestBody.ReadFrom(r.Body)
			defer r.Body.Close()
			if err != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}
			bodyString := requestBody.String()
			minifiedReq := minifyRequestBody(bodyString)

			cache := GetCache(r.Context())
			clientSecret, err := cache.Get(r.Context(), fmt.Sprintf("client_credential:%s", clientID))
			if err != nil || clientSecret == "" {
				MakeLogEntry("ClientAuthenticated - GetClientSecret").Error(err)
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}

			validSignature := validateSignature(signature, clientID, clientSecret, minifiedReq, timestamp)
			if !validSignature {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}

			// Rewind the request body
			r.Body = ioutil.NopCloser(bytes.NewReader(requestBody.Bytes()))
			next.ServeHTTP(w, r)
		})
	}
}

func validateSignature(signature, clientID, clientSecret, body, createdTime string) bool {
	data := clientID + body + createdTime
	hmac := hmac.New(sha512.New, []byte(clientSecret))

	// compute the HMAC
	hmac.Write([]byte(data))
	dataHmac := hmac.Sum(nil)
	expectedSignature := hex.EncodeToString(dataHmac)

	// convert the timeStamp to a Unix timestamp
	currentTime := time.Now().Unix()
	timeStampInt, err := strconv.ParseInt(createdTime, 10, 64)
	if err != nil {
		MakeLogEntry("validateSignature - parseCreateTime").Error(err)
		return false
	}

	// validate expired signature
	if currentTime > timeStampInt+15 {
		MakeLogEntry("validateSignature - expire signature").Error(err)
		return false
	}
	//compare result
	return strings.EqualFold(expectedSignature, signature)
}

func minifyRequestBody(requestBody string) string {
	// Remove leading and trailing white spaces
	minifiedBody := strings.TrimSpace(requestBody)

	// Remove new lines and tabs
	minifiedBody = strings.ReplaceAll(minifiedBody, "\n", "")
	minifiedBody = strings.ReplaceAll(minifiedBody, "\t", "")

	// Remove unnecessary white spaces between characters
	minifiedBody = strings.Join(strings.Fields(minifiedBody), "")

	return minifiedBody
}
