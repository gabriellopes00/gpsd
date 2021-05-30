# position

- latitude `float64`
- longitude `float64`

# tasks

### request

- receive victim location from its cellphone via http request POST :: /api/help (body: current position, header authToken)
- validate her credentials

### get helper positions

- send a signal to get ALL registered helpers position, via websocket or something like this
- receive ALL helpers position

### select useful helpers

- cache them positions per max of 10min
- get only helpers in victim radius from the cache, maybe 10km (if there is no one in this radius, send to nearest ones outside the radius, and tell to the victim that maybe the help can get late)
- if the data from the cache had been gotten right now, use them, else, use them to select the helpers, but send other request to get more accurate positions, and cache this newer positions instead of the the older ones

### calculate distance and path until the victim

- calculate distance between selected helpers and the victim
- get path from the selected helpers to the victim (paths, time until the victim, means of transport and them time until the victim...)

### send call

- send call only to nearest 10 helpers, and maybe a simple notification to the most distant
- if no one accept the request, send to distant ones

## rules

- to run in parallel ==> send signal ---> receive location ---> cache them ---> get only ones in the radius ---> calculate the distance ---> get paths with google api ---> send calls and notifications
- if some one accept a call, make the others know that there is one going to help the victim, with them position and time until the victim

### calcs - run in parallel

- verify victim's 10km radius helpers
- get more accurate positions from them
- receive the coordinate from the helpers
- calculate the nearest
- make google maps api calculate the path from each helper to the victim (time, distance, vehicles vias...)
- send a call to the nearest ones

### mobile interaction

#### permissions

- run in background (only to receive the signals)
- gps access (in background only on signals recoil)
- run the app in lock screen

#### app

- a map with user current position
<!-- - navbar on bottom with a button with "search for victims" that will find for victims, including the most distant (nearest cities) -->
- perfil page
- recently data charts page

##### call page

- map showing the path form the helper to the victim
- the time until the victim with different means of transport
- the distance until the victim location
- show if there is some one going to help the victim (distance, travel time)
- a button to accept or not the call (if accept, set one more helper going to help, else close the app screen)

- denunciate misuse of the app --> disable misused account
