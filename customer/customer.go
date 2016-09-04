package customer

import (
	"net/http"
	"github.com/gin-gonic/gin"
	r "gopkg.in/dancannon/gorethink.v2"
	db "meshwalker.com/mws/pkg/database"
	log "github.com/Sirupsen/logrus"
	"meshwalker.com/mws/pkg/types"
	//msg "meshwalker.com/mws/pkg/messages"
	"time"
)

const (
	dbName		= "mws"
	tableName	= "customer"
)


func GetById(c *gin.Context) {
	customer_id := c.Param("id")
	var result Customer

	cur, err := r.DB(dbName).Table(tableName).Get(customer_id).Run(db.Session)
	defer cur.Close()
	if err != nil {
		log.Error(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	if cur.IsNil() {
		c.Status(http.StatusNotFound)
		return
	}

	if err := cur.One(&result); err != nil {
		log.Error(err)
		c.Status(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, &result)
	return
}


// Create a customer
func Create(c *gin.Context) {
	var newCustomer	RestCustomer
	if err := c.BindJSON(&newCustomer); err != nil {
		log.Error(err)
		c.Status(http.StatusBadRequest)
		return
	}

	if newCustomer.Password != newCustomer.ConfirmPassword {
		msg := &types.ErrMsg{
			Status:	"error",
			Message: "Entered passwords do not match",
		}
		c.JSON(http.StatusBadRequest, msg)
		return
	}

	if newCustomer.MobileNumber == "" {
		msg := &types.ErrMsg{
			Status:	"error",
			Message: "You have to provide a valid mobile phone number!",
		}
		c.JSON(http.StatusBadRequest, msg)
		return
	}

	// Check if user already exists
	if cur, err := r.DB(dbName).Table(tableName).Filter(
		r.Row.Field("mobile_number").Eq(newCustomer.MobileNumber)).Run(db.Session); err != nil {
		log.Error(err)
		defer cur.Close()
		return
	} else {
		defer cur.Close()
		var res []interface{}

		cur.All(&res)
		if(len(res) != 0) {
			log.Error("This user already existes!")
			c.JSON(http.StatusBadRequest, &types.ErrMsg{
				Status: "error",
				Message: "The provided mail address is already taken",
			})
			return
		} else {
			log.Info("Cool, a new user!")
		}
	}

	// Generate a salted password
	saltedPassword, salt := GenSaltedPassword(newCustomer.Password)

	// Map RestCustomer to database Customer object
	customer := &Customer{
		FirstName:	newCustomer.FirstName,
		LastName:	newCustomer.LastName,
		Address:	newCustomer.Address,
		ZIP:		newCustomer.ZIP,
		City:		newCustomer.City,
		Country:	newCustomer.Country,
		Email:		newCustomer.Email,
		MobileNumber:	newCustomer.MobileNumber,
		Password:	saltedPassword,
		PasswordSalt:	salt,
		CreatedAt:	time.Now(),
		ModifiedAt:	time.Now(),
	}

	// Create customer on database
	cur, err := r.DB(dbName).Table(tableName).Insert(customer).RunWrite(db.Session)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Creating user failed",
		})
		return
	}

	c.String(http.StatusOK, cur.GeneratedKeys[0])
	return
	/*
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": message,
		"nick":    nick,
	})*/
}

func Update(c *gin.Context) {

}


// Delete customer
func Delete(c *gin.Context) {
	customer_id := c.Param("id")

	cur, err := r.DB(dbName).Table(tableName).Get(customer_id).Delete().RunWrite(db.Session)
	if err != nil {
		log.Error(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	if cur.Deleted != 1 {
		c.Status(http.StatusNotFound)
	}

	c.Status(http.StatusOK)
	return
}
