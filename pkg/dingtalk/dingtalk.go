/*
 * @Author: berryberry
 * @since: 2019-11-07 20:37:31
 * @LastModifiedBy: berryberry
 * @LastModifiedTime: Do not edit
 */
package dingtalk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Bot struct {
	webhook string
}

func New(webhook string) *Bot {
	bot := new(Bot)
	bot.webhook = webhook
	return bot
}

// Webhook set the webhook where the message to send
func (this *Bot) Webhook(webhook string) *Bot {
	this.webhook = webhook
	return this
}

func (this *Bot) SendText(content string) error {
	message := map[string]interface{}{
		"msgtype": "text",
		"text": map[string]interface{}{
			"content": content,
		},
	}

	data, err := json.Marshal(message)
	if err != nil {
		return err
	}

	resp, err := http.Post(this.webhook, "application/json", bytes.NewReader(data))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	respBodyData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var respBody map[string]interface{}
	err = json.Unmarshal(respBodyData, &respBody)
	if err != nil {
		return err
	}

	var ok bool

	fmt.Println(string(respBodyData))
	// 解析ErrorCode
	var errCode float64
	_errCode, ok := respBody["errcode"]
	if !ok {
		return fmt.Errorf("errcode not found in response json")
	}
	if errCode, ok = _errCode.(float64); !ok {
		return fmt.Errorf("errcode is not int type")
	}

	// 如果ErrorCode不为0, 说明钉钉API调用失败
	if errCode != 0 {
		var errMessage string
		_errMessage, ok := respBody["errmsg"]
		if !ok {
			return fmt.Errorf("errmsg not found in response json")
		}
		if errMessage, ok = _errMessage.(string); !ok {
			return fmt.Errorf("errmsg is not string type")
		}
		return fmt.Errorf("errcode is not 0(send message faild): %s", errMessage)
	}

	return nil
}

func (this *Bot) SendMarkdownText(content string) error {
	message := map[string]interface{}{
		"msgtype": "markdown",
		"markdown": map[string]interface{}{
			"title": "ginfizz error!",
			"text":  content,
		},
	}

	data, err := json.Marshal(message)
	if err != nil {
		return err
	}

	resp, err := http.Post(this.webhook, "application/json", bytes.NewReader(data))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	respBodyData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var respBody map[string]interface{}
	err = json.Unmarshal(respBodyData, &respBody)
	if err != nil {
		return err
	}

	var ok bool

	// 解析ErrorCode
	var errCode float64
	_errCode, ok := respBody["errcode"]
	if !ok {
		return fmt.Errorf("errcode not found in response json")
	}
	if errCode, ok = _errCode.(float64); !ok {
		return fmt.Errorf("errcode is not int type")
	}

	// 如果ErrorCode不为0, 说明钉钉API调用失败
	if errCode != 0 {
		var errMessage string
		_errMessage, ok := respBody["errmsg"]
		if !ok {
			return fmt.Errorf("errmsg not found in response json")
		}
		if errMessage, ok = _errMessage.(string); !ok {
			return fmt.Errorf("errmsg is not string type")
		}
		return fmt.Errorf("errcode is not 0(send message faild): %s", errMessage)
	}

	return nil

}
