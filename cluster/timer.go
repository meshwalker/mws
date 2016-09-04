package cluster

import (
	log "github.com/Sirupsen/logrus"
	gocache "github.com/patrickmn/go-cache"
	"time"
	"net"
)


type HealthCheckPool  struct {
	Cache	*gocache.Cache
}


type HealthCheckTimer struct {
	IP	net.IP
	Stop	chan bool
}


func (p *HealthCheckPool) Init(){
	p.Cache = gocache.New(gocache.NoExpiration, gocache.NoExpiration)
}

func (p *HealthCheckPool) AddTimer(cluster_id string, cIP net.IP) {
	stopChan := make(chan bool, 1)

	go func() {
		timer := time.NewTimer(time.Second*65)

		select {
		case <- timer.C:
			log.Info("Timer for Cluster expired: ",cluster_id)
			log.Info("Client IP: ",cIP)
		//setRedirect()
		case <- stopChan:
			timer.Stop()
		}
		close(stopChan)
	}()

	hcTimer := &HealthCheckTimer{
		IP:	cIP,
		Stop:	stopChan,
	}

	if err := p.Cache.Add(cluster_id, hcTimer, gocache.NoExpiration); err != nil {
		log.Error(err)
	}
	return
}


func (p *HealthCheckPool) RemoveTimer(cluster_id string) bool {
	x, found := p.Cache.Get(cluster_id)
	if found {
		hcTimer := x.(*HealthCheckTimer)
		hcTimer.Stop <- true

		p.Cache.Delete(cluster_id)
		return true
	}

	return false
}