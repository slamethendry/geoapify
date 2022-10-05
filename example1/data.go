package main

import (
	p "github.com/slamethendry/geoapify"
)

var office = p.LonLat{12.515119332498072, 41.896917189378605} // San Lorenzo

var workPeriod = [2]uint{0, 25200}

var agents = []p.Agent{
	{
		ID:          "004",
		Start:       office,
		End:         p.LonLat{12.508980161984935, 41.8246168491674}, // Meravigliosa
		TimeWindows: [][2]uint{workPeriod},
	},
	{
		ID:          "007",
		Start:       office,
		End:         p.LonLat{12.442004926398404, 41.805514283894546}, // Mostacciano
		TimeWindows: [][2]uint{workPeriod},
	},
}

var agentColours = []string{"orange", "green", "blue", "red"}

var jobs = []p.Job{
	{
		ID:       "Pantheon",
		LonLat:   p.LonLat{12.477070552476265, 41.89900899218839},
		Duration: 3600,
	},
	{
		ID:       "Muratella",
		LonLat:   p.LonLat{12.397600148423832, 41.83036665198902},
		Duration: 3600,
	},
	{
		ID:       "Foro Romano",
		LonLat:   p.LonLat{12.485081946228776, 41.89250838299456},
		Duration: 1800,
	},
	{
		ID:       "Anagnina",
		LonLat:   p.LonLat{12.58896415803558, 41.8408469922023},
		Duration: 3600,
	},
	{
		ID:       "Colosseo",
		LonLat:   p.LonLat{12.492317861372474, 41.8914836916335},
		Duration: 3600,
	},
	{
		ID:       "Ciampino",
		LonLat:   p.LonLat{12.603442297616573, 41.803032794020844},
		Duration: 1800,
	},
	{
		ID:       "San Pietro",
		LonLat:   p.LonLat{12.453992479516694, 41.903639665810346},
		Duration: 1800,
	},
	{
		ID:       "San Angelo",
		LonLat:   p.LonLat{12.466966091909574, 41.90352142276983},
		Duration: 1800,
	},
	{
		ID:       "Trevi",
		LonLat:   p.LonLat{12.483584114098038, 41.90245787178557},
		Duration: 7200,
	},
}

var request = p.Request{
	Mode:    "motorcycle",
	Agents:  agents,
	Jobs:    jobs,
	Traffic: "approximated",
}

var batchRequest = p.BatchRequest{
	Inputs: []p.BatchInput{
		{
			Body: request,
		},
	},
}
