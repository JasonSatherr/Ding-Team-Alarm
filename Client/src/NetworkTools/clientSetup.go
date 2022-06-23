package NetworkTools

import (
	"fmt"
	"github.com/JasonSatherr/Ding-Team-Alarm/tree/main/Client/src/convenience"
	"net"
)

type fileBuildingClient struct{
	SERVER_TYPE = "tcp"
	connection = nil
}

//Connect is a function that connects the client to a designated host at a designated port...
func (fileBuildingClient f) Connect(hostName string, portNumber string){

	//establish connection
	f.connection, err := net.Dial(f.SERVER_TYPE, hostName+":"+portNumber)
	convenience.HandlePotErr(err)
}

//writeHiToConnection is a function that will write hello to an active connection
func (fileBuildingClient f) writeHiToConnection(){
	_, err := f.connection.Write([]byte("Hello, this is a generic message from the client!"))
	convenience.HandlePotErr(err)
}

//CloseConnection is a function of fileBuildingClient that will safely close the connection that it has with the data server
func (fileBuildingClient f) CloseConnection(){
	err := f.connection.Close()
	convenience.HandlePotErr(err)
	f.connection = nil
}

