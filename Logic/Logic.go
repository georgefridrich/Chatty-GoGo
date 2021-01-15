package Logic 

import (
	"fmt" // package that implements formatted I/O
	"net/http"
	"github.com/gin-gonic/gin"  // GinGonic Framework
)

// This is not ideal but necessary, we'll hold contents in
// memory until we move to using a database

type userdata struct {
	Username string `json:"Username"`
	Subject string `json:"Subject"`
	Message string `json:"Message"`
	Messagefrom string `json"Messagefrom"`
}

type allUserdata []userdata

var saveduserdata = allUserdata{
	{
		Username: "super.mario",
		Subject: "lunch?",
		Message: "in the mood for BBQ today?",
		Messagefrom: "bob.ross",
	},
	{
		Username: "bob.ross",
		Subject: "",
		Message: "",
		Messagefrom: "",
	},
}

// We'll be calling these from Routes, create the functions
// to be exportable

func HomeLink(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	// Print to consumer
	c.JSON(http.StatusOK, gin.H {"message":"Welcome to Chatty-GoGo!"}) 
	// Print to console - not really needed with GinGonic 
	// but for now we'll leave it
	fmt.Println("Endpoint: Home Page") 
}

func UserRegistration(c *gin.Context) {
	var newUser userdata
	var err error

	c.BindJSON(&newUser)
    if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.Header("Content-Type", "application/json")
		//c.JSON(http.StatusOK, newUser)
		saveduserdata = append(saveduserdata, newUser)
		c.JSON(http.StatusOK, newUser)
	}
	fmt.Println("Endpoint: UserRegistration")
}

func ListUsers(c *gin.Context) {
	var err error
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, saveduserdata)
	}
	fmt.Println("Endpoint: ListUsers")
}

func RemoveUser(c *gin.Context) {
	//var err error
	usernames := c.Params.ByName("username")

	for i, individualUserdata := range saveduserdata {
		if individualUserdata.Username == usernames {
			saveduserdata = append(saveduserdata[:i], saveduserdata[i+1:]...)
            c.JSON(http.StatusOK, gin.H {"username: " + usernames: "has been successfully removed."})
		}
	}
	fmt.Println("Endpoint: RemoveUser")
}

func UpdateMessageforUser(c *gin.Context) {

	usernames := c.Params.ByName("username")
	c.Header("Content-Type", "application/json")

	var updatedUserdata userdata
	//var err error  *add error checking later

	c.BindJSON(&updatedUserdata)

    for i, individualUserdata := range saveduserdata {
		if individualUserdata.Username == usernames {
			individualUserdata.Subject = updatedUserdata.Subject
			individualUserdata.Message = updatedUserdata.Message
			individualUserdata.Messagefrom = updatedUserdata.Messagefrom
			saveduserdata = append(saveduserdata[:i], individualUserdata)
			c.JSON(http.StatusOK, gin.H {"Subject,Message,Messagefrom for: " + usernames: "successfully updated"})
		}
	}
    fmt.Println("Endpoint: UpdateMessage")
}