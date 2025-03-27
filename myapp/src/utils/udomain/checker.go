package udomain

import (
	"context"
	"fmt"
	"myapp/pkg/constant"
	"myapp/pkg/types"
	"myapp/src/utils/tgbot"
	"myapp/src/utils/uredis"
	"os"
	"strings"
	"time"

	"github.com/samber/lo"
	"github.com/sirupsen/logrus"
)

var (
	calcCh        = make(chan struct{}, 100)
	sendToCsAlert = false
	banedDomains  = make(map[string]bool)
)

func init() {
	sendToCsAlert = os.Getenv(constant.SEND_TO_CS_ALERT) == "true"

	banedDomains["tjzqit.com"] = true
}

func sendState(states map[types.DomainState]map[string]int64) {

	hostState := make(map[string][]int64)

	for s, domains := range states {
		for d, cnt := range domains {
			if _, found := hostState[d]; !found {
				hostState[d] = make([]int64, types.DomainStateMax)
			}
			hostState[d][int(s)] = cnt
		}
	}

	out := make([]string, 0, 48)
	out = append(out, "域名出问题啦, 快让技术查看!")

	allStates := lo.RangeFrom[types.DomainState](0, int(types.DomainStateMax))

	for d, ss := range hostState {
		out = append(out, fmt.Sprintf("%v:", d))
		lo.ForEach(allStates, func(s types.DomainState, _ int) {
			if ss[s] > 0 {
				out = append(out, fmt.Sprintf("\t%v: %v", s, ss[s]))
			}
		})
	}

	if len(out) <= 1 {
		return
	}

	text := strings.Join(out, "\n")

	logrus.WithContext(context.Background()).
		WithField("hostState", hostState).
		Debugln(text)

	if sendToCsAlert {
		tgbot.SendAlert(text)
	}

	tgbot.SendTechAlert(text)
}

func CheckDomainStat() {

	alertDomains := make(map[types.DomainState]map[string]int64)
	oldAlertDomains := make(map[types.DomainState]map[string]int64)

	doCheck := func(state types.DomainState, val int64) {
		keys := genKeyByStats(state)
		keyLen := len(keys)

		vals := make(map[string]int64)
		oldVals := make(map[string]int64)
		newVals := make(map[string]int64)
		for i, k := range keys {
			m, _ := uredis.HGetAll(k)

			for kk, vv := range m {
				if banedDomains[kk] {
					continue
				}
				vals[kk] += vv
			}

			if i > 0 {
				for kk, vv := range m {
					if banedDomains[kk] {
						continue
					}
					newVals[kk] += vv
				}
			}

			if i < keyLen-2 {
				for kk, vv := range m {
					if banedDomains[kk] {
						continue
					}
					oldVals[kk] += vv
				}
			}
		}

		for k, v := range newVals {
			if v >= val {
				if _, found := oldVals[k]; !found {
					if _, found := alertDomains[state]; !found {
						alertDomains[state] = make(map[string]int64)
					}
					alertDomains[state][k] = v
				}
				oldVals[k] = v
			} else {
				delete(oldVals, k)
			}
		}

		deletedKeys := make([]string, 0, 10)
		for k, v := range oldVals {
			if v < val {
				deletedKeys = append(deletedKeys, k)
			}
		}
		lo.ForEach(deletedKeys, func(k string, _ int) {
			delete(oldVals, k)
		})

		if len(oldVals) > 0 {
			oldAlertDomains[state] = oldVals
		}

		// logrus.WithContext(context.Background()).
		// 	WithFields(logrus.Fields{
		// 		"state":    state,
		// 		"vals":     vals,
		// 		"new_vals": newVals,
		// 		"old_vals": oldVals,
		// 		"alert":    alertDomains,
		// 		"oldAlert": oldAlertDomains,
		// 	}).
		// 	Debugln("newdocheckernewdocheckernewdochecker")
	}

	doCheck(types.DomainState404, 1)
	doCheck(types.DomainStateContent, 1)
	doCheck(types.DomainStateNetworkException, 100)
	doCheck(types.DomainStateTimeout, 150)
	doCheck(types.DomainStateNotAll, 300)

	// logrus.WithContext(context.Background()).
	// 	WithFields(logrus.Fields{
	// 		"alert":    alertDomains,
	// 		"oldAlert": oldAlertDomains,
	// 	}).
	// 	Debugln("sendState(alertDomains)")

	// send
	now := time.Now().Unix()

	if len(alertDomains) > 0 {
		sendState(alertDomains)
		uredis.Set("send_at", now, 24*time.Hour)
	}

	lastSentAt := uredis.GetInt64("send_at")
	if len(oldAlertDomains) > 0 && now-lastSentAt >= constant.SENT_ALERT_INTERVAL_SEC {
		sendState(oldAlertDomains)
		uredis.Set("send_at", now, 24*time.Hour)
	}
}
