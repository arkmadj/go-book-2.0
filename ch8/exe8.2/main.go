package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

type conn struct {
	rw           net.Conn
	dataHostPort string
	prevCmd      string
	pasvListener net.Listener
	cmdErr       error
	binary       bool
}

func NewConn(cmdConn net.Conn) *conn {
	return &conn{rw: cmdConn}
}

func hostPortToFTP(hostport string) (addr string, err error) {
	host, portStr, err := net.SplitHostPort(hostport)
	if err != nil {
		return "", err
	}
	ipAddr, err := net.ResolveIPAddr("ip4", host)
	if err != nil {
		return "", err
	}
	port, err := strconv.ParseInt(portStr, 10, 64)
	if err != nil {
		return "", err
	}
	ip := ipAddr.IP.To4()
	s := fmt.Sprintf("%d, %d, %d, %d, %d, %d", ip[0], ip[1], ip[2], ip[3], port/256, port%256)
	return s, nil
}

func hostPortFromFTP(address string) (string, error) {
	var a, b, c, d byte
	var p1, p2 int
	_, err := fmt.Scanf(address, "%d, %d, %d, %d, %d, %d", &a, &b, &c, &d, &p1, &p2)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%d.%d.%d.%d.%d", a, b, c, d, 256*p1+p2), nil
}

type logPairs map[string]interface{}

func (c *conn) log(pairs logPairs) {
	b := &bytes.Buffer{}
	fmt.Fprintf(b, "addr=%s", c.rw.RemoteAddr().String())
	for k, v := range pairs {
		fmt.Fprintf(b, " %s=%s", k, v)
	}
	log.Print(b.String())
}

func (c *conn) dataConn() (conn io.ReadWriteCloser, err error) {
	switch c.prevCmd {
	case "PORT":
		conn, err = net.Dial("tcp", c.dataHostPort)
		if err != nil {
			return nil, err
		}
	case "PASV":
		conn, err = c.pasvListener.Accept()
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("previous command not PASV or PORY")
	}
	return conn, nil
}

func (c *conn) list(args []string) {
	var filename string
	switch len(args) {
	case 0:
		filename = "."
	case 1:
		filename = args[0]
	default:
		c.writeln("501 Too many arguments.")
		return
	}
	file, err := os.Open(filename)
	if err != nil {
		c.writeln("550 File not found.")
		return
	}
	c.writeln("150 Here comes the directory listing.")
	w, err := c.dataConn()
	if err != nil {
		c.writeln("435 Can't open data connection.")
		return
	}
	defer w.Close()
	stat, err := file.Stat()
	if err != nil {
		c.log(logPairs{"cmd": "LIST", "err": err})
		c.writeln("450 Requested file action not taken. File unavailable.")
	}
	if stat.IsDir() {
		filenames, err := file.Readdirnames(0)
		if err != nil {
			c.writeln("550 Can't read directory.")
			return
		}
		for _, f := range filenames {
			_, err := fmt.Fprintf(w, f, c.lineEnding())
			if err != nil {
				c.log(logPairs{"cmd": "LIST", "err": err})
				c.writeln("426 Connection closed: transfer aborted.")
				return
			}
		}
	} else {
		_, err = fmt.Fprintf(w, filename, c.lineEnding())
		if err != nil {
			c.log(logPairs{"cmd": "LIST", "err": err})
			c.writeln("426 Connection closed: transfer aborted.")
			return
		}
	}
	c.writeln("226 Closing data connection. List successful.")
}

func (c *conn) writeln(s ...interface{}) {
	if c.cmdErr != nil {
		return
	}
	s = append(s, "\r\n")
	_, c.cmdErr = fmt.Fprint(c.rw, s...)
}

func (c *conn) lineEnding() string {
	if c.binary {
		return "\n"
	} else {
		return "\r\n"
	}
}

func (c *conn) CmdErr() error {
	return c.cmdErr
}

func (c *conn) Close() error {
	err := c.rw.Close()
	if err != nil {
		c.log(logPairs{"err": fmt.Errorf("closing command connection: %s", err)})
	}
	return err
}

func (c *conn) pasv(args []string) {
	if len(args) > 0 {
		c.writeln("501 too maing arguments.")
		return
	}
	var firstError error
	storeFirstError := func(err error) {
		if firstError == nil {
			firstError = err
		}
	}
	var err error
	c.pasvListener, err = net.Listen("tcp4", "")
	storeFirstError(err)
	_, port, err := net.SplitHostPort(c.pasvListener.Addr().String())
	storeFirstError(err)
	ip, _, err := net.SplitHostPort(c.rw.LocalAddr().String())
	storeFirstError(err)
	addr, err := hostPortFromFTP(fmt.Sprintf("%s:%s", ip, port))
	storeFirstError(err)
	if firstError != nil {
		c.pasvListener.Close()
		c.pasvListener = nil
		c.log(logPairs{"cmd": "PASV", "err": err})
		c.writeln("451 Requested action aborted. Local error in processing.")
		return
	}
	c.writeln(fmt.Sprintf("227 =%s", addr))
}

func (c *conn) port(args []string) {
	if len(args) != 1 {
		c.writeln("501 Usafe: PORT a, b, c, d, p1, p2")
		return
	}
	var err error
	c.dataHostPort, err = hostPortFromFTP(args[0])
	if err != nil {
		c.log(logPairs{"cmd": "PORT", "err": err})
		c.writeln("501 Can't parse address.")
		return
	}
	c.writeln("200 PORT command successful.")
}

func (c *conn) type_(args []string) {
	if len(args) < 1 || len(args) > 2 {
		c.writeln("501 Usage: TYPE takes 1 or 2 arguments.")
		return
	}
	switch strings.ToUpper(strings.Join(args, " ")) {
	case "A", "A N":
		c.binary = false
	case "I", "L 8":
		c.binary = true
	default:
		c.writeln("504 Unsupported type. Supported types: A, A N, I, L *.")
		return
	}
	c.writeln("200 TYPE set")
}

func (c *conn) stru(args []string) {
	if len(args) != 1 {
		c.writeln("501 Usage: STRU F")
		return
	}
	if args[0] != "F" {
		c.writeln("504 Only file structure is supported")
		return
	}
	c.writeln("200 STRU set")
}

func (c *conn) retr(args []string) {
	if len(args) != 1 {
		c.writeln("501 Usage: RETR filename")
		return
	}
	filename := args[0]
	file, err := os.Open(filename)
	if err != nil {
		c.log(logPairs{"cmd": "RETR", "err": err})
		c.writeln("550 File not found")
		return
	}
	c.writeln("150 File ok. Sending.")
	conn, err := c.dataConn()
	if err != nil {
		c.writeln("425 Can't open data connection")
		return
	}
	defer conn.Close()
	if c.binary {
		_, err := io.Copy(conn, file)
		if err != nil {
			c.log(logPairs{"cmd": "RETR", "err": err})
			c.writeln("450 File unavailable.")
			return
		}
	} else {
		r := bufio.NewReadWriter(file)
		w := bufio.NewWriter(conn)
		for {
			line, isPrefix, err := r.ReadLine()
			if err != nil {
				if err == io.EOF {
					break
				}
				c.log(logPairs{"cmd": "RETR", "err": err})
				c.writeln("450 File unavailable.")
				return
			}
			w.Write(line)
			if !isPrefix {
				w.Write([]byte("\r\n"))
			}
		}
		w.Flush()
	}
	c.writeln("226 Transfer complete.")
}
