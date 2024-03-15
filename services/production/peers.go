package production

import (
	"strconv"
	"time"

	"github.com/anacrolix/torrent"
)

type PeerScheduler struct {
	Count      int
	LastUpdate time.Time
	IsUpdating bool
}

var Peers = make(map[string]*PeerScheduler)

func DisplayCachedPeersOrSchedule(magnet string) string {
	scheduler := Peers[magnet]
	if scheduler == nil {
		Peers[magnet] = &PeerScheduler{}
		go UpdatePeers(magnet, Peers[magnet])
		return "scheduled"
	}

	if scheduler.IsUpdating {
		return "counting..."
	}

	if time.Since(scheduler.LastUpdate) > 1*time.Minute {
		go UpdatePeers(magnet, scheduler)
	}

	return strconv.Itoa(scheduler.Count)
}

func UpdatePeers(magnet string, scheduler *PeerScheduler) error {
	if scheduler.IsUpdating {
		return nil
	}

	scheduler.IsUpdating = true
	defer func() { scheduler.IsUpdating = false }()

	c, err := torrent.NewClient(torrent.NewDefaultClientConfig())
	if err != nil {
		return err
	}
	defer c.Close()

	t, err := c.AddMagnet(magnet)
	if err != nil {
		return err
	}

	<-t.GotInfo()
	scheduler.Count = len(t.KnownSwarm())
	scheduler.LastUpdate = time.Now()

	return nil
}
