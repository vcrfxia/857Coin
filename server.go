package main

import (
	"math/rand"
	"fmt"
	"math"
)

type triplet struct{
	//TODO: fill this out 
}

type server struct{
	peers []chan triplet
	me int 
	num_processors int 
	stored_triplets []triplet
	rand_generator *rand.Rand
	reply_channel chan triplet
}

/*
Function to create a new server
@param peers list of peers
@param me: the server's index into peers
@param num_processors: number of processors
*/
func Make(peers []chan triplet, me int, num_processors int, reply_channel chan triplet) *server{

	//store all variable into server state
	sv := &server{}
	sv.peers = peers
	sv.me = me
	sv.num_processors = num_processors

	sv.rand_generator = rand.New(rand.NewSource(int64(sv.me)))

	return sv
}

func (sv *server) start(){
	fmt.Println("starting server", sv.me)

	//TODO: implement this
}


func construct_triplets(){

}



func main(){
	N := math.Exp2(20)
	num_servers := int(math.Pow(N,0.333))

	fmt.Println("Running server to solve puzzle with N:", N)

	//channel to push replies
	reply_channel := make(chan triplet)

	//list of channels to give to servers
	channels := make([]chan triplet, num_servers)

	//make list of channels
	for i := 0; i < num_servers; i++{
		channels[i] = make(chan triplet)
	}

	fmt.Println("Done initializing channels")

	//list of servers 
	servers := make([]server, num_servers)

	//make a fuckton of servers
	for i := 0; i < num_servers; i++{
		servers[i] = *Make(channels, i, num_servers, reply_channel)
	}

	//start all servers
	for i:= 0; i < num_servers; i ++{
		servers[i].start()
	}

	fmt.Println("Starting all servers, waiting for results")

	//print out anything the servers return 
	for ret := range(reply_channel){
		fmt.Println("got", ret)
	}
}