/*
 * @Author: berryberry
 * @since: 2019-11-07 20:39:27
 * @LastModifiedBy: berryberry
 * @LastModifiedTime: Do not edit
 */
package dingtalk_test

import (
	"fmt"
	"testing"

	"github.com/berryberrrry/ginfizz/pkg/dingtalk"
)

const (
	webhook = "https://oapi.dingtalk.com/robot/send?access_token=×××××××××"
)

func TestDingtalk_1(t *testing.T) {
	bot := dingtalk.New(webhook)
	fmt.Println(bot)
	err := bot.SendText("this is text")
	if err != nil {
		panic(err)
	}

}
