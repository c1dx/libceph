package conn

import (
	"github.com/ceph/go-ceph/rados"
)

// CephConn ...
type CephConn struct {
	Rados *rados.Conn
}

// GetCephConn ...
func GetCephConn(user string) (*CephConn, error) {
	cephConn := new(CephConn)
	if err := cephConn.Connect(user); err != nil {
		return nil, err
	}
	return cephConn, nil
}

// Connect ...
func (cephCli *CephConn) Connect(user string) error {
	conn, err := rados.NewConnWithUser(user)
	if err != nil {
		return err
	}

	err = conn.ReadDefaultConfigFile()
	if err != nil {
		return err
	}
	// args := []string{ "--mon-host", "10.11.12.196" }
	// conn.ParseCmdLineArgs(args)
	err = conn.Connect()
	if err != nil {
		return err
	}

	cephCli.Rados = conn

	return nil
}
