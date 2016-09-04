package cluster

import (
	"net/http"
	"github.com/gin-gonic/gin"
	r "gopkg.in/dancannon/gorethink.v2"
	db "meshwalker.com/mws/pkg/database"
	log "github.com/Sirupsen/logrus"
	"meshwalker.com/mws/pkg/types"
	"time"
)

const (
	dbName		= "mws"
	tableName	= "cluster"
)


func GetById(c *gin.Context) {
	cluster_id := c.Param("cid")
	var result Cluster

	if cluster_id == "" {
		c.JSON(http.StatusBadRequest, ErrMsgNoClusterId())
		return
	}

	cur, err := r.DB(dbName).Table(tableName).Get(cluster_id).Run(db.Session)
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


func Create(c *gin.Context) {
	var newCluster RestNewCluster
	if err := c.BindJSON(&newCluster); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	// Domain field is empty
	if newCluster.Name == "" || newCluster.BaseDomain == "" {
		msg := &types.ErrMsg{
			Status: "error",
			Message: "You  have to provid a valid domain as a cluster name",
		}
		log.Error(msg)
		c.JSON(http.StatusBadRequest, msg)
		return
	}

	fullDomain := newCluster.Name+"."+newCluster.BaseDomain

	if err := clusterExists(c, &newCluster, fullDomain); err != nil  {
		log.Error(err)
		return
	}


	cluster := &Cluster{
		CustomerId:	newCluster.CustomerId,
		Name:		newCluster.Name,
		BaseDomain:	newCluster.BaseDomain,
		FullDomain:	fullDomain,
		CreatedAt: 	time.Now(),
		ModifiedAt: 	time.Now(),
	}

	// Store new cluster in database
	cur, err := r.DB(dbName).Table(tableName).Insert(cluster).RunWrite(db.Session)
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
}


func Update(c *gin.Context) {

}


// Remove cluster from database
func Delete(c *gin.Context) {
	cluster_id := c.Param("cid")

	if cluster_id == "" {
		c.JSON(http.StatusBadRequest, ErrMsgNoClusterId())
		return
	}

	cur, err := r.DB(dbName).Table(tableName).Get(cluster_id).Delete().RunWrite(db.Session)
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


func clusterExists(c *gin.Context, nCluster *RestNewCluster, fullDomain string) error {
	// Check if cluster already exists
	if cur, err := r.DB(dbName).Table(tableName).Filter(
		r.Row.Field("full_domain").Eq(fullDomain)).Run(db.Session); err != nil {
		defer cur.Close()
		return err
	} else {
		defer cur.Close()
		var res []interface{}

		cur.All(&res)
		if(len(res) != 0) {
			log.Error("This cluster already existes!")
			c.JSON(http.StatusBadRequest, &types.ErrMsg{
				Status: "error",
				Message: "A cluster with the provided domainis already exists",
			})
			return nil
		} else {
			log.Info("Cool, a new cluster!")
			return nil
		}
	}

	// Check if user has already an associated cluster
	if cur, err := r.DB(dbName).Table(tableName).Filter(
		r.Row.Field("customer_id").Eq(nCluster.CustomerId)).Run(db.Session); err != nil {
		defer cur.Close()
		return err
	} else {
		defer cur.Close()
		var res []interface{}

		cur.All(&res)
		if(len(res) != 0) {
			log.Error("This has already an associated cluster")
			c.JSON(http.StatusBadRequest, &types.ErrMsg{
				Status: "error",
				Message: "This user is already owner of a cluster",
			})
			return nil
		} else {
			log.Info("Cool, a new cluster!")
			return nil
		}
	}
}