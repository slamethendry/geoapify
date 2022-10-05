package planner_test

import (
	p "github.com/slamethendry/geoapify/planner"
)

var reqJSON = []byte(`{"mode":"drive","agents":[{"start_location":[10.985736727197894,48.2649593],"end_location":[10.896261152517647,48.33227795],"pickup_capacity":4},{"start_location":[10.984995162319564,48.264409],"end_location":[10.896261152517647,48.33227795],"pickup_capacity":7},{"start_location":[10.990289270853658,48.2675442],"end_location":[10.896261152517647,48.33227795],"pickup_capacity":4},{"start_location":[10.987279662433057,48.2677992],"end_location":[10.896261152517647,48.33227795],"pickup_capacity":7},{"start_location":[10.983010635351173,48.262753950000004],"end_location":[10.896261152517647,48.33227795],"pickup_capacity":4}],"jobs":[{"location":[10.98698105,48.25615875],"duration":300,"pickup_amount":1},{"location":[10.9845547,48.26311145],"duration":300,"pickup_amount":1},{"location":[10.984630924828402,48.263248250000004],"duration":300,"pickup_amount":2},{"location":[10.968364837855287,48.262043399999996],"duration":300,"pickup_amount":1},{"location":[10.984364769628737,48.25542385],"duration":300,"pickup_amount":1},{"location":[10.984062746838354,48.25549435],"duration":300,"pickup_amount":1},{"location":[10.983802751265776,48.25558745],"duration":300,"pickup_amount":2},{"location":[10.983222005227521,48.255775],"duration":300,"pickup_amount":1},{"location":[10.983499356818182,48.25569725],"duration":300,"pickup_amount":1},{"location":[10.982919152872745,48.2558497],"duration":300,"pickup_amount":1},{"location":[10.983681544239769,48.25621035],"duration":300,"pickup_amount":2},{"location":[10.983236456481574,48.2560687],"duration":300,"pickup_amount":1},{"location":[10.984312143079265,48.25577875],"duration":300,"pickup_amount":1},{"location":[10.981143603167904,48.257296600000004],"duration":300,"pickup_amount":1},{"location":[10.9807393,48.25748695],"duration":300,"pickup_amount":1},{"location":[10.981209348235852,48.25786594111741],"duration":300,"pickup_amount":1},{"location":[10.980955539642784,48.2562265],"duration":300,"pickup_amount":1},{"location":[10.979089323915998,48.25726365],"duration":300,"pickup_amount":1},{"location":[10.979089323915998,48.25726365],"duration":300,"pickup_amount":1},{"location":[10.978800955841443,48.25723825],"duration":300,"pickup_amount":1}]}`)

