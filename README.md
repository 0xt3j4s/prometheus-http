# prometheus-http
- Built my first metric exporter using Go and Prometheus, as I wanted to learn Prometheus.
- This is a simple HTTP server that exposes metrics to count the number of requests made to the `/ping` endpoint.
- The server runs on port 8090 and the metrics can be viewed at `localhost:8090/metrics`.
- The metrics can be viewed in Prometheus at `localhost:9090/`.

## Pre-requisites
1. [Go](https://golang.org/doc/install)
2. [Prometheus](https://prometheus.io/docs/prometheus/latest/getting_started/) 

## Usage
1. Clone this repository, using the following command:
    ```bash
    git clone https://github.com/0xt3j4s/prometheus-http.git
    cd prometheus-http
    ```

2. Build the server
    ```bash
    go build server.go
    ```

3. Run the server
    ```bash
    ./server
    ```

4. To check if the server is running, open the following URL in your browser: http://localhost:8090/ping and then poggers.

5. To view the metrics, open the following URL in your browser: http://localhost:8090/metrics.

## Outputs
1. The following output is displayed when the server is running:
    - In the terminal:
    ```bash
    Server is listening on port 8090...
    ```
    - In the browser: 
    <br>
    ![Server is running](./images/server.png)
2. The following output is displayed when the metrics are viewed:
    - In the browser: 
    <br>
    ![Metrics](./images/metrics.png)
3. The following graph can be viewed in Prometheus:
    - Visit `localhost:9090/` and plot a graph after entering the expression `ping_request_count` over the last 1 minute (or any other duration): 
    <br>
    ![Graph](./images/prometheus.png)
 
## Contributing
Contributions are welcome! If you find any issues or have suggestions for improvement, please create an issue or submit a pull request.

1. Fork
2. Clone your fork
3. Create a new branch
4. Make your changes
5. Commit and push
6. Create a pull request

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.