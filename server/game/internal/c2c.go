package internal

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/name5566/leaf/log"
	"io/ioutil"
	"net/http"
	"server/conf"
	"strings"
	"time"
)

type User struct {
	NickName string
	UserID   uint32
	Balance  float64
	Avatar   string
}

type UserCallback func(data *User)

type Client4Center struct {
	token         string
	conn          *websocket.Conn
	isServerLogin bool
	userWaitEvent map[string]UserCallback
}

func NewClient4Center() *Client4Center {
	wsURL := "ws" + strings.TrimPrefix(conf.Server.CenterServer, "http")
	c, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	log.Debug("连接中心服 %+v", wsURL)
	if err != nil {
		log.Fatal("dial error %v", err)
	}
	return &Client4Center{
		token:         "",
		isServerLogin: false,
		conn:          c,
		userWaitEvent: make(map[string]UserCallback),
	}
}

/***********************************************************

	请求服务器token并连接中心服

************************************************************/

// 从中心服请求token
func (c4c *Client4Center) ReqToken() {
	req, err := http.NewRequest("GET", conf.Server.TokenServer, nil)
	if err != nil {
		log.Fatal("生成请求失败")
		panic(err)
	}
	params := req.URL.Query()
	params.Add("dev_key", conf.Server.DevKey)
	params.Add("dev_name", conf.Server.DevName)
	req.URL.RawQuery = params.Encode()

	log.Debug("请求Token %+v", req.URL.String())

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		log.Fatal("请求中心服token失败", err, resp.StatusCode)
	}

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("响应体读取失败", err)
	}

	log.Debug(string(bs))
	tokenResp := TokenResp{}

	err = json.Unmarshal(bs, &tokenResp)

	if err != nil {
		log.Fatal("Token响应解析失败", err)
	}

	if tokenResp.StatusCode != 200 {
		log.Fatal("Token响应码不是200", tokenResp.StatusCode)
	}

	c4c.token = tokenResp.TokenMsg.Token

	log.Debug("Token更新完成 %+v", c4c.token)
}

/*****************************************

	监听中心服返回数据并处理

******************************************/

func (c4c *Client4Center) HeartBeatAndListen() {
	ticker := time.NewTicker(time.Minute * 60)
	go func() {
		for {
			<-ticker.C
			c4c.heartBeat()
		}
	}()

	go func() {
		for {
			log.Debug("监听服务器返回数据")

			_, message, err := c4c.conn.ReadMessage()
			if err != nil {
				log.Error("Read msg error", err.Error())
			}

			log.Debug("Msg from center %v", string(message))

			var msg Server2CenterMsg
			err = json.Unmarshal(message, &msg)
			if err != nil {
				log.Error("Json Unmarshal error", err.Error())
			}
			switch msg.Event {
			case msgServerLogin:
				c4c.onServerLogin(message)
			case msgUserLogin:
				c4c.onUserLogin(message)
			case msgUserLogout:
				c4c.onUserLogout(message)
			case msgUserLoseScore:
				c4c.onUserLoseScore(message)
			case msgUserWinScore:
				c4c.onUserWinScore(message)
			default:
				c4c.onErrorReq(message)
			}
		}
	}()

	c4c.ServerLoginCenter()
}

func (c4c *Client4Center) onServerLogin(msg []byte) {
	log.Debug("收到了中心服确认服务器登陆消息 %v", string(msg))
	sLogin := ServerLoginResp{}

	err := json.Unmarshal(msg, &sLogin)
	if err != nil {
		log.Error("解析服务器登录返回数据错误", err)
	}

	data := sLogin.Data
	status := data.Status
	code := data.Code
	taxPercent := data.Msg.PlatformTaxPercent

	c4c.isServerLogin = true
	log.Debug("%+v %+v %+v", status, code, taxPercent)
}

// 收到用户登录返回之后
func (c4c *Client4Center) onUserLogin(msg []byte) {
	loginResp := UserLoginResp{}
	err := json.Unmarshal(msg, &loginResp)
	if err != nil {
		log.Error("解析中心服返回数据出错")
	}

	userData := loginResp.Data

	code := userData.Code
	if code == centerStatusSuccess {
		gameUser := userData.Msg.GameUser
		gameAccount := userData.Msg.GameAccount

		if loginCallBack, ok := c4c.userWaitEvent[fmt.Sprintf("%+vlogin", gameUser.UserID)]; ok {
			loginCallBack(&User{
				UserID:   gameUser.UserID,
				NickName: gameUser.GameNick,
				Avatar:   gameUser.GameIMG,
				Balance:  gameAccount.Balance,
			})
		} else {
			log.Error("找不到用户回调")
		}
	} else {
		log.Error("中心服务器状态码", code)
	}
}

func (c4c *Client4Center) onUserLogout(msg []byte) {
	logoutResp := UserLogoutResp{}
	err := json.Unmarshal(msg, &logoutResp)
	if err != nil {
		log.Error("解析中心服返回数据出错")
	}

	userData := logoutResp.Data

	code := userData.Code
	if code == centerStatusSuccess {
		gameUser := userData.Msg.GameUser
		gameAccount := userData.Msg.GameAccount

		if loginCallBack, ok := c4c.userWaitEvent[fmt.Sprintf("%+vlogout", gameUser.UserID)]; ok {
			loginCallBack(&User{
				UserID:   gameUser.UserID,
				NickName: gameUser.GameNick,
				Avatar:   gameUser.GameIMG,
				Balance:  gameAccount.Balance,
			})
		} else {
			log.Error("找不到用户回调")
		}
	} else {
		log.Error("中心服务器状态码", code)
	}
}

func (c4c *Client4Center) onUserWinScore(msg []byte) {
	winResp := SyncScoreResp{}
	err := json.Unmarshal(msg, winResp)
	if err != nil {
		log.Error("解析加钱返回错误", err)
	}

	syncData := winResp.Data
	if syncData.Code != centerStatusSuccess {
		if loginCallBack, ok := c4c.userWaitEvent[fmt.Sprintf("%+vwin", syncData.Msg.ID)]; ok {
			loginCallBack(&User{})
		} else {
			log.Error("找不到用户回调")
		}
	} else {
		log.Error("中心服务器状态码", syncData.Code)
	}
}

func (c4c *Client4Center) onUserLoseScore(msg []byte) {
	loseResp := SyncScoreResp{}
	err := json.Unmarshal(msg, loseResp)
	if err != nil {
		log.Error("解析减钱返回错误", err)
	}

	syncData := loseResp.Data
	if syncData.Code != centerStatusSuccess {
		if loginCallBack, ok := c4c.userWaitEvent[fmt.Sprintf("%+vwin", syncData.Msg.ID)]; ok {
			loginCallBack(&User{})
		} else {
			log.Error("找不到用户回调")
		}
	} else {
		log.Error("中心服务器状态码", syncData.Code)
	}
}

func (c4c *Client4Center) onErrorReq(msg []byte) {
	log.Debug("错误事件 %v", string(msg))
}

/*****************************************************

	向中心服发送事件

******************************************************/

// 服务器登录中心服
func (c4c *Client4Center) ServerLoginCenter() {
	serverLoginMsg := ServerLoginReq{
		msgServerLogin,
		ServerLoginReqData{
			Host:   conf.Server.CenterServer,
			Port:   conf.Server.CenterServerPort,
			GameID: conf.Server.GameID,
			Token:  c4c.token,
			DevKey: conf.Server.DevKey,
		},
	}

	c4c.sendMsg2Center(serverLoginMsg)
}

// todo 向中心服发送心跳
func (c4c *Client4Center) heartBeat() {

}

// 操作用户数据一定要等中心服确认消息返回之后再进行展示或其他操作

// UserLoginCenter 用户登录
func (c4c *Client4Center) UserLoginCenter(userID uint32, password string, callback UserCallback) {
	if !c4c.isServerLogin {
		log.Debug("Game Server NOT Ready! Need login to Center Server!")
		return
	}

	log.Debug("UserLoginCenter c4c.Token- %+v", c4c.token)

	userLoginMsg := UserLoginReq{
		msgUserLogin,
		UserLoginReqData{
			userID,
			password,
			c4c.token,
			conf.Server.GameID,
			conf.Server.DevKey,
		},
	}

	c4c.sendMsg2Center(userLoginMsg)
	c4c.userWaitEvent[fmt.Sprintf("%+vlogin", userID)] = callback
}

// UserLogoutCenter 用户登出
func (c4c *Client4Center) UserLogoutCenter(userID uint32, callback UserCallback) {
	if !c4c.isServerLogin {
		log.Debug("Game Server NOT Ready! Need login to Center Server!")
		return
	}

	log.Debug("UserLogoutCenter c4c.Token- %+v", c4c.token)

	logoutResp := UserLogoutReq{
		Event: msgUserLogout,
		Data: UserLogoutReqData{
			UserID: userID,
			Token:  c4c.token,
			GameID: conf.Server.GameID,
			DevKey: conf.Server.DevKey,
		},
	}

	c4c.sendMsg2Center(logoutResp)
	c4c.userWaitEvent[fmt.Sprintf("%+vlogout", userID)] = callback
}

func (c4c *Client4Center) UserWinScore(userID uint32, timeUnix uint32, timeStr, payReason string, callback UserCallback) {
	if !c4c.isServerLogin {
		log.Debug("Game Server NOT Ready! Need login to Center Server!")
		return
	}

	log.Debug("UserWinScore c4c.Token- %+v", c4c.token)

	logoutResp := SyncScoreReq{
		Event: msgUserWinScore,
		Data: SyncScoreReqData{
			Auth: ServerAuth{
				Token:  c4c.token,
				DevKey: conf.Server.DevKey,
			},

			Info: SyncScoreReqDataInfo{
				UserID:     userID,
				CreateTime: timeUnix,
				PayReason:  payReason,
				Money:      3.33,
				LockMoney:  0,
				PreMoney:   0,
				Order:      fmt.Sprintf("%+v_%+v_win", userID, timeStr),
				GameID:     conf.Server.GameID,
				RoundID:    fmt.Sprintf("%+v_%+v_wintest***", userID, timeStr),
			},
		},
	}

	c4c.sendMsg2Center(logoutResp)
	c4c.userWaitEvent[fmt.Sprintf("%+vwin", userID)] = callback
}

func (c4c *Client4Center) UserLoseScore(userID uint32, timeUnix uint32, timeStr, payReason string, callback UserCallback) {
	if !c4c.isServerLogin {
		log.Debug("Game Server NOT Ready! Need login to Center Server!")
		return
	}

	log.Debug("UserWinScore c4c.Token- %+v", c4c.token)

	logoutResp := SyncScoreReq{
		Event: msgUserLoseScore,
		Data: SyncScoreReqData{
			Auth: ServerAuth{
				Token:  c4c.token,
				DevKey: conf.Server.DevKey,
			},

			Info: SyncScoreReqDataInfo{
				UserID:     userID,
				CreateTime: timeUnix,
				PayReason:  payReason,
				Money:      -1,
				LockMoney:  0,
				PreMoney:   0,
				Order:      fmt.Sprintf("%+v_%+v_lose", userID, timeStr),
				GameID:     conf.Server.GameID,
				RoundID:    fmt.Sprintf("%+v_%+v_losetest***", userID, timeStr),
			},
		},
	}

	c4c.sendMsg2Center(logoutResp)
	c4c.userWaitEvent[fmt.Sprintf("%+vlose", userID)] = callback
}

// 向中心服发送消息的基础函数
func (c4c *Client4Center) sendMsg2Center(data interface{}) {
	bs, err := json.Marshal(data)
	if err != nil {
		log.Error("解析失败", err)
	}
	log.Debug("发送数据 %v", string(bs))

	err = c4c.conn.WriteMessage(websocket.TextMessage, bs)
	if err != nil {
		log.Fatal("发送数据失败", err)
	}
}