var resJSON = []byte(`{"type":"FeatureCollection","properties":{"mode":"drive","params":{"mode":"drive","agents":[{"start_location":[10.985736727197894,48.2649593],"end_location":[10.896261152517647,48.33227795],"pickup_capacity":4},{"start_location":[10.984995162319564,48.264409],"end_location":[10.896261152517647,48.33227795],"pickup_capacity":7},{"start_location":[10.990289270853658,48.2675442],"end_location":[10.896261152517647,48.33227795],"pickup_capacity":4},{"start_location":[10.987279662433057,48.2677992],"end_location":[10.896261152517647,48.33227795],"pickup_capacity":7},{"start_location":[10.983010635351173,48.262753950000004],"end_location":[10.896261152517647,48.33227795],"pickup_capacity":4}],"jobs":[{"location":[10.98698105,48.25615875],"duration":300,"pickup_amount":1},{"location":[10.9845547,48.26311145],"duration":300,"pickup_amount":1},{"location":[10.984630924828402,48.263248250000004],"duration":300,"pickup_amount":2},{"location":[10.968364837855287,48.262043399999996],"duration":300,"pickup_amount":1},{"location":[10.984364769628737,48.25542385],"duration":300,"pickup_amount":1},{"location":[10.984062746838354,48.25549435],"duration":300,"pickup_amount":1},{"location":[10.983802751265776,48.25558745],"duration":300,"pickup_amount":2},{"location":[10.983222005227521,48.255775],"duration":300,"pickup_amount":1},{"location":[10.983499356818182,48.25569725],"duration":300,"pickup_amount":1},{"location":[10.982919152872745,48.2558497],"duration":300,"pickup_amount":1},{"location":[10.983681544239769,48.25621035],"duration":300,"pickup_amount":2},{"location":[10.983236456481574,48.2560687],"duration":300,"pickup_amount":1},{"location":[10.984312143079265,48.25577875],"duration":300,"pickup_amount":1},{"location":[10.981143603167904,48.257296600000004],"duration":300,"pickup_amount":1},{"location":[10.9807393,48.25748695],"duration":300,"pickup_amount":1},{"location":[10.981209348235852,48.25786594111741],"duration":300,"pickup_amount":1},{"location":[10.980955539642784,48.2562265],"duration":300,"pickup_amount":1},{"location":[10.979089323915998,48.25726365],"duration":300,"pickup_amount":1},{"location":[10.979089323915998,48.25726365],"duration":300,"pickup_amount":1},{"location":[10.978800955841443,48.25723825],"duration":300,"pickup_amount":1}]}},"features":[{"geometry":{"type":"MultiLineString","coordinates":[[[10.985737,48.264959],[10.984631,48.263248]],[[10.984631,48.263248],[10.984555,48.263111]],[[10.984555,48.263111],[10.986981,48.256159]],[[10.986981,48.256159],[10.896261,48.332278]]]},"type":"Feature","properties":{"agent_index":0,"time":2066,"start_time":0,"end_time":2066,"distance":20977,"legs":[{"time":42,"distance":290,"from_waypoint_index":0,"to_waypoint_index":1,"steps":[{"from_index":0,"to_index":1,"time":42,"distance":290}]},{"time":2,"distance":5,"from_waypoint_index":1,"to_waypoint_index":2,"steps":[{"from_index":0,"to_index":1,"time":2,"distance":5}]},{"time":80,"distance":906,"from_waypoint_index":2,"to_waypoint_index":3,"steps":[{"from_index":0,"to_index":1,"time":80,"distance":906}]},{"time":1042,"distance":19776,"from_waypoint_index":3,"to_waypoint_index":4,"steps":[{"from_index":0,"to_index":1,"time":1042,"distance":19776}]}],"mode":"drive","actions":[{"index":0,"type":"start","start_time":0,"duration":0,"waypoint_index":0},{"index":1,"type":"job","start_time":42,"duration":300,"job_index":2,"waypoint_index":1},{"index":2,"type":"job","start_time":344,"duration":300,"job_index":1,"waypoint_index":2},{"index":3,"type":"job","start_time":724,"duration":300,"job_index":0,"waypoint_index":3},{"index":4,"type":"end","start_time":2066,"duration":0,"waypoint_index":4}],"waypoints":[{"original_location":[10.985736727197894,48.2649593],"location":[10.985737,48.264959],"start_time":0,"duration":0,"actions":[{"index":0,"type":"start","start_time":0,"duration":0,"waypoint_index":0}],"next_leg_index":0},{"original_location":[10.984630924828402,48.263248250000004],"location":[10.984631,48.263248],"start_time":42,"duration":300,"actions":[{"index":1,"type":"job","start_time":42,"duration":300,"job_index":2,"waypoint_index":1}],"prev_leg_index":0,"next_leg_index":1},{"original_location":[10.9845547,48.26311145],"location":[10.984555,48.263111],"start_time":344,"duration":300,"actions":[{"index":2,"type":"job","start_time":344,"duration":300,"job_index":1,"waypoint_index":2}],"prev_leg_index":1,"next_leg_index":2},{"original_location":[10.98698105,48.25615875],"location":[10.986981,48.256159],"start_time":724,"duration":300,"actions":[{"index":3,"type":"job","start_time":724,"duration":300,"job_index":0,"waypoint_index":3}],"prev_leg_index":2,"next_leg_index":3},{"original_location":[10.896261152517647,48.33227795],"location":[10.896261,48.332278],"start_time":2066,"duration":0,"actions":[{"index":4,"type":"end","start_time":2066,"duration":0,"waypoint_index":4}],"prev_leg_index":3}]}},{"geometry":{"type":"MultiLineString","coordinates":[[[10.984995,48.264409],[10.983682,48.25621]],[[10.983682,48.25621],[10.983803,48.255587]],[[10.983803,48.255587],[10.984063,48.255494]],[[10.984063,48.255494],[10.984312,48.255779]],[[10.984312,48.255779],[10.984365,48.255424]],[[10.984365,48.255424],[10.896261,48.332278]]]},"type":"Feature","properties":{"agent_index":1,"time":2752,"start_time":0,"end_time":2752,"distance":21348,"legs":[{"time":117,"distance":1063,"from_waypoint_index":0,"to_waypoint_index":1,"steps":[{"from_index":0,"to_index":1,"time":117,"distance":1063}]},{"time":15,"distance":75,"from_waypoint_index":1,"to_waypoint_index":2,"steps":[{"from_index":0,"to_index":1,"time":15,"distance":75}]},{"time":13,"distance":37,"from_waypoint_index":2,"to_waypoint_index":3,"steps":[{"from_index":0,"to_index":1,"time":13,"distance":37}]},{"time":9,"distance":15,"from_waypoint_index":3,"to_waypoint_index":4,"steps":[{"from_index":0,"to_index":1,"time":9,"distance":15}]},{"time":8,"distance":36,"from_waypoint_index":4,"to_waypoint_index":5,"steps":[{"from_index":0,"to_index":1,"time":8,"distance":36}]},{"time":1090,"distance":20122,"from_waypoint_index":5,"to_waypoint_index":6,"steps":[{"from_index":0,"to_index":1,"time":1090,"distance":20122}]}],"mode":"drive","actions":[{"index":0,"type":"start","start_time":0,"duration":0,"waypoint_index":0},{"index":1,"type":"job","start_time":117,"duration":300,"job_index":10,"waypoint_index":1},{"index":2,"type":"job","start_time":432,"duration":300,"job_index":6,"waypoint_index":2},{"index":3,"type":"job","start_time":745,"duration":300,"job_index":5,"waypoint_index":3},{"index":4,"type":"job","start_time":1054,"duration":300,"job_index":12,"waypoint_index":4},{"index":5,"type":"job","start_time":1362,"duration":300,"job_index":4,"waypoint_index":5},{"index":6,"type":"end","start_time":2752,"duration":0,"waypoint_index":6}],"waypoints":[{"original_location":[10.984995162319564,48.264409],"location":[10.984995,48.264409],"start_time":0,"duration":0,"actions":[{"index":0,"type":"start","start_time":0,"duration":0,"waypoint_index":0}],"next_leg_index":0},{"original_location":[10.983681544239769,48.25621035],"location":[10.983682,48.25621],"start_time":117,"duration":300,"actions":[{"index":1,"type":"job","start_time":117,"duration":300,"job_index":10,"waypoint_index":1}],"prev_leg_index":0,"next_leg_index":1},{"original_location":[10.983802751265776,48.25558745],"location":[10.983803,48.255587],"start_time":432,"duration":300,"actions":[{"index":2,"type":"job","start_time":432,"duration":300,"job_index":6,"waypoint_index":2}],"prev_leg_index":1,"next_leg_index":2},{"original_location":[10.984062746838354,48.25549435],"location":[10.984063,48.255494],"start_time":745,"duration":300,"actions":[{"index":3,"type":"job","start_time":745,"duration":300,"job_index":5,"waypoint_index":3}],"prev_leg_index":2,"next_leg_index":3},{"original_location":[10.984312143079265,48.25577875],"location":[10.984312,48.255779],"start_time":1054,"duration":300,"actions":[{"index":4,"type":"job","start_time":1054,"duration":300,"job_index":12,"waypoint_index":4}],"prev_leg_index":3,"next_leg_index":4},{"original_location":[10.984364769628737,48.25542385],"location":[10.984365,48.255424],"start_time":1362,"duration":300,"actions":[{"index":5,"type":"job","start_time":1362,"duration":300,"job_index":4,"waypoint_index":5}],"prev_leg_index":4,"next_leg_index":5},{"original_location":[10.896261152517647,48.33227795],"location":[10.896261,48.332278],"start_time":2752,"duration":0,"actions":[{"index":6,"type":"end","start_time":2752,"duration":0,"waypoint_index":6}],"prev_leg_index":5}]}},{"geometry":{"type":"MultiLineString","coordinates":[[[10.990289,48.267544],[10.968365,48.262043]],[[10.968365,48.262043],[10.896261,48.332278]]]},"type":"Feature","properties":{"agent_index":2,"time":2233,"start_time":0,"end_time":2233,"distance":22813,"legs":[{"time":531,"distance":2667,"from_waypoint_index":0,"to_waypoint_index":1,"steps":[{"from_index":0,"to_index":1,"time":531,"distance":2667}]},{"time":1402,"distance":20146,"from_waypoint_index":1,"to_waypoint_index":2,"steps":[{"from_index":0,"to_index":1,"time":1402,"distance":20146}]}],"mode":"drive","actions":[{"index":0,"type":"start","start_time":0,"duration":0,"waypoint_index":0},{"index":1,"type":"job","start_time":531,"duration":300,"job_index":3,"waypoint_index":1},{"index":2,"type":"end","start_time":2233,"duration":0,"waypoint_index":2}],"waypoints":[{"original_location":[10.990289270853658,48.2675442],"location":[10.990289,48.267544],"start_time":0,"duration":0,"actions":[{"index":0,"type":"start","start_time":0,"duration":0,"waypoint_index":0}],"next_leg_index":0},{"original_location":[10.968364837855287,48.262043399999996],"location":[10.968365,48.262043],"start_time":531,"duration":300,"actions":[{"index":1,"type":"job","start_time":531,"duration":300,"job_index":3,"waypoint_index":1}],"prev_leg_index":0,"next_leg_index":1},{"original_location":[10.896261152517647,48.33227795],"location":[10.896261,48.332278],"start_time":2233,"duration":0,"actions":[{"index":2,"type":"end","start_time":2233,"duration":0,"waypoint_index":2}],"prev_leg_index":1}]}},{"geometry":{"type":"MultiLineString","coordinates":[[[10.98728,48.267799],[10.983222,48.255775]],[[10.983222,48.255775],[10.983236,48.256069]],[[10.983236,48.256069],[10.982919,48.25585]],[[10.982919,48.25585],[10.983499,48.255697]],[[10.983499,48.255697],[10.980956,48.256226]],[[10.980956,48.256226],[10.981144,48.257297]],[[10.981144,48.257297],[10.980739,48.257487]],[[10.980739,48.257487],[10.896261,48.332278]]]},"type":"Feature","properties":{"agent_index":3,"time":3424,"start_time":0,"end_time":3424,"distance":21140,"legs":[{"time":170,"distance":1613,"from_waypoint_index":0,"to_waypoint_index":1,"steps":[{"from_index":0,"to_index":1,"time":170,"distance":1613}]},{"time":18,"distance":47,"from_waypoint_index":1,"to_waypoint_index":2,"steps":[{"from_index":0,"to_index":1,"time":18,"distance":47}]},{"time":7,"distance":19,"from_waypoint_index":2,"to_waypoint_index":3,"steps":[{"from_index":0,"to_index":1,"time":7,"distance":19}]},{"time":14,"distance":48,"from_waypoint_index":3,"to_waypoint_index":4,"steps":[{"from_index":0,"to_index":1,"time":14,"distance":48}]},{"time":47,"distance":258,"from_waypoint_index":4,"to_waypoint_index":5,"steps":[{"from_index":0,"to_index":1,"time":47,"distance":258}]},{"time":23,"distance":143,"from_waypoint_index":5,"to_waypoint_index":6,"steps":[{"from_index":0,"to_index":1,"time":23,"distance":143}]},{"time":16,"distance":47,"from_waypoint_index":6,"to_waypoint_index":7,"steps":[{"from_index":0,"to_index":1,"time":16,"distance":47}]},{"time":1029,"distance":18965,"from_waypoint_index":7,"to_waypoint_index":8,"steps":[{"from_index":0,"to_index":1,"time":1029,"distance":18965}]}],"mode":"drive","actions":[{"index":0,"type":"start","start_time":0,"duration":0,"waypoint_index":0},{"index":1,"type":"job","start_time":170,"duration":300,"job_index":7,"waypoint_index":1},{"index":2,"type":"job","start_time":488,"duration":300,"job_index":11,"waypoint_index":2},{"index":3,"type":"job","start_time":795,"duration":300,"job_index":9,"waypoint_index":3},{"index":4,"type":"job","start_time":1109,"duration":300,"job_index":8,"waypoint_index":4},{"index":5,"type":"job","start_time":1456,"duration":300,"job_index":16,"waypoint_index":5},{"index":6,"type":"job","start_time":1779,"duration":300,"job_index":13,"waypoint_index":6},{"index":7,"type":"job","start_time":2095,"duration":300,"job_index":14,"waypoint_index":7},{"index":8,"type":"end","start_time":3424,"duration":0,"waypoint_index":8}],"waypoints":[{"original_location":[10.987279662433057,48.2677992],"location":[10.98728,48.267799],"start_time":0,"duration":0,"actions":[{"index":0,"type":"start","start_time":0,"duration":0,"waypoint_index":0}],"next_leg_index":0},{"original_location":[10.983222005227521,48.255775],"location":[10.983222,48.255775],"start_time":170,"duration":300,"actions":[{"index":1,"type":"job","start_time":170,"duration":300,"job_index":7,"waypoint_index":1}],"prev_leg_index":0,"next_leg_index":1},{"original_location":[10.983236456481574,48.2560687],"location":[10.983236,48.256069],"start_time":488,"duration":300,"actions":[{"index":2,"type":"job","start_time":488,"duration":300,"job_index":11,"waypoint_index":2}],"prev_leg_index":1,"next_leg_index":2},{"original_location":[10.982919152872745,48.2558497],"location":[10.982919,48.25585],"start_time":795,"duration":300,"actions":[{"index":3,"type":"job","start_time":795,"duration":300,"job_index":9,"waypoint_index":3}],"prev_leg_index":2,"next_leg_index":3},{"original_location":[10.983499356818182,48.25569725],"location":[10.983499,48.255697],"start_time":1109,"duration":300,"actions":[{"index":4,"type":"job","start_time":1109,"duration":300,"job_index":8,"waypoint_index":4}],"prev_leg_index":3,"next_leg_index":4},{"original_location":[10.980955539642784,48.2562265],"location":[10.980956,48.256226],"start_time":1456,"duration":300,"actions":[{"index":5,"type":"job","start_time":1456,"duration":300,"job_index":16,"waypoint_index":5}],"prev_leg_index":4,"next_leg_index":5},{"original_location":[10.981143603167904,48.257296600000004],"location":[10.981144,48.257297],"start_time":1779,"duration":300,"actions":[{"index":6,"type":"job","start_time":1779,"duration":300,"job_index":13,"waypoint_index":6}],"prev_leg_index":5,"next_leg_index":6},{"original_location":[10.9807393,48.25748695],"location":[10.980739,48.257487],"start_time":2095,"duration":300,"actions":[{"index":7,"type":"job","start_time":2095,"duration":300,"job_index":14,"waypoint_index":7}],"prev_leg_index":6,"next_leg_index":7},{"original_location":[10.896261152517647,48.33227795],"location":[10.896261,48.332278],"start_time":3424,"duration":0,"actions":[{"index":8,"type":"end","start_time":3424,"duration":0,"waypoint_index":8}],"prev_leg_index":7}]}},{"geometry":{"type":"MultiLineString","coordinates":[[[10.983011,48.262754],[10.981209,48.257866]],[[10.981209,48.257866],[10.979089,48.257264]],[[10.979089,48.257264],[10.978801,48.257238]],[[10.978801,48.257238],[10.896261,48.332278]]]},"type":"Feature","properties":{"agent_index":4,"time":2365,"start_time":0,"end_time":2365,"distance":19873,"legs":[{"time":95,"distance":756,"from_waypoint_index":0,"to_waypoint_index":1,"steps":[{"from_index":0,"to_index":1,"time":95,"distance":756}]},{"time":33,"distance":220,"from_waypoint_index":1,"to_waypoint_index":2,"steps":[{"from_index":0,"to_index":1,"time":33,"distance":220}]},{"time":17,"distance":51,"from_waypoint_index":2,"to_waypoint_index":3,"steps":[{"from_index":0,"to_index":1,"time":17,"distance":51}]},{"time":1020,"distance":18846,"from_waypoint_index":3,"to_waypoint_index":4,"steps":[{"from_index":0,"to_index":1,"time":1020,"distance":18846}]}],"mode":"drive","actions":[{"index":0,"type":"start","start_time":0,"duration":0,"waypoint_index":0},{"index":1,"type":"job","start_time":95,"duration":300,"job_index":15,"waypoint_index":1},{"index":2,"type":"job","start_time":428,"duration":300,"job_index":18,"waypoint_index":2},{"index":3,"type":"job","start_time":728,"duration":300,"job_index":17,"waypoint_index":2},{"index":4,"type":"job","start_time":1045,"duration":300,"job_index":19,"waypoint_index":3},{"index":5,"type":"end","start_time":2365,"duration":0,"waypoint_index":4}],"waypoints":[{"original_location":[10.983010635351173,48.262753950000004],"location":[10.983011,48.262754],"start_time":0,"duration":0,"actions":[{"index":0,"type":"start","start_time":0,"duration":0,"waypoint_index":0}],"next_leg_index":0},{"original_location":[10.981209348235852,48.25786594111741],"location":[10.981209,48.257866],"start_time":95,"duration":300,"actions":[{"index":1,"type":"job","start_time":95,"duration":300,"job_index":15,"waypoint_index":1}],"prev_leg_index":0,"next_leg_index":1},{"original_location":[10.979089323915998,48.25726365],"location":[10.979089,48.257264],"start_time":428,"duration":600,"actions":[{"index":2,"type":"job","start_time":428,"duration":300,"job_index":18,"waypoint_index":2},{"index":3,"type":"job","start_time":728,"duration":300,"job_index":17,"waypoint_index":2}],"prev_leg_index":1,"next_leg_index":2},{"original_location":[10.978800955841443,48.25723825],"location":[10.978801,48.257238],"start_time":1045,"duration":300,"actions":[{"index":4,"type":"job","start_time":1045,"duration":300,"job_index":19,"waypoint_index":3}],"prev_leg_index":2,"next_leg_index":3},{"original_location":[10.896261152517647,48.33227795],"location":[10.896261,48.332278],"start_time":2365,"duration":0,"actions":[{"index":5,"type":"end","start_time":2365,"duration":0,"waypoint_index":4}],"prev_leg_index":3}]}}]}`)

