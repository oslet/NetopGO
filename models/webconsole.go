package models

import (
	"bytes"
	"fmt"
	"github.com/astaxie/beego"
	gossh "golang.org/x/crypto/ssh"
	"golang.org/x/net/websocket"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type SSH struct {
	User    string
	Pwd     string
	Addr    string
	Client  *gossh.Client
	Session *gossh.Session
}

func (s *SSH) Connect() (*SSH, error) {
	config := &gossh.ClientConfig{}
	config.SetDefaults()
	config.User = s.User
	config.Auth = []gossh.AuthMethod{gossh.Password(s.Pwd)}
	client, err := gossh.Dial("tcp", s.Addr, config)
	if nil != err {
		return nil, err
	}
	s.Client = client
	session, err := client.NewSession()
	if nil != err {
		return nil, err
	}
	s.Session = session
	return s, nil
}

func (s *SSH) Exec(cmd string) (string, error) {
	var buf bytes.Buffer
	s.Session.Stdout = &buf
	s.Session.Stderr = &buf
	err := s.Session.Run(cmd)
	if err != nil {
		return "", err
	}
	defer s.Session.Close()
	stdout := buf.String()
	fmt.Printf("Stdout:%v\n", stdout)
	return stdout, nil
}

func SSHWebSocketHandler(ws *websocket.Conn) {
	ctx := NewContext(nil, ws.Request())

	vm_info := ctx.GetFormValue("vm_info")
	cols := Str2Int(ctx.GetFormValue("cols"))
	rows := Str2Int(ctx.GetFormValue("rows"))
	vm_info = strings.Replace(vm_info, " ", "+", 1)

	beego.Info(vm_info)
	beego.Info(cols)
	beego.Info(rows)

	de_vm_info, err := AESDecode(vm_info, AesKey)
	beego.Info(de_vm_info)
	de_vm_info_arr := strings.Split(de_vm_info, "\n")
	user := strings.TrimSpace(de_vm_info_arr[0])
	passwd := strings.TrimSpace(de_vm_info_arr[1])
	vmAddr := strings.TrimSpace(de_vm_info_arr[2])
	beego.Info(user)
	beego.Info(passwd)
	beego.Info(vmAddr)

	sh := &SSH{
		User: user,
		Pwd:  passwd,
		Addr: vmAddr,
	}

	sh, err = sh.Connect()
	if nil != err {
		beego.Error(err)
		return
	}

	session := sh.Session
	defer session.Close()
	modes := gossh.TerminalModes{
		gossh.ECHO:          1,
		gossh.TTY_OP_ISPEED: 14400,
		gossh.TTY_OP_OSPEED: 14400,
	}

	if err = session.RequestPty("xterm-256color", rows, cols, modes); err != nil {
		beego.Error(err)
		return
	}

	w, err := session.StdinPipe()
	if nil != err {
		beego.Error(err)
		return
	}
	go func() {
		io.Copy(w, ws)
	}()

	r, err := session.StdoutPipe()
	if nil != err {
		beego.Error(err)
		return
	}
	go func() {
		io.Copy(ws, r)
	}()

	er, err := session.StderrPipe()
	if nil != err {
		beego.Error(err)
		return
	}
	go func() {
		io.Copy(ws, er)
	}()

	if err := session.Shell(); nil != err {
		beego.Error(err)
		return
	}

	if err := session.Wait(); nil != err {
		beego.Error(err)
		return
	}

}

type Context struct {
	r *http.Request
	w http.ResponseWriter
	v map[string]interface{}
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	ctx := &Context{
		r: r,
		w: w,
	}
	err := ctx.parseForm()
	if nil != err {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	return ctx
}

func (c *Context) parseForm() error {
	err := c.r.ParseForm()
	if nil != err {
		return err
	}
	paramMap := make(map[string]interface{})
	s := c.r.Form
	for k, v := range s {
		if nil != paramMap[k] {
			paramArr := make([]interface{}, 0, 0)
			paramArr = append(paramArr, paramMap[k])
			paramArr = append(paramArr, v)
		} else {
			paramMap[k] = v
		}
	}
	c.v = paramMap
	return nil
}

func (c *Context) GetFormValue(key string) string {
	fv := c.v[key]
	if nil != fv {
		return strings.TrimSpace(fv.([]string)[0])
	} else {
		return ""
	}
}

func Str2Int(s string) int {
	v, _ := strconv.ParseInt(s, 10, 32)
	return int(v)
}
