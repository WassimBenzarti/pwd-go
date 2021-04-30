package api

import (
	"testing"

	. "github.com/stretchr/testify/assert"
	"github.com/wassimbenzarti/pwd-go/models"
)

const COOKIE = "_biz_uid=b63210ac5bb948caa2798b01bf044443; ajs_anonymous_id=%2234e65bb7-e690-4eb8-b94f-368eea918af5%22; _mkto_trk=id:929-FJL-178&token:_mch-play-with-docker.com-1619110250665-61259; _biz_flagsA=%7B%22Version%22%3A1%2C%22ViewThrough%22%3A%221%22%2C%22XDomain%22%3A%221%22%7D; ajs_user_id=%2289bfdfaa-9c15-45e7-9796-f76fc9f5cd00%22; _ga=GA1.2.39277319.1619110269; _biz_sid=1d1f57; id=MTYxOTYyMzU5NHxQaVoyM3NTV1NMZTlvODE3dW40ZHAyV050YzdMTndJUVNPVzZkRE5RMUo2ZjAxbzJ5dlZmN1h1eUNqTVNmaE9KVll0ZUFnSmF1TUR3djRqRHJOejlDSFpfcVZZZElJc1JEcmtadmxvU2hyOVpzd01rV3U2UlJzbGJ3WFJRZmtSeHdjRVFPX0pvc3VaekNMR2Rmdk85czZPUUJsN2cwYzRGaXdHRGxDNVJjaUFObHkteE16dFhOVzQtMVVqeEw3OHU5QnV4TzYwM0JZQTZsZW1aUHRfTVlNeFFKeks0emRfb0RCT3BySS1oTGQ0Q0tpNHhUZ2YyV0ctNUZSUjBNeXN4c0M0RzhiaGVlX1RIamFyMDJjcTRSa1lFSUwwPXy2oMcOkbQdmZyrZ7B12rL5-UiB-ufANpWTeZSqi-kClw==; _gid=GA1.2.858234351.1619623596; _biz_nA=55; _biz_pendingA=%5B%22m%2Fipv%3F_biz_r%3D%26_biz_h%3D802059049%26_biz_u%3Db63210ac5bb948caa2798b01bf044443%26_biz_s%3D1d1f57%26_biz_l%3Dhttps%253A%252F%252Flabs.play-with-docker.com%252F%26_biz_t%3D1619624740836%26_biz_i%3DPlay%2520with%2520Docker%26_biz_n%3D54%26rnd%3D345399%22%5D"

var id = ""

func TestSession(t *testing.T) {
	id := CreateSession(COOKIE)

	IsType(t, "", id)
}

func TestGetSession(t *testing.T) {
	id := CreateSession(COOKIE)
	session := GetSession(id)
	IsType(t, models.Session{}, session)
	NotEmpty(t, session.Host)
	NotEmpty(t, session.Id)
	NotEmpty(t, session.PwdIPAddress)
	Equal(t, true, session.Ready)
}