var req = p.Request{
	Mode: "drive",
	Agents: []p.Agent{
		{
			Start:          p.LonLat{10.985736727197894, 48.2649593},
			End:            p.LonLat{10.896261152517647, 48.33227795},
			PickupCapacity: 4,
		},
		{
			Start:          p.LonLat{10.984995162319564, 48.264409},
			End:            p.LonLat{10.896261152517647, 48.33227795},
			PickupCapacity: 7,
		},
		{
			Start:          p.LonLat{10.990289270853658, 48.2675442},
			End:            p.LonLat{10.896261152517647, 48.33227795},
			PickupCapacity: 4,
		},
		{
			Start:          p.LonLat{10.987279662433057, 48.2677992},
			End:            p.LonLat{10.896261152517647, 48.33227795},
			PickupCapacity: 7,
		},
		{
			Start:          p.LonLat{10.983010635351173, 48.262753950000004},
			End:            p.LonLat{10.896261152517647, 48.33227795},
			PickupCapacity: 4,
		},
	},
	Jobs: []p.Job{
		{
			LonLat:       p.LonLat{10.98698105, 48.25615875},
			Duration:     300,
			PickupAmount: 1,
		},
		{
			LonLat:       p.LonLat{10.9845547, 48.26311145},
			Duration:     300,
			PickupAmount: 1,
		},
		{
			LonLat:       p.LonLat{10.984630924828402, 48.263248250000004},
			Duration:     300,
			PickupAmount: 2,
		},
		{
			LonLat:       p.LonLat{10.968364837855287, 48.262043399999996},
			Duration:     300,
			PickupAmount: 1,
		},
		{
			LonLat:       p.LonLat{10.984364769628737, 48.25542385},
			Duration:     300,
			PickupAmount: 1,
		},
		{
			LonLat:       p.LonLat{10.984062746838354, 48.25549435},
			Duration:     300,
			PickupAmount: 1,
		},
		{
			LonLat:       p.LonLat{10.983802751265776, 48.25558745},
			Duration:     300,
			PickupAmount: 2,
		},
		{
			LonLat:       p.LonLat{10.983222005227521, 48.255775},
			Duration:     300,
			PickupAmount: 1,
		},
		{
			LonLat:       p.LonLat{10.983499356818182, 48.25569725},
			Duration:     300,
			PickupAmount: 1,
		},
		{
			LonLat:       p.LonLat{10.982919152872745, 48.2558497},
			Duration:     300,
			PickupAmount: 1,
		},
		{
			LonLat:       p.LonLat{10.983681544239769, 48.25621035},
			Duration:     300,
			PickupAmount: 2,
		},
		{
			LonLat:       p.LonLat{10.983236456481574, 48.2560687},
			Duration:     300,
			PickupAmount: 1,
		},
		{
			LonLat:       p.LonLat{10.984312143079265, 48.25577875},
			Duration:     300,
			PickupAmount: 1,
		},
		{
			LonLat:       p.LonLat{10.981143603167904, 48.257296600000004},
			Duration:     300,
			PickupAmount: 1,
		},
		{
			LonLat:       p.LonLat{10.9807393, 48.25748695},
			Duration:     300,
			PickupAmount: 1,
		},
		{
			LonLat:       p.LonLat{10.981209348235852, 48.25786594111741},
			Duration:     300,
			PickupAmount: 1,
		},
		{
			LonLat:       p.LonLat{10.980955539642784, 48.2562265},
			Duration:     300,
			PickupAmount: 1,
		},
		{
			LonLat:       p.LonLat{10.979089323915998, 48.25726365},
			Duration:     300,
			PickupAmount: 1,
		},
		{
			LonLat:       p.LonLat{10.979089323915998, 48.25726365},
			Duration:     300,
			PickupAmount: 1,
		},
		{
			LonLat:       p.LonLat{10.978800955841443, 48.25723825},
			Duration:     300,
			PickupAmount: 1,
		},
	},
}

