# Connect-gin (obsolete)

The backend web server for Connect. Built using Gin, Golang-JWT (and lib/pq as the database driver to connect to Postgres (Supabase)).

## Why did you abandon this project?

I really wanted to build something cool with Go (I was really intrigued by the language and its famed "speed"... and also the name `Go` sounds really cool :P). However, I decided to switch to [express.js](https://github.com/nithinrdy/connect-express) partway through:

1. I'm an absolute newbie when it comes to Go so I may be wrong here, but from what I've observed, the packages built for Go don't seem to be as mature as the ones built for Node.

    I really had to go digging to find something as simple (at least in case of Node) as a web socket library. (I still wanted to stick to Go and decided to switch over to express.js and socket.io to build the signalling server and continue with Gin for the rest of the backend (commit history for the `connect-express` repo shows these changes), but, as you can see, I eventually abandoned the project altogether, because...).

2. So Gin apparently doesn't serve the static build files of a React (CRA) project the right way?

    The way express handles this is pretty straightforward: you point to the server the directory that contains the static files, and then serve the `index.html` file on one of the routes and the scripts linked in it take over. The key point here is that the above approach takes care of the client side routing (at least when built using react-router).

    This isn't the case with Gin, which completely breaks the frontend routing unless the user begins from the base route. Something like `gin-contrib/static` didn't help either (assuming it should have; or maybe I used it incorrectly). The only way to get this to work at this point was to manually serve the the individual HTML files for every route (which is just as tedious as it sounds).

## The app is still fully-functional... ish

And so, there ended my 'Connect' journey with Gin. However, this project is still "fully-functional" though -- just set up a signalling server and link the react-app with this server for the basic routes (auth, etc.), and to the signalling server for websocket communication -- boom, you're 'CONNECT'-ed (pun intended).

You may also want to find a way to fix the main issue which forced me to give up on Gin and this project -- serving react's build files without having the client side routing completely break down (just use `express`? *wink wink* or run it in development mode forever because who cares about performance).

Still like Go, but will probably stick to Node for now.

## Installing

1. Clone the repo

2. Install all the dependencies using `go get .`.

3. Create a .env file in the root directory and add the following environment variables: `DB_PASSWORD`, `DB_HOST_URL`, `DB_PORT`. See the models folder for the database schemas when setting up your own database.

4. Run the server using `go run main.go`.
