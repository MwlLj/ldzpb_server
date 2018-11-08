package sessionmgr

import (
	// "strconv"
	"time"
)

var Memory_type_sqlite int = 1
var Memory_type_memory int = 2
var Memory_type_mysql int = 3

type CSessionMgr interface {
	ConnectDb()
	DisConnectDb()
	CreateSession(timeout uint64, userdata string) (sessionId string, err error)
	DestroySession(sessionId string) error
	SessionIsVaild(sessionId string) (isVaild bool, err error)
	ResetLosevaildTime(sessionId string) error
	DeleteSessionAfterLosevaild(nowTimeStamp uint64) error
	GetUserdata(sessionId string) (userdata string, err error)
}

func getNowTimeStamp() uint64 {
	t := time.Now()
	return uint64(t.UTC().UnixNano())
}

var endChannel chan (bool)

func New(memoryType int) CSessionMgr {
	var mgr CSessionMgr
	switch memoryType {
	case Memory_type_mysql:
		mgr = NewSessionMgrMysql()
		mgr.ConnectDb()
	}
	endChannel = make(chan (bool))
	go func() {
		for {
			select {
			case <-endChannel:
				return
			default:
				mgr.DeleteSessionAfterLosevaild(getNowTimeStamp())
				time.Sleep(time.Millisecond * 1000)
			}
		}
	}()
	return mgr
}

func Destroy(mgr CSessionMgr) {
	if mgr != nil {
		mgr.DisConnectDb()
		close(endChannel)
	}
}