var res = p.Plan{
	Type: "FeatureCollection",
	Properties: p.Properties{
		Mode: "drive",
		Params: p.Params{
			Mode: "drive",
			Agents: []p.Agent{
				{
					Start:          p.LonLat{10.985736727197894, 48.2649593},
					End:            p.LonLat{10.896261152517647, 48.33227795},
					PickupCapacity: 4,
				},
				{
					Start:          p.LonLat{10.984995162319564, 48.264409},
					End:            p.LonLat{10.896261152517647, 48.33227795},
					PickupCapacity: 7,
				},
				{
					Start:          p.LonLat{10.990289270853658, 48.2675442},
					End:            p.LonLat{10.896261152517647, 48.33227795},
					PickupCapacity: 4,
				},
				{
					Start:          p.LonLat{10.987279662433057, 48.2677992},
					End:            p.LonLat{10.896261152517647, 48.33227795},
					PickupCapacity: 7,
				},
				{
					Start:          p.LonLat{10.983010635351173, 48.262753950000004},
					End:            p.LonLat{10.896261152517647, 48.33227795},
					PickupCapacity: 4,
				},
			},
			Jobs: []p.Job{
				{
					LonLat:       p.LonLat{10.98698105, 48.25615875},
					Duration:     300,
					PickupAmount: 1,
				},
				{
					LonLat:       p.LonLat{10.9845547, 48.26311145},
					Duration:     300,
					PickupAmount: 1,
				},
				{
					LonLat:       p.LonLat{10.984630924828402, 48.263248250000004},
					Duration:     300,
					PickupAmount: 2,
				},
				{
					LonLat:       p.LonLat{10.968364837855287, 48.262043399999996},
					Duration:     300,
					PickupAmount: 1,
				},
				{
					LonLat:       p.LonLat{10.984364769628737, 48.25542385},
					Duration:     300,
					PickupAmount: 1,
				},
				{
					LonLat:       p.LonLat{10.984062746838354, 48.25549435},
					Duration:     300,
					PickupAmount: 1,
				},
				{
					LonLat:       p.LonLat{10.983802751265776, 48.25558745},
					Duration:     300,
					PickupAmount: 2,
				},
				{
					LonLat:       p.LonLat{10.983222005227521, 48.255775},
					Duration:     300,
					PickupAmount: 1,
				},
				{
					LonLat:       p.LonLat{10.983499356818182, 48.25569725},
					Duration:     300,
					PickupAmount: 1,
				},
				{
					LonLat:       p.LonLat{10.982919152872745, 48.2558497},
					Duration:     300,
					PickupAmount: 1,
				},
				{
					LonLat:       p.LonLat{10.983681544239769, 48.25621035},
					Duration:     300,
					PickupAmount: 2,
				},
				{
					LonLat:       p.LonLat{10.983236456481574, 48.2560687},
					Duration:     300,
					PickupAmount: 1,
				},
				{
					LonLat:       p.LonLat{10.984312143079265, 48.25577875},
					Duration:     300,
					PickupAmount: 1,
				},
				{
					LonLat:       p.LonLat{10.981143603167904, 48.257296600000004},
					Duration:     300,
					PickupAmount: 1,
				},
				{
					LonLat:       p.LonLat{10.9807393, 48.25748695},
					Duration:     300,
					PickupAmount: 1,
				},
				{
					LonLat:       p.LonLat{10.981209348235852, 48.25786594111741},
					Duration:     300,
					PickupAmount: 1,
				},
				{
					LonLat:       p.LonLat{10.980955539642784, 48.2562265},
					Duration:     300,
					PickupAmount: 1,
				},
				{
					LonLat:       p.LonLat{10.979089323915998, 48.25726365},
					Duration:     300,
					PickupAmount: 1,
				},
				{
					LonLat:       p.LonLat{10.979089323915998, 48.25726365},
					Duration:     300,
					PickupAmount: 1,
				},
				{
					LonLat:       p.LonLat{10.978800955841443, 48.25723825},
					Duration:     300,
					PickupAmount: 1,
				},
			},
		},
	},
	Features: []p.Feature{
		{
			Geometry: p.Geometry{
				Type: "MultiLineString",
				Coordinates: [][]p.LonLat{
					{
						{10.985737, 48.264959},
						{10.984631, 48.263248},
					},
					{
						{10.984631, 48.263248},
						{10.984555, 48.263111},
					},
					{
						{10.984555, 48.263111},
						{10.986981, 48.256159},
					},
					{
						{10.986981, 48.256159},
						{10.896261, 48.332278},
					},
				},
			},
			Type: "Feature",
			Properties: p.AgentPlan{
				AgentIndex: 0,
				Time:       2066,
				StartTime:  0,
				EndTime:    2066,
				Distance:   20977,
				Legs: []p.Leg{
					{
						Time:              42,
						Distance:          290,
						FromWaypointIndex: 0,
						ToWaypointIndex:   1,
						Steps: []p.Step{
							{
								FromIndex: 0,
								ToIndex:   1,
								Time:      42,
								Distance:  290,
							},
						},
					},
					{
						Time:              2,
						Distance:          5,
						FromWaypointIndex: 1,
						ToWaypointIndex:   2,
						Steps: []p.Step{
							{
								FromIndex: 0,
								ToIndex:   1,
								Time:      2,
								Distance:  5,
							},
						},
					},
					{
						Time:              80,
						Distance:          906,
						FromWaypointIndex: 2,
						ToWaypointIndex:   3,
						Steps: []p.Step{
							{
								FromIndex: 0,
								ToIndex:   1,
								Time:      80,
								Distance:  906,
							},
						},
					},
					{
						Time:              1042,
						Distance:          19776,
						FromWaypointIndex: 3,
						ToWaypointIndex:   4,
						Steps: []p.Step{
							{
								FromIndex: 0,
								ToIndex:   1,
								Time:      1042,
								Distance:  19776,
							},
						},
					},
				},
				Mode: "drive",
				Actions: []p.Action{
					{
						Index:         0,
						Type:          "start",
						StartTime:     0,
						Duration:      0,
						WaypointIndex: 0,
					},
					{
						Index:         1,
						Type:          "job",
						StartTime:     42,
						Duration:      300,
						JobIndex:      2,
						WaypointIndex: 1,
					},
					{
						Index:         2,
						Type:          "job",
						StartTime:     344,
						Duration:      300,
						JobIndex:      1,
						WaypointIndex: 2,
					},
					{
						Index:         3,
						Type:          "job",
						StartTime:     724,
						Duration:      300,
						JobIndex:      0,
						WaypointIndex: 3,
					},
					{
						Index:         4,
						Type:          "end",
						StartTime:     2066,
						Duration:      0,
						WaypointIndex: 4,
					},
				},
				Waypoints: []p.Waypoint{
					{
						OriginalLoc: p.LonLat{10.985736727197894, 48.2649593},
						Location:    p.LonLat{10.985737, 48.264959},
						StartTime:   0,
						Duration:    0,
						Actions: []p.Action{
							{
								Index:         0,
								Type:          "start",
								StartTime:     0,
								Duration:      0,
								WaypointIndex: 0,
							},
						},
						NextLegIndex: 0,
					},
					{
						OriginalLoc: p.LonLat{10.984630924828402, 48.263248250000004},
						Location: p.LonLat{10.984631,
							48.263248},
						StartTime: 42,
						Duration:  300,
						Actions: []p.Action{
							{
								Index:         1,
								JobIndex:      2,
								Type:          "job",
								StartTime:     42,
								Duration:      300,
								WaypointIndex: 1,
							},
						},
						PrevLegIndex: 0,
						NextLegIndex: 1,
					},
					{
						OriginalLoc: p.LonLat{10.9845547, 48.26311145},
						Location:    p.LonLat{10.984555, 48.263111},
						StartTime:   344,
						Duration:    300,
						Actions: []p.Action{
							{
								Index:         2,
								Type:          "job",
								StartTime:     344,
								Duration:      300,
								JobIndex:      1,
								WaypointIndex: 2,
							},
						},
						PrevLegIndex: 1,
						NextLegIndex: 2,
					},
					{
						OriginalLoc: p.LonLat{10.98698105, 48.25615875},
						Location:    p.LonLat{10.986981, 48.256159},
						StartTime:   724,
						Duration:    300,
						Actions: []p.Action{
							{
								Index:         3,
								Type:          "job",
								StartTime:     724,
								Duration:      300,
								JobIndex:      0,
								WaypointIndex: 3,
							},
						},
						PrevLegIndex: 2,
						NextLegIndex: 3,
					},
					{
						OriginalLoc: p.LonLat{10.896261152517647,
							48.33227795},
						Location: p.LonLat{10.896261,
							48.332278},
						StartTime: 2066,
						Duration:  0,
						Actions: []p.Action{
							{
								Index:         4,
								Type:          "end",
								StartTime:     2066,
								Duration:      0,
								WaypointIndex: 4,
							},
						},
						PrevLegIndex: 3,
					},
				},
			},
		},
	},
}

const simpleJSON = `{
    "mode": "drive",
    "agents": [
        {
            "start_location": [
                106.72683176435947,
                -6.28476945
            ],
            "end_location": [
                106.7457861328624,
                -6.29673725
            ]
        }
    ],
    "jobs": [
        {
            "id": "order-xyz",
            "location": [
                106.7213665,
                -6.2895597
            ],
            "duration": 300
        },
        {
            "id": "order-abc",
            "location": [
                106.7157832,
                -6.2853618
            ],
            "duration": 300
        }
    ]
}`

var agent1 = p.Agent{
	Start: p.LonLat{106.72683176435947, -6.28476945},
	End:   p.LonLat{106.7457861328624, -6.29673725},
}

var customers = []p.Job{
	{
		ID:       "order-xyz",
		LonLat:   p.LonLat{106.7213665, -6.2895597},
		Duration: 300,
	},
	{
		ID:       "order-abc",
		LonLat:   p.LonLat{106.7157832, -6.2853618},
		Duration: 300,
	},
}

var simpleRequest = p.Request{
	Mode:   "drive",
	Agents: []p.Agent{agent1},
	Jobs:   customers,
}
