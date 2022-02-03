package socket

import (
        "net"
        "os"
        "fmt"
)

var conn net.Conn
var hostName string
var err error

func ConnLogServer() (error) {

	//You should insert the log server's IP address directly on the code or map it to the DNS address within the host file.
        conn, err = net.Dial("tcp", "logServer:9617")
        if nil != err{
                fmt.Println("GDCPE Log Server connection error")
        }
        hostName, err = os.Hostname()

        return err
}

func WriteEvalLog(msgFromGDCPEA string) (error) {
        conn.Write([]byte(hostName+msgFromGDCPEA+"\n"))
        return nil
}
