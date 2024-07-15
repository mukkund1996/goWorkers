# goWorkers

`goWorkers` is intended to be a simulation of a worker-pool accepting map-reduce jobs. It comprises of the following elements:

- Web-server - based on GO gin framework.
- Worker pool - accepts map-reduce jobs
- Web-client - to monitor the workers

## Getting started

### Running the web-server

Run the following command to start the webserver:

```bash
go run webserver/server/main.go
```

### Running the web-client

Run the following command to set up the dev build of the web-client on port `5173`.

```bash
cd client && npm run dev
```

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, feel free to open an issue or create a pull request.

## License

This project is licensed under the MIT License. See the LICENSE file for details.

## Acknowledgements

Inspired by various worker pool implementations in Go.
Special thanks to the Go community for their valuable contributions and support.
