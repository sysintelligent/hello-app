# Hello-App

This repository contains a simple application demonstrating how to interact with Kubernetes and expose metrics using Prometheus.

## Prerequisites

- Go installed on your machine.
- Access to a Kubernetes cluster.
- `config.json` file containing the application version.

## Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/yourusername/hello-app.git
    ```

2. Navigate to the project directory:

    ```bash
    cd hello-app
    ```

3. Compile the application:

    ```bash
    go build
    ```

## Usage

1. Ensure that your Kubernetes cluster is running.
2. Make sure you have a `config.json` file containing the application version in the root directory.
3. Execute the binary:

    ```bash
    ./hello-app
    ```

4. Access the application at `http://localhost:8080`.

## Metrics

- Metrics are exposed at `/metrics` endpoint.

## Testing

```bash
export KUBECONFIG=~/.kube/config
go run main.go
curl localhost:8080
```

## Docker Image Creation

To create a Docker image:

```bash
docker build -t hello-app .
docker tag hello-app sysintelligent/hello-app:2.0.1
docker push sysintelligent/hello-app:2.0.1
```

## Contributing

This project welcomes contributions. Feel free to fork and submit a pull request with your changes.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Thanks to the Kubernetes community for their excellent work.
- Special thanks to the Prometheus team for providing powerful monitoring capabilities.