# CP4Foo

This is a Foo CloudPak that used to demo how cloudpak integrate with common service.

## Integration Parts

1. Operator dependency in CSV: https://github.ibm.com/zhiwei/cp4foo-operator/blob/master/deploy/olm-catalog/cp4foo-operator/0.0.1/cp4foo-operator.v0.0.1.clusterserviceversion.yaml#L30-L35

    ```
    spec:
      customresourcedefinitions:
        required:
        - description: CommonService is the Schema for the commonservices API
          displayName: CommonService
          kind: CommonService
          name: commonservices.operator.ibm.com
          version: v1alpha1
    ```

2. Create the IBM Common Services OperandRequest during start of the operator

    https://github.ibm.com/zhiwei/cp4foo-operator/blob/master/cmd/manager/main.go#L126-L131


3. Handle future updates

    https://github.ibm.com/zhiwei/cp4foo-operator/blob/master/pkg/bootstrap/constant.go#L8-L9

    https://github.ibm.com/zhiwei/cp4foo-operator/blob/master/pkg/bootstrap/init.go#L100-L105


## How to use

### 1.Create IBM Common Services OperatorSource

```yaml
apiVersion: operators.coreos.com/v1
kind: OperatorSource
metadata:
  name: opencloud-operators
  namespace: openshift-marketplace
spec:
  authorizationToken: {}
  displayName: IBMCS Operators
  endpoint: https://quay.io/cnr
  publisher: IBM
  registryNamespace: opencloudio
  type: appregistry
```

### 2.Create CP4Foo OperatorSource

```yaml
apiVersion: operators.coreos.com/v1
kind: OperatorSource
metadata:
  name: cp4foo-operators
  namespace: openshift-marketplace
spec:
  authorizationToken: {}
  displayName: CP4Foo Operators
  endpoint: https://quay.io/cnr
  publisher: IBM
  registryNamespace: ibmcloud
  type: appregistry
```

### 3.Create a namespace cp4foo-operator

Create a project/namespace `cp4foo-operator` in OpenShift cluster.


### 4.Install CP4Foo operator

Search CP4Foo operator in OperatorHub and install CP4Foo operator into `cp4foo-operator` namespace.


### 5.Done

Waiting for a few minutes, the CP4Foo operator and IBM Common Services will be installed.

NOTE: In this demo, CP4Foo only depends on IBM Licensing, so the cluster will only have Licensing and its dependencies installed.
