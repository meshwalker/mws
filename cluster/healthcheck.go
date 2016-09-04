package cluster

import (
	"net/http"
	"github.com/gin-gonic/gin"
	r "gopkg.in/dancannon/gorethink.v2"
	db "meshwalker.com/mws/pkg/database"
	log "github.com/Sirupsen/logrus"
	"meshwalker.com/mws/pkg/types"
	"net"
)

var pool *HealthCheckPool

func init() {
	pool = &HealthCheckPool{}
	pool.Init()
}


func GetMyIP(c *gin.Context) {
	clientip := c.ClientIP()
	c.JSON(http.StatusOK, gin.H{"ip": clientip})
}


func HealthCheck(c *gin.Context) {
	cluster_id := c.Param("cid")
	if cluster_id == "" {
		c.JSON(http.StatusBadRequest, &types.ErrMsg{
			Status: "error",
			Message: "Missing cluster id",
		})
		return
	}

	var hcMsg HealthCheckMsg
	c.Bind(&hcMsg)
	cIP := net.ParseIP(hcMsg.IP)

	// If cluster specific timer is in cache (true), remove it and restart the timer
	if resetTimer(cluster_id, cIP) {
		c.Status(http.StatusOK)
		return
	}

	// Compare client suggestion with ip of request
	if !cIP.Equal(net.ParseIP(c.ClientIP())) {
		c.JSON(http.StatusBadRequest, &types.ErrMsg{
			Status: "error",
			Message: "Detecting cluster ip failed",
		})
		return
	}

	// Verify new meshwalker cluster ip
	if !verifyMeshwalkerCluster(cIP, hcMsg.Token) {
		log.Error("Health check of cluster "+cluster_id+" failed")
		return
	}

	if !updateDBClusterIP(c, cluster_id, cIP) {
		return
	}

	// Start HealthCheckTimer
	pool.AddTimer(cluster_id, cIP)

	return
}


// Check if stored ip address is up to date
func resetTimer(cluster_id string, cIP net.IP) bool {
	x, found := pool.Cache.Get(cluster_id)
	if found {
		timer := x.(*HealthCheckTimer)
		if timer.IP.Equal(cIP) {
			timer.Stop <- true
			pool.AddTimer(cluster_id, cIP)
			return true
		}
	}

	return false
}


// If the server received a new ip address of a meshwalker cluster
// the tcp should verify if the cluster is really accessible via this ip
func verifyMeshwalkerCluster(cIP net.IP, token string) bool {
	client := &http.Client{}

	req, err := http.NewRequest("POST", "https://"+cIP.String(), nil)
	req.PostForm.Add("token", token)
	resp, err := client.Do(req)
	if err != nil {
		log.Error(err)
		return false
	}
	defer resp.Body.Close()
	return true
}


// Update IP of cluster
func updateDBClusterIP(c *gin.Context, cluster_id string, cIP net.IP) bool {
	var cluster Cluster
	cur, err := r.DB(dbName).Table(tableName).Get(cluster_id).Run(db.Session)
	if err != nil {
		log.Error(err)
		c.Status(http.StatusInternalServerError)
		return false
	}
	defer cur.Close()

	cur.One(&cluster)
	cluster.IP = cIP

	if _, err := r.DB(dbName).Table(tableName).Insert(cluster).RunWrite(db.Session); err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Storing cluster ip failed",
		})
		return false
	}

	return true
}