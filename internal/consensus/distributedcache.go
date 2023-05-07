package consensus

import "log"

func (d *DistributedCache) Join(name, addr string) error {
	//TODO implement me
	log.Printf("%s joined at %s", name, addr)
	return nil
}

func (d *DistributedCache) Leave(name string) error {
	//TODO implement me
	log.Printf("%s left", name)
	return nil
}
