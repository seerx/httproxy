package block

import (
	"log"
	"os/exec"
)

// RejectIP 阻止 IP 地址
func RejectIP(IP string) bool {
	log.Println("禁止 IP :" + IP)
	out, err := exec.Command("bash",
		"/usr/local/proxy/deny.sh",
		IP,
	).Output()
	if err != nil {
		log.Println("禁止 IP 时发生错误:" + err.Error())
	}
	log.Println(string(out))
	// reject := fmt.Sprintf("/usr/bin/firewall-cmd --permanent --add-rich-rule='rule family=\"ipv4\" source address=\"%s\" port protocol=\"tcp\" port=\"80\" reject'", IP)
	// log.Println(reject)
	// reject = fmt.Sprintf("--add-rich-rule='rule family=\"ipv4\" source address=\"%s\" port protocol=\"tcp\" port=\"80\"", IP)
	// out, err := exec.Command("python3",
	// 	"/usr/bin/firewall-cmd",
	// 	"--permanent",
	// 	reject,
	// 	"reject").Output()
	// if err != nil {
	// 	log.Println("禁止 IP 时发生错误:" + err.Error())
	// 	return false
	// }
	// log.Println(string(out))

	// reload := "/usr/bin/firewall-cmd --reload"
	// log.Println(reload)
	// out, err = exec.Command("python3", "/usr/bin/firewall-cmd", "--reload").Output()
	// if err != nil {
	// 	log.Println("重新加载防火墙时发生错误:" + err.Error())
	// 	return false
	// }
	// log.Println(string(out))
	return true
}
