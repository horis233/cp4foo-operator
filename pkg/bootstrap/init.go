package bootstrap

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/ghodss/yaml"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/klog"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

// InitResources initialize resources at the bootstrap of operator
func InitResources(mgr manager.Manager) error {
	client := mgr.GetClient()
	reader := mgr.GetAPIReader()
	err := updateConfigFromYaml([]byte(operandConfig), client, reader)
	if err != nil {
		klog.Error(err, "Update OperandConfig error")
	}
	timeout := time.After(300 * time.Second)
	ticker := time.NewTicker(30 * time.Second)
	for {
		klog.Info("try to create IBM Common Services OperandConfig and OperandRegistry")
		select {
		case <-timeout:
			return fmt.Errorf("Timeout to create the ODLM resource")
		case <-ticker.C:
			// create OperandRequest
			err := createOrUpdateFromYaml([]byte(operandRequest), client, reader)
			if err != nil {
				klog.Error(err, "create OperandRequest error")
			} else {
				return nil
			}
		}
	}
}

func yamlToObject(yamlContent []byte) (*unstructured.Unstructured, error) {
	obj := &unstructured.Unstructured{}
	jsonSpec, err := yaml.YAMLToJSON(yamlContent)
	if err != nil {
		return nil, fmt.Errorf("could not convert yaml to json: %v", err)
	}

	if err := obj.UnmarshalJSON(jsonSpec); err != nil {
		return nil, fmt.Errorf("could not unmarshal resource: %v", err)
	}

	return obj, nil
}

func getObject(obj *unstructured.Unstructured, reader client.Reader) (*unstructured.Unstructured, error) {
	found := &unstructured.Unstructured{}
	found.SetGroupVersionKind(obj.GetObjectKind().GroupVersionKind())

	err := reader.Get(context.TODO(), types.NamespacedName{Name: obj.GetName(), Namespace: obj.GetNamespace()}, found)

	return found, err
}

func createObject(obj *unstructured.Unstructured, client client.Client) error {
	err := client.Create(context.TODO(), obj)
	if err != nil && !errors.IsAlreadyExists(err) {
		return fmt.Errorf("could not Create resource: %v", err)
	}

	return nil
}

func updateObject(obj *unstructured.Unstructured, client client.Client) error {
	if err := client.Update(context.TODO(), obj); err != nil {
		return fmt.Errorf("could not update resource: %v", err)
	}

	return nil
}

func createOrUpdateFromYaml(yamlContent []byte, client client.Client, reader client.Reader) error {
	obj, err := yamlToObject(yamlContent)
	if err != nil {
		return err
	}

	if obj.GetNamespace() == "" {
		obj.SetNamespace(os.Getenv("WATCH_NAMESPACE"))
	}

	objInCluster, err := getObject(obj, reader)
	if errors.IsNotFound(err) {
		return createObject(obj, client)
	} else if err != nil {
		return err
	}

	version, _ := strconv.Atoi(obj.GetAnnotations()["version"])
	versionInCluster, _ := strconv.Atoi(objInCluster.GetAnnotations()["version"])

	if version > versionInCluster {
		return updateObject(obj, client)
	}

	return nil
}

func updateConfigFromYaml(yamlContent []byte, client client.Client, reader client.Reader) error {
	obj, err := yamlToObject(yamlContent)
	if err != nil {
		return err
	}

	err = wait.PollImmediate(time.Second*10, time.Second*300, func() (bool, error) {
		_, err := getObject(obj, reader)
		if errors.IsNotFound(err) {
			return false, nil
		} else if err != nil {
			return true, err
		}
		return true, updateObject(obj, client)
	})

	return err
}