package cluster

import (
	"time"
	"net"
	"meshwalker.com/mws/pkg/types"
	"errors"
)

type Cluster struct {
	Id		string	`json:"id,omitempty" gorethink:"id,omitempty"`
	Name		string	`json:"name,omitempty" gorethink:"name,omitempty"`
	BaseDomain	string	`json:"base_domain,omitempty" gorethink:"base_domain,omitempty"`
	FullDomain	string	`json:"full_domain,omitempty" gorethink:"full_domain,omitempty"`
	IP		net.IP	`json:"ip,omitempty" gorethink:"ip,omitempty"`
	Token		string	`json:"token,omitempty" gorethink:"token,omitempty"`
	CustomerId	string	`json:"customer_id,omitempty" gorethink:"customer_id,omitempty"`
	Nodes		[]Node	`json:"nodes" gorethink:"nodes"`
	CreatedAt	time.Time	`json:"created_at,omitempty" gorethink:"created_at,omitempty"`
	ModifiedAt	time.Time	`json:"modified_at,omitempty" gorethink:"modified_at,omitempty"`
}

type HealthCheckMsg struct {
	IP	string	`json:"ip,omitempty" `
	Token	string	`json:"token,omitempty"`
}

type HealthCheckConfirmMsg struct {
	Token		string	`json:"token,omitempty"`
}

type RestNewCluster struct {
	CustomerId	string	`json:"customer_id"`
	Name		string	`json:"name" gorethink:"name"`
	BaseDomain	string	`json:"base_domain" gorethink:"base_domain"`
}

type RestNewNode struct {
	UId		string	`json:"uid,omitempty" gorethink:"uid,omitempty"`
	Name	string	`json:"name" gorethink:"name"`
}

type Node struct {
	UId		string	`json:"uid,omitempty" gorethink:"uid,omitempty"`
	Name		string	`json:"name,omitempty" gorethink:"name,omitempty"`
	Token		string	`json:"token,omitempty" gorethink:"token,omitempty"`
	CreatedAt	time.Time	`json:"created_at,omitempty" gorethink:"created_at,omitempty"`
	ModifiedAt	time.Time	`json:"modified_at,omitempty" gorethink:"modified_at,omitempty"`
}


func (cluster *Cluster) AddNode(item Node)  {
	cluster.Nodes = append(cluster.Nodes, item)
}


func (c *Cluster) RemoveNode(node_id string) error {
	for i := 0; i < len(c.Nodes) ; i++ {
		if c.Nodes[i].UId == node_id {
			c.Nodes = append(c.Nodes[:i], c.Nodes[i+1:]...)
			/*copy(c.Nodes[i:], c.Nodes[i+1:])
			c.Nodes[len(c.Nodes)-1] = nil // or the zero value of T
			c.Nodes = c.Nodes[:len(c.Nodes)-1]*/
			return nil
		}
	}
	return errors.New("Provided node id not found")
}


func ErrMsgNoClusterId() *types.ErrMsg {
	return &types.ErrMsg{
		Status: "error",
		Message: "Missing cluster id",
	}
}
