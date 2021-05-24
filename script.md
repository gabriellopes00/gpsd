# position

- latitude `float64`
- longitude `float64`

# tasks

- receive victim location from its cellphone
- validate her credentials
- send a signal to get victim city's helpers current position
- get helpers position
- calculate the nearest ones
- get path from them to the victim
- send call to the helpers

## rules

- to run in parallel ==> send signal ---> receive location ---> calculate nearest ones ---> get paths with google api ---> send calls
- if some one accept a call, make the others know that there is one going to help the victim, with its position and time to the victim

### calcs - run in parallel

- verify for victim city's helpers
- send a signal via web socket or something like this
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
