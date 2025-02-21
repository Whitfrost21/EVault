# EVault - E-Waste Scraper and Recycling System

EVault is an innovative e-waste management solution designed to accept and process recycling requests specifically for electronic vehicles. The system allows users to schedule waste collection requests, choose their location on a real-time map, and utilizes GraphHopper API to trace optimal routes, calculate time, and distance. The system uses Go for the backend with a Fyne-based GUI for easy user interaction.

!["Evault Dashboard"](https://github.com/Whitfrost21/EVault/blob/master/Screenshots/dashboard.jpg?raw=true)

## Features

- **E-Waste Collection Requests**: Submit the request with specified details.
  !["Add Requests"](https://github.com/Whitfrost21/EVault/blob/master/Screenshots/Reqform.jpg)
- **Real-Time Map Integration**: Can select location on real-time map, (i used osm with leaflet.js for this).
- **Route Calculation**: Uses GraphHopper API to trace the best routes and calculate the distance and time to the pickup location.
  !["Map"](https://github.com/Whitfrost21/EVault/blob/master/Screenshots/map.png)

- **Automatic Monitoring**: The system checks pending requests every 5 minutes to ensure smooth operation.
- **In-App Notifications**: Notifications are sent to the users when a collection request is completed.

  !["notifications"](https://github.com/Whitfrost21/EVault/blob/master/Screenshots/notifications.jpg)

- **Scalability**: Designed to scale to handle increasing numbers of users and requests efficiently.

## Technologies Used

- **Go (Golang)**: Core backend logic for processing requests and server-side operations.
- **Fyne Framework**: GUI framework for frontend development and user interface.
- **GraphHopper API**: For route tracing, distance, and time calculation for pickups.
- **OSM with leaflet.js**:For real-time mapview integration.

## Installation

To run this project locally, follow these steps:

### 1. Clone the Repository

```bash
git clone https://github.com/Whitfrost21/EVault.git
cd EVault
```

### 2. Install Dependencies

Install the required dependencies by running:

```bash
go mod tidy
```

### 3. Set Up GraphHopper API Key

Sign up for a GraphHopper account to obtain an API key. Then, set the API key as an environment variable:

```bash
export GRAPHHOPPER_API_KEY="your_api_key"
```

### 4. Run the Application

Start the EVault server with the following command:

```bash
go run main.go
```

> 🔔 **Note:** I have'nt define any quality determination and cost evaluation method just them manually ,may be thinking to improve them in future .

## Contributing

We welcome contributions! If you'd like to improve the project, feel free to fork the repository, create a new branch, and submit a pull request. Please ensure to follow the coding standards and write tests for new features.

## License

This project is licensed under the Apache License - see the [LICENSE](LICENSE) file for details.

## Acknowledgements

- [GraphHopper](https://www.graphhopper.com) for the route calculation and distance/time tracking API.
- [Fyne](https://fyne.io) for the GUI framework.
- [Go](https://golang.org) for building the backend logic.

## Contact

For any questions or feedback, please feel free to reach out to [zoreprajwal495@gmail.com].
