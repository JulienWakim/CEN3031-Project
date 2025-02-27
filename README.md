# PropertyProbe
PropertyProbe will be a powerful tool for real estate investors and professionals trying to find and analyze pre-foreclosure properties. It uses advanced algorithms to search through pre-foreclosure listings and compile all the relevant information on each property. With this tool, users can quickly and easily find properties that meet their investment criteria and make informed decisions about which properties to pursue. The program is user-friendly and easy to navigate, allowing users to filter properties by location.

## Requirements

To run and use PropertyProbe, you'll need:

- [Go version go1.19.5](https://golang.org/dl/) or later
- [Angular CLI version 15.1.2](https://cli.angular.io/) or later
- [Node.js](https://nodejs.org/) version 14.x or later

## Getting Started

1. Clone the MyApp repository to your local machine:

```
git clone https://github.com/AkshayAshok2/CEN3031-Project.git
```

2. Start Angular and Go servers: 

Option 1: If on Mac OS or Linux, run the shell script to install the required packages, run the Angular server, and start the Go server:

```
./run.sh
```

Option 2: If on Windows, navigate to the `client` directory, install the required packages, and start the Angular development server: 
```
cd client
npm install
npm run start &
```
In a separate terminal window, navigate to the `api` directory and start the Go server:
```
cd ../api
go run httpd/main.go
```

3. Open `http://localhost:4200` in your web browser to access PropertyProbe.


### Frontend Engineers:
Jacob Kanfer, Akshay Ashok

### Backend Engineers:
Carson Schmidt, Julien Wakim

