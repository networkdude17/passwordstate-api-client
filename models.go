// Data Models

package passwordstateclient

// Model - Password
type Password struct {
	PasswordID  int    `json:"PasswordID"`
	UserName 	string `json:"UserName"`
	Password    string `json:"Password"`
}
