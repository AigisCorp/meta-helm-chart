# Meta Helm Chart

Meta Helm Chart is a versatile Helm chart that simplifies the deployment and configuration of Kubernetes applications using only a `values.yaml` file. It is designed to be user-friendly and suitable for a wide range of use cases, making it accessible to everyone.

## Table of Contents

- [Meta Helm Chart](#meta-helm-chart)
  - [Introduction](#introduction)
  - [Features](#features)
  - [Getting Started](#getting-started)
  - [Usage](#usage)
  - [Contributing](#contributing)
  - [License](#license)

## Introduction

Helm is a popular package manager for Kubernetes that allows you to define, install, and upgrade even the most complex Kubernetes applications. While Helm provides a powerful templating mechanism, creating Helm charts can still be complex and require expertise.

Meta Helm Chart simplifies this process by allowing you to define your application's configuration using a single `values.yaml` file. You can configure various aspects of your application without the need to create or modify Helm templates. This makes it easy for both beginners and experienced users to deploy Kubernetes applications.

## Features

- **Simplified Configuration**: Define your application's configuration using a `values.yaml` file without dealing with Helm templates.

- **Customizable**: Easily customize various aspects of your application, such as image versions, resources, environment variables, and more, all in the `values.yaml` file.

- **Compatibility**: Works with Helm 3 and can be used with any Kubernetes cluster.

- **Extensible**: You can still leverage Helm's templating capabilities if needed by including custom templates alongside the `values.yaml` file.

## Getting Started

1. **Installation**: Ensure you have Helm 3 installed on your system. If not, you can follow the [Helm installation guide](https://helm.sh/docs/intro/install/) to get it set up.

2. **Create your own values.yaml from [this file](example/values.yaml)**:

3. **Template with your custom values.yaml**:

   ```bash
   helm template meta-helm-chart oci://registry-1.docker.io/aigiscorp/meta-helm-chart --values yourownvalues.yaml
   ```

4. **Install/Upgrade with your custom values.yaml**:

   ```bash
   helm install meta-helm-chart oci://registry-1.docker.io/aigiscorp/meta-helm-chart --values yourownvalues.yaml
   ```

   ```bash
   helm upgrade meta-helm-chart oci://registry-1.docker.io/aigiscorp/meta-helm-chart --values yourownvalues.yaml
   ```

## Usage

The primary usage of the Meta Helm Chart is to define and manage your application's configuration in the `values.yaml` file. You can find a sample `values.yaml` file in this repository as a starting point. Modify it according to your application's requirements.

For advanced users who want to leverage Helm's templating features, you can include custom templates in the Helm chart alongside the `values.yaml` file. These templates can be referenced in your `values.yaml` file, giving you the flexibility to combine the simplicity of `values.yaml` with the power of Helm templates.

For more detailed usage instructions and examples, refer to the documentation provided in this repository.

## Contributing

We welcome contributions from the community. If you have ideas for improvements, feature requests, or bug reports, please open an issue in this repository. If you'd like to contribute code, please fork the repository, make your changes, and submit a pull request.

Please review our [Contribution Guidelines](CONTRIBUTING) for more information on how to get involved.

## License

This project is licensed under the Apache 2.0 License. See [LICENSE](LICENSE) for more information.
