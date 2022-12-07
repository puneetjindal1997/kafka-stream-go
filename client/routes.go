package main

// Routes for disaster management
var GetDisasterLogs = Routes{
	//routes logs
	Route{"Get logs", "GET", "/logs", GetLogs},
}
