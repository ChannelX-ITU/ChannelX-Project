package channel

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/user"
	"path/filepath"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"encoding/base64"
	"strings"
)

// getClient uses a Context and Config to retrieve a Token
// then generate a Client. It returns the generated Client.
func getClient(ctx context.Context, config *oauth2.Config) *http.Client {
	cacheFile, err := tokenCacheFile()
	if err != nil {
		log.Fatalf("Unable to get path to cached credential file. %v", err)
	}
	tok, err := tokenFromFile(cacheFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(cacheFile, tok)
	}
	return config.Client(ctx, tok)
}

// getTokenFromWeb uses Config to request a Token.
// It returns the retrieved Token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var code string
	if _, err := fmt.Scan(&code); err != nil {
		log.Fatalf("Unable to read authorization code %v", err)
	}

	tok, err := config.Exchange(oauth2.NoContext, code)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web %v", err)
	}
	return tok
}

// tokenCacheFile generates credential file path/filename.
// It returns the generated credential path/filename.
func tokenCacheFile() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	tokenCacheDir := filepath.Join(usr.HomeDir, ".credentials")
	os.MkdirAll(tokenCacheDir, 0700)
	return filepath.Join(tokenCacheDir,
		url.QueryEscape("gmail-go-quickstart.json")), err
}

// tokenFromFile retrieves a Token from a given file path.
// It returns the retrieved Token and any read error encountered.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	t := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(t)
	defer f.Close()
	return t, err
}

// saveToken uses a file path to create a file and store the
// token in it.
func saveToken(file string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", file)
	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func GetGmailService() *gmail.Service {
	ctx := context.Background()

	s := Secret{Installed:Cred{ClientId:"121965707693-tp5icl8v8vsk991didual14djb7d30qo.apps.googleusercontent.com", ProjectId:"turnkey-timer-189309", AuthProvi:"https://www.googleapis.com/oauth2/v1/certs", AuthUri:"https://accounts.google.com/o/oauth2/auth", TokenUri:"https://accounts.google.com/o/oauth2/token", ClientSecre:"8MSSa87g1NhmnmNh1tMyhjox", Redirect:[]string{"urn:ietf:wg:oauth:2.0:oob", "http://localhost"}}}
	b, _ := json.Marshal(s)
	// If modifying these scopes, delete your previously saved credentials
	// at ~/.credentials/gmail-go-quickstart.json
	config, err := google.ConfigFromJSON(b, gmail.GmailModifyScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(ctx, config)

	srv, err := gmail.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve gmail Client %v", err)
	}

	user := "me"
	_, err = srv.Users.Labels.List(user).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve labels. %v", err)
	}

	return srv
}

type Secret struct {
	Installed	Cred	`json:"installed"`
}

type Cred struct {
	ClientId	string 	`json:"client_id"`
	ProjectId	string	`json:"project_id"`
	AuthUri		string	`json:"auth_uri"`
	TokenUri	string	`json:"token_uri"`
	AuthProvi	string	`json:"auth_provider_x509_cert_url"`
	ClientSecre	string	`json:"client_secret"`
	Redirect	[]string	`json:"redirect_uris"`
}

func (s *Server) SendMail(to string, subject string, msg string) {
	data := []byte(msg)
	str := base64.URLEncoding.EncodeToString(data)
	message := gmail.Message{Raw: str, Payload:&gmail.MessagePart{Headers:[]*gmail.MessagePartHeader{{Name:"to", Value:to}, {Name:"subject", Value:subject}}, Body:&gmail.MessagePartBody{Data:str}}}
	temp := []byte("From: 'me'\r\n" +
		"To:  " + to +"\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" + msg)
	message.Raw = base64.StdEncoding.EncodeToString(temp)
	message.Raw = strings.Replace(message.Raw, "/", "_", -1)
	message.Raw = strings.Replace(message.Raw, "+", "-", -1)
	message.Raw = strings.Replace(message.Raw, "=", "", -1)

	s.mailMan.Users.Messages.Send("me", &message).Do()
}
