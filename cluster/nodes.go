package cluster

import (
	r "gopkg.in/dancannon/gorethink.v2"
	db "meshwalker.com/mws/pkg/database"
	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"meshwalker.com/mws/pkg/types"
	"github.com/satori/go.uuid"
)

func AddNode(c *gin.Context) {
	cluster_id := c.Param("cid")
	var newNode RestNewNode
	var cluster Cluster
	c.Bind(&newNode)

	if cluster_id == "" {
		c.JSON(http.StatusBadRequest, ErrMsgNoClusterId())
		return
	}

	node := &Node{
		UId:	uuid.NewV4().String(),
		Name:	newNode.Name,
		CreatedAt:	time.Now(),
		ModifiedAt:	time.Now(),
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

	if err := cur.One(&cluster); err != nil {
		log.Error(err)
		c.Status(http.StatusNotFound)
		return
	}

	cluster.AddNode(*node)
	cursor, errReplace := r.DB(dbName).Table(tableName).Replace(cluster).RunWrite(db.Session)
	if errReplace != nil {
		log.Error(errReplace)
		c.Status(http.StatusInternalServerError)
		return
	}

	if cursor.Replaced != 1 {
		c.Status(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, &cluster)
	return
}


func RemoveNode(c *gin.Context) {
	cluster_id := c.Param("cid")
	node_id := c.Param("nid")
	var cluster Cluster

	if cluster_id == "" || node_id == "" {
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

	if err := cur.One(&cluster); err != nil {
		log.Error(err)
		c.Status(http.StatusNotFound)
		return
	}

	if err := cluster.RemoveNode(node_id); err != nil {
		c.JSON(http.StatusNotFound, &types.ErrMsg{
			Status: "error",
			Message: err.Error(),
		})
		return
	}

	cursor, errReplace := r.DB(dbName).Table(tableName).Replace(cluster).RunWrite(db.Session)
	if errReplace != nil {
		log.Error(errReplace)
		c.Status(http.StatusInternalServerError)
		return
	}

	if cursor.Replaced != 1 {
		c.Status(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, &cluster)
	return
}