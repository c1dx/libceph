package client

import (
	"libceph/actions"
	"libceph/common"
)

var cephClient *CephClient

// GetCephConn ...
func GetCephClient() (*CephClient, error) {
	if cephClient == nil {
		cephClient = new(CephClient)
		if err := cephClient.Initial(); err != nil {
			return nil, err
		}
	}
	return cephClient, nil
}

type CephClient struct {
	Cluster *actions.Cluster
	Pool    *actions.Pool
	Mon     *actions.Mon
	Osd     *actions.Osd
	Pg      *actions.Pg
}

func (this *CephClient) Initial() error {
	cephConn, err := common.GetCephConn()
	if err != nil {
		return err
	}
	this.Cluster = &actions.Cluster{
		ActionBase: actions.ActionBase{
			CephConn: cephConn,
		},
	}
	this.Pool = &actions.Pool{
		ActionBase: actions.ActionBase{
			CephConn: cephConn,
		},
	}
	this.Mon = &actions.Mon{
		ActionBase: actions.ActionBase{
			CephConn: cephConn,
		},
	}
	this.Osd = &actions.Osd{
		ActionBase: actions.ActionBase{
			CephConn: cephConn,
		},
	}
	this.Pg = &actions.Pg{
		ActionBase: actions.ActionBase{
			CephConn: cephConn,
		},
	}
	return nil
}
