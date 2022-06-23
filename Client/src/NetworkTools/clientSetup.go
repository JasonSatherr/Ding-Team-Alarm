package NetworkTools

import (
	"fmt"
	"github.com/JasonSatherr/Ding-Team-Alarm/tree/main/Client/src/convenience"
	"net"
)

type fileBuildingClient struct{
	SERVER_TYPE string `default:"tcp"`
	connection net.Conn `default:nil`
}

//Connect is a function that connects the client to a designated host at a designated port...
func (f * fileBuildingClient) Connect(hostName string, portNumber string){

	//establish connection
	//var err Error
	//connection, err := net.Dial("tcp", hostName+":"+portNumber)
	connection, err := net.Dial("tcp", "golang.org:http")
	convenience.HandlePotErr(err)
	fmt.Print("GORKGORKGORK")
	fmt.Print(connection)
	f.connection = connection
	fmt.Print(f.connection)
	if f.connection == nil{
		fmt.Print("\n We FAILED TO MAKE A CONNECTION \n")
	}else{
		fmt.Print("\n We made A CONNECTION \n")
	}
	
}

//writeHiToConnection is a function that will write hello to an active connection
func (f fileBuildingClient) writeHiToConnection(){
	_, err := f.connection.Write([]byte("Hello, this is a generic message from the client!"))
	convenience.HandlePotErr(err)
}

//CloseConnection is a function of fileBuildingClient that will safely close the connection that it has with the data server
func (f fileBuildingClient) CloseConnection(){
	err := f.connection.Close()
	convenience.HandlePotErr(err)
	f.connection = nil
}

