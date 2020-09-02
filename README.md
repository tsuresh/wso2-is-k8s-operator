
## WSO2 Identity Server - K8S Operator
The following CRD operator can be used to deploy WSO2 IS on your Kubernates Cluster. If you want to deploy the Identity Server via Helm operator please refer to the given link:  [https://github.com/wso2/kubernetes-is](https://github.com/wso2/kubernetes-is)

## Prerequisites[](https://sdk.operatorframework.io/docs/building-operators/golang/tutorial/#prerequisites)

-   Access to a Kubernetes v1.11.3+ cluster (v1.16.0+ if using  `apiextensions.k8s.io/v1`  CRDs).
-   User logged with admin permission. See  [how to grant yourself cluster-admin privileges or be logged in as admin](https://cloud.google.com/kubernetes-engine/docs/how-to/role-based-access-control#iam-rolebinding-bootstrap)
- [Homebrew](https://brew.sh/) installed
- Git command line installed and configured


## Installation [Dev]

Please follow the following instructions to install Operator-SDK in your development environment.

    brew install operator-sdk

Clone the repository by running the following command 

    git clone https://github.com/tsuresh/wso2-is-k8s-operator.git

Run the following command to install dependancies 

    make install

Navigate to the project directory 

    cd wso2-is-operator

Apply the CRD by running the following command 

    kubectl apply -f config/crd/bases/wso2.wso2.com_wso2is.yaml

You are free to change any configurations at **config/samples/wso2_v1_wso2is.yaml**
Once you do the config changes apply the config by running

    kubectl apply -f config/samples/wso2_v1_wso2is.yaml
    
Finally run the following command to run the operator in your cluster 

    make run

## Installation [Prod]

To be released
