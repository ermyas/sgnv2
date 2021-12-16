package utils

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/celer-network/goutils/log"
	"github.com/spf13/viper"
)

func SendWithdrawAlert(addr, withdraw, deposit, delta, token string) {
	msg := fmt.Sprintf("withdraw refused, total withdraw more than %s of total deposit, token:`%s`, usr_addr: `%s`, has deposited:`%s`, has withdrawn: `%s`, want to withdraw:`%s`", "105%", token, addr, deposit, withdraw, delta)
	if viper.GetString("env") == "prod" {
		sendSlackP1AlertProd("Withdraw 105% Deposit Alert", msg)
	} else if viper.GetString("env") == "test" {
		sendSlackP1AlertTest("Withdraw 105% Deposit Alert", msg)
	}
	log.Warnf(msg)
}

type BalanceAlert struct {
	Token    string
	Balance  string
	Addr     string
	Withdraw string
	Deposit  string
}

type StatusAlertInfo struct {
	ChainId  uint64
	TxHash   string
	Ut       time.Time
	IsPegged bool
}

func SendBalanceAlert(alerts []*BalanceAlert) {
	msg := "find abnormal lp balance, total_deposit * 1.05 < (total_withdrawal + current_lp_balance):\n"
	for _, alert := range alerts {
		msg = msg + fmt.Sprintf("token:`%s`, balance: `%s`, usr_addr: `%s`, total withdraw: `%s`, total deposit:`%s` \n", alert.Token, alert.Balance, alert.Addr, alert.Withdraw, alert.Deposit)
	}
	log.Warnf(msg)
	if viper.GetString("env") == "prod" {
		sendSlackP1AlertProd("Abnormal LP Balance Alert", msg)
	} else if viper.GetString("env") == "test" {
		sendSlackP1AlertTest("Abnormal LP Balance Alert", msg)
	}
}

func SendProblematicBlockNumberAlert(addrs []string) {
	msg := "those addresses are left behind at reporting their chain's current block number.\n"
	for _, addr := range addrs {
		msg = msg + addr + "\n"
	}
	log.Warnf(msg)
	if viper.GetString("env") == "prod" {
		sendSlackStatusAlertProd("current block number behind alert", msg)
	}
}

func SendStatusAlert(alerts []*StatusAlertInfo, key string) {
	msg := "find abnormal status, " + key + ":\n"
	for _, alert := range alerts {
		msg = msg + fmt.Sprintf("chainId:`%d`, txHash:`%s`, updateTime:`%s`", alert.ChainId, alert.TxHash, alert.Ut)
		if alert.IsPegged {
			msg = msg + ", note: this is pegged"
		}
		msg = msg + " \n"
	}
	log.Warnf(msg)
	if viper.GetString("env") == "prod" {
		sendSlackStatusAlertProd(key, msg)
	} else if viper.GetString("env") == "test" {
		sendSlackP1AlertTest(key, msg)
	}
}

func sendSlackP1AlertTest(title string, msg string) {
	url := "https://hooks.slack.com/services/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"
	body := `{
		"channel": "#cbridge-v2-alert-test",
			"username": "%s",
			"text": "%s",
			"icon_emoji": "https://svblockchain.slack.com/services/BRARCSVU3?settings=1"
	}`

	err := httpDoPost(url, fmt.Sprintf(body, title, msg))
	if err != nil {
		log.Error("send alert error:", err)
	}
}

func sendSlackP1AlertProd(title string, msg string) {
	url := "https://hooks.slack.com/services/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"
	body := `{
		"channel": "#cbridge-v2-prod2-alert-p1",
			"username": "%s",
			"text": "%s",
			"icon_emoji": "https://svblockchain.slack.com/services/BRARCSVU3?settings=1"
	}`

	err := httpDoPost(url, fmt.Sprintf(body, title, msg))
	if err != nil {
		log.Error("send alert error:", err)
	}
}

func sendSlackStatusAlertProd(title string, msg string) {
	url := "https://hooks.slack.com/services/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"
	body := `{
		"channel": "#cbridge-v2-prod2-alert-status",
			"username": "%s",
			"text": "%s",
			"icon_emoji": "https://svblockchain.slack.com/services/BRARCSVU3?settings=1"
	}`

	err := httpDoPost(url, fmt.Sprintf(body, title, msg))
	if err != nil {
		log.Error("send alert error:", err)
	}
}

func httpDo(method string, url string, msg string) error {
	client := &http.Client{}
	body := bytes.NewBuffer([]byte(msg))
	req, err := http.NewRequest(method,
		url,
		body)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json;charset=utf-8")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return nil
}

// post方式
func httpDoPost(url string, msg string) error {
	return httpDo("POST", url, msg)
}

// get方式
func httpDoGet(url string, msg string) error {
	return httpDo("GET", url, msg)
}
