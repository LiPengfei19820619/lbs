package bll

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"net/http"

	"zte/ims/lbs/securityctrl/model"
)

// Start 启动安全管控SOAP消息转发服务
func Start() {
	selfIPAddr := GetSelfIPAddr()
	selfPort := GetSelfListenPort()

	startHTTPServer(selfIPAddr, selfPort)
}

func startHTTPServer(ipaddr string, port int) error {
	http.HandleFunc("/", httpHandler)
	laddr := fmt.Sprint(ipaddr, ":", port)
	http.ListenAndServe(laddr, nil)

	return nil
}

// Session 会话
type Session struct {
	sentResponse bool
	requestBody  []byte
	responseChan chan *http.Response
}

func httpHandler(w http.ResponseWriter, req *http.Request) {
	nodes := GetAllAdjacentNodes()
	if len(nodes) == 0 {
		io.WriteString(w, "No node to transfer!")
		return
	}

	fmt.Println("receive http request from security ctrl platform")
	fmt.Println("req.ContentLength:", req.ContentLength)
	fmt.Println("req.Header Content-Length:", req.Header.Get("Content-Length"))

	var session Session
	session.sentResponse = false
	session.requestBody = make([]byte, req.ContentLength)
	req.Body.Read(session.requestBody)
	session.responseChan = make(chan *http.Response, len(nodes))

	for _, node := range nodes {
		fmt.Println("node: ", node.IPAddr, ":", node.Port)
		go (&session).transferRequestToNode(req, node)
	}

	for i := 0; i < len(nodes); i++ {
		res := <-session.responseChan

		(&session).handleResponseFromNode(&w, res)
	}
}

func (session *Session) transferRequestToNode(req *http.Request, node model.AdjacentNode) error {
	fmt.Println("transfer request to node: ", node.IPAddr, ":", node.Port)
	url := req.URL
	url.Scheme = "http"
	url.Host = fmt.Sprint(node.IPAddr, ":", node.Port)

	/* body := make([]byte, 10240)
	req.Body.Read(body)
	fmt.Println("req.body:", string(body)) */

	reqNew, err := http.NewRequest(req.Method, url.String(), bytes.NewReader(session.requestBody))
	if err != nil {
		fmt.Println("http.NewResquent error:", err.Error())
		session.responseChan <- nil
		return err
	}

	copyHeader(reqNew.Header, req.Header)

	//reqNew.Header.Add("Content-Length", req.Header.Get("Content-Length"))
	reqNew.Host = fmt.Sprint(node.IPAddr, ":", node.Port)
	reqNew.ContentLength = req.ContentLength

	/* reqNew.Body.Read(body)
	fmt.Println("reqNew.body:", string(body)) */

	client := http.Client{}
	res, err := client.Do(reqNew)
	if err != nil {
		fmt.Println("client.Do error:", err.Error())
		session.responseChan <- nil
		return err
	}

	for key := range reqNew.Header {
		fmt.Println("request to node[", node.IPAddr, ":", node.Port, "], ", key, ":", reqNew.Header.Get(key))
	}

	session.responseChan <- res
	return nil
}

func (session *Session) handleResponseFromNode(w *http.ResponseWriter, res *http.Response) {
	if res == nil {
		return
	}
	fmt.Println("receive http response from node")
	defer res.Body.Close()

	if !session.sentResponse {
		copyHeader((*w).Header(), res.Header)

		io.Copy(*w, res.Body)

		session.sentResponse = true
	}
}

func copyHeader(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}

func startTCPServer(ipaddr string, port int) error {
	laddr := fmt.Sprint(ipaddr, ":", port)
	listerer, err := net.Listen("tcp", laddr)

	if err != nil {
		fmt.Println("Error Listerning, error:", err.Error())
		return err
	}
	// Close the listener when the application closes.
	defer listerer.Close()

	fmt.Println("Listerning on: ", laddr)

	for {
		conn, err := listerer.Accept()
		if err != nil {
			fmt.Println("Error Accepting: ", err.Error())
			break
		}

		go handleRequest(conn)
	}

	return nil
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {
	// Make a buffer to hold incoming data.
	buf := make([]byte, 1024)
	// Read the incoming connection into the buffer.
	reqLen, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	fmt.Println("Receive: ", buf, "len: ", reqLen)

	// Send a response back to person contacting us.
	conn.Write([]byte("Message received."))
	// Close the connection when you're done with it.
	conn.Close()
}
