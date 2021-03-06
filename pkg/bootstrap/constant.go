package bootstrap

const operandRequest = `
apiVersion: operator.ibm.com/v1alpha1
kind: OperandRequest
metadata:
  name: cp4foo-common-service
  annotations:
    version: "1"
spec:
  requests:
  - registry: common-service
    registryNamespace: ibm-common-services
    operands:
    - name: ibm-licensing-operator
`

const operandConfig = `
apiVersion: operator.ibm.com/v1alpha1
kind: OperandConfig
metadata:
  name: common-service
  namespace: ibm-common-services
  annotations:
    version: "1"
spec:
  services:
  - name: ibm-metering-operator
    spec:
      metering: {}
      meteringUI: {}
      meteringReportServer: {}
      operandBindInfo: {}
      operandRequest: {}
  - name: ibm-licensing-operator
    spec:
      IBMLicensing:
        logLevel: DEBUG
      operandBindInfo: {}
      operandRequest: {}
  - name: ibm-mongodb-operator
    spec:
      mongoDB: {}
      operandRequest: {}
  - name: ibm-cert-manager-operator
    spec:
      certManager: {}
      issuer: {}
      certificate: {}
      clusterIssuer: {}
  - name: ibm-iam-operator
    spec:
      authentication: {}
      oidcclientwatcher: {}
      pap: {}
      policycontroller: {}
      policydecision: {}
      secretwatcher: {}
      securityonboarding: {}
      operandBindInfo: {}
      operandRequest: {}
  - name: ibm-healthcheck-operator
    spec:
      healthService: {}
  - name: ibm-commonui-operator
    spec:
      commonWebUI: {}
      operandRequest: {}
  - name: ibm-management-ingress-operator
    spec:
      managementIngress: {}
      operandBindInfo: {}
      operandRequest: {}
  - name: ibm-ingress-nginx-operator
    spec:
      nginxIngress: {}
  - name: ibm-auditlogging-operator
    spec:
      auditLogging: {}
      operandRequest: {}
  - name: ibm-catalog-ui-operator
    spec:
      catalogUI: {}
      operandRequest: {}
  - name: ibm-platform-api-operator
    spec:
      platformApi: {}
      operandRequest: {}
  - name: ibm-helm-api-operator
    spec:
      helmApi: {}
      operandRequest: {}
  - name: ibm-helm-repo-operator
    spec:
      helmRepo: {}
      operandRequest: {}
  - name: ibm-monitoring-exporters-operator
    spec:
      exporter: {}
      operandRequest: {}
  - name: ibm-monitoring-prometheusext-operator
    spec:
      prometheusExt: {}
      operandRequest: {}
  - name: ibm-monitoring-grafana-operator
    spec:
      grafana: {}
      operandRequest: {}
  - name: ibm-elastic-stack-operator
    spec:
      elasticStack: {}
      operandBindInfo: {}
      operandRequest: {}
`
