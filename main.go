package main

import (
	"bufio"
	"os"
	"io"
	"regexp"
	"strings"
	"fmt"
	"os/exec"
	"net/http"
	"bytes"
	"io/ioutil"
	"encoding/json"
	//"ember/http/rpc"
	//"ember/cli"
)

func main() {
	//file := newRawFile("url.txt.bak")

	//key := "http"
	scanner, err := OpenScanner()
	if err != nil {
		return
	}

	domain := "http://7xit6e.com1.z0.glb.clouddn.com/"

	//reg := regexp.MustCompile(`[^\t\n]+`)
	reg := regexp.MustCompile(`[^ ]+`)
	var x []string
	var url string
	var tag string
	var title string
	var idx int
	var i int
	for scanner.scanner.Scan() {
		line := scanner.scanner.Text()
		if len(line) < 20 {
			continue
		}
		i = i + 1
		word := reg.FindAllString(line, -1)
		/*
		if strings.Contains(word[2], key) {
			x = append(x, word)
		}
		*/
		idx = strings.Index(line, word[1])
		idx = idx + len(word[1]) + 1
		title = string(([]byte(line))[idx:])
		url = word[0]
		tag = word[1]
		_ = url
		_ = tag
		_ = title
		//println(url, tag, title)
		qiniuurl := fmt.Sprintf("%s%d%s",domain ,i ,".mp3")
		println(qiniuurl)
		//x = append(x, word[0])
		//x = append(x, word[1])
		//x = append(x, title)
	}
	scanner.Close()
	fmt.Printf("[x:%#v]\n", x)
	//return x, err

	XiamiaDownload()

	msg, err := json.Marshal("ojb")
	if err != nil {
		println(err.Error())
	}
	ret, err := postNewMusic(string(msg))
	_ = ret
	if err != nil {
		println(err.Error())
	}

	/*
	_ = file
	name := "1.mp3"
	println(domain + name)
	println("hello lua")
	*/

	hub := cli.NewRpcHub(os.Args[1:], &Server{}, &Client{})
	hub.Run()

}

func postNewMusic(msg string) (ret string, err error) {
	url := "http://127.0.0.1:9910/NewMusic"
	var resp *http.Response

	resp, err = http.Post(url, "text/json", bytes.NewReader([]byte(msg)))
	if err != nil {
		return
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	ret = string(data)
	return

}

func XiamiaDownload () {
	//arv := []string{"-a"}
	//c := exec.Command("./run.sh", arv...)
	c := exec.Command("./run.sh")
	d,_ := c.Output()
	println(string(d))
	/*
	var output bytes.Buffer
	//cmd := exec.Command("cat")
	//cmd := exec.Command("python main.py ")
	//cmd := exec.Command("./main.py ")
	cmd := exec.Command("/usr/bin/python main.py ")
	cmd.Stdout = &output
	stdin, _ := cmd.StdinPipe()
	cmd.Start() //执行开始
	stdin.Write([]byte("http://www.xiami.com/song/1769629623"))
	stdin.Close()
	cmd.Wait()                                        //等待执行完成
	fmt.Printf("The output is: %s\n", output.Bytes()) //The output is: widuu test!
	*/

}

func OpenScanner() (scanner Scanner, err error) {
	fd, err := os.OpenFile("url.txt.bak", os.O_RDWR | os.O_APPEND | os.O_CREATE, 0640)
	if err != nil {
		println(err.Error())
		return scanner,err
	}
	return Scanner{bufio.NewScanner(fd), fd}, err
}

func (p *Scanner) Close() ( err error) {
	p.fd.Close()
	return
}

type Scanner struct {
	scanner *bufio.Scanner
	fd *os.File
}

func (p *RawFile) Read() (ret string, err error) {
	r := bufio.NewReader(p.fd)
	for {
		line, err := r.ReadString('\n')
		if io.EOF == err || nil != err {
			break
		}
		ret = ret + line
	}
	return ret, err
}

func newRawFile(name string) RawFile {
	fd, err := os.OpenFile(name, os.O_RDWR | os.O_APPEND | os.O_CREATE, 0640)
	if err != nil {
	}
	return RawFile{fd, name}
}

type RawFile struct {
	fd *os.File
	name string
}


func (p *Server) Echo(msg string) (echo string, err error) {
	echo = msg
	return
}

func (p *Server) Trait() map[string][]string {
	return map[string][]string {
		"Echo": {"msg"},
	}
}

type Client struct {
	Echo func(msg string) (echo string, err error)
}

type Server struct{}

