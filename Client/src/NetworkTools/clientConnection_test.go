package NetworkTools

import ("testing"
)

//make sure that the server is up first...
//
func TestCanMakeConnection(t *testing.T) {
	client := fileBuildingClient{}
	client.Connect("golang.org","http")

	if client.connection == nil{
		t.Errorf("We failed to make the connection to golang :(")
	}
    // total := Sum(5, 5)
    // if total != 10 {
    //    t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
    // }
}