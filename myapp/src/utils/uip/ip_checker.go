package uip

import (
	"fmt"
	"time"

	"myapp/pkg/constant"
	"myapp/src/utils/uredis"
)

const (
	ipCheckKey = "for_set_expirs"
)

func keyCheckIP() string {
	nowUnix := time.Now().Unix()
	return fmt.Sprintf("check_ip_dupi:%v", nowUnix/constant.STAT_SESSION_DURANG_SEC)
}

func Init() {
	go func() {
		key := keyCheckIP()
		if checkIP(key, ipCheckKey) {
			uredis.Expire(key, 72*time.Hour)
		}

		t := time.NewTicker(30 * time.Second)
		for range t.C {
			key := keyCheckIP()
			if checkIP(key, ipCheckKey) {
				uredis.Expire(key, 72*time.Hour)
			}
		}
	}()
}

func checkIP(key, ip string) bool {
	i, e := uredis.AddSet(key, ip)
	if e != nil {
		return true
	}

	return i > 0
}

func CheckIP(ip string) bool {
	return checkIP(keyCheckIP(), ip)
}
