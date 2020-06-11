package block

import (
	"fmt"
	"os/exec"

	"log"
)

// RejectIP 阻止 IP 地址
func RejectIP(IP string) bool {
	reject := fmt.Sprintf(`firewall-cmd --permanent --add-rich-rule="rule family="ipv4" source address="%s" port protocol="tcp" port="80" reject"`, IP)
	log.Println(reject)
	out, err := exec.Command(reject).Output()
	if err != nil {
		log.Println("禁止 IP 时发生错误:" + err.Error())
		return false
	}
	log.Println(string(out))

	reload := "firewall-cmd --reload"
	log.Println(reload)
	out, err = exec.Command(reload).Output()
	if err != nil {
		log.Println("重新加载防火墙时发生错误:" + err.Error())
		return false
	}
	log.Println(string(out))
	return true
}
