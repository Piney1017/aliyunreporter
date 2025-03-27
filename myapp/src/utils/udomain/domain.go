package udomain

import (
	"fmt"
	"myapp/pkg/constant"
	"myapp/pkg/types"
	"myapp/src/utils/uredis"
	"myapp/src/utils/utime"
	"strings"
	"time"
)

func DomainNormal(domain string) string {

	d := strings.ToLower(domain)

	parts := strings.Split(d, ".")

	lparts := len(parts)
	if lparts <= 2 {
		return d
	}

	return strings.Join(parts[lparts-2:], ".")
}

func keyStatDomain(statAt int64, stats types.DomainState) string {
	return fmt.Sprintf("%v:%v", statAt, stats)
}

func StatDomain(domain string, stats types.DomainState) {
	dn := DomainNormal(domain)
	key := keyStatDomain(utime.StatAt(), stats)
	n, _ := uredis.HIncr(key, dn, 1)
	if n < 2 {
		uredis.Expire(key, constant.STAT_DOMAIN_HASH_KEEP_DURANG*time.Second)
	}
}

func genKeyByStats(stats types.DomainState) []string {
	keys := make([]string, 0, 16)
	statAt := utime.StatAt()

	for i := int64(0); i < 16; i++ {
		keys = append(keys,
			keyStatDomain(statAt-i, stats),
		)
	}

	return keys
}
