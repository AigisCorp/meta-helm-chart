# Getting started

**_Installation_**: This step is about ensuring that Helm 3 is installed on your system. Helm is a tool for managing Kubernetes applications, and Helm 3 is the version in use. If you don't have Helm 3 installed, you can follow the Helm installation guide to set it up. This guide will provide you with instructions for installing Helm on your system.

**_Create your own values.yaml_**: In this step, you are instructed to create a custom values.yaml file. This file is used to configure various aspects of your application deployment on Kubernetes. You can start with a sample values.yaml file provided at the link (example/values.yaml) and customize it to match your specific application's requirements. This customization is where you define things like the version of your application, resource allocation (CPU, memory), and environment variables.

**_Template with your custom values.yaml_**: This step involves using Helm to generate Kubernetes manifests based on your custom values.yaml file. It provides the following commands:

```sh
helm repo add meta-helm-chart https://aigiscorp.github.io/meta-helm-chart/
```

This adds a Helm repository called meta-helm-chart from a specific URL. Helm repositories are collections of charts (packages of pre-configured Kubernetes resources) that can be installed and managed.

```sh
helm repo update
```

This command updates the local list of Helm repositories to ensure you have the latest information.

```sh
helm template your-chart-name meta-helm-chart/meta-helm-chart --values yourownvalues.yaml
```

This command uses the Helm chart from the meta-helm-chart repository and your custom values.yaml file to generate Kubernetes manifests without actually installing anything. This is useful for previewing the configuration before applying it to your cluster.

**_Install/Upgrade with your custom values.yaml_**: These commands allow you to install or upgrade your application on your Kubernetes cluster using your custom values.yaml file:

```sh
helm install your-chart-name meta-helm-chart/meta-helm-chart --values yourownvalues.yaml
```

This command installs the Helm chart with the specified name (your-chart-name) using your custom values.yaml.

```sh
helm upgrade your-chart-name meta-helm-chart/meta-helm-chart --values yourownvalues.yaml
```

This command upgrades the Helm chart with the specified name, again using your custom values.yaml. This is useful when you need to make changes to your application's configuration.

**_Package with your custom values.yaml_**: This step is about packaging your custom configuration into a Helm chart that can be easily shared or distributed. It involves the following commands:

```sh
helm pull meta-helm-chart/meta-helm-chart --untar
```

This command downloads the Helm chart from the meta-helm-chart repository and extracts its contents into a directory.

```sh
cp yourownvalues.yaml meta-helm-chart/values.yaml
```

This copies your custom values.yaml into the Helm chart directory as values.yaml.

```sh
mv meta-helm-chart your-chart-name
```

This renames the directory to match the name you want for your Helm chart.

```sh
helm package your-chart-name
```

This packages the directory as a Helm chart with the name you specified. The resulting .tgz file can be easily shared or installed on other Kubernetes clusters.nstallation: This step is about ensuring that Helm 3 is installed on your system. Helm is a tool for managing Kubernetes applications, and Helm 3 is the version in use. If you don't have Helm 3 installed, you can follow the Helm installation guide to set it up. This guide will provide you with instructions for installing Helm on your system.
