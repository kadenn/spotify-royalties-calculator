## How to run this application locally on your machine?

Download the application:
```sh
git clone https://github.com/kadenn/spotify-royalties-calculator.git
```
**Note:** Client depends on the Server, do not run the Client without running the Server first.


### Server
**Note:** Make sure you replace the example.env file in the server directory with the one I shared with you.

```sh
cd spotify-royalties-calculator/server
```
```sh
go run .
```
To run tests:
```sh
go test .
```

### Client
```sh
cd spotify-royalties-calculator/client
```
```sh
yarn install
```
```sh
export REACT_APP_API_URL=http://localhost:8080 
```
```sh
yarn start
```
You can now view client in the browser: http://localhost:3000


