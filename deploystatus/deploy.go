package deploymentStatus

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var err error
var config *rest.Config
var kubeconfig *string

func Testdeploy(deployna string) string {

	map_variable := make(map[string]int32)
	map_variable2 := make(map[string]int32)

	// 创建 clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	listOption := metav1.ListOptions{}

	// 调用 List 接口传入namespace字段，获取 Deployment 列表
	deploymentList, err := clientset.AppsV1().Deployments(deployna).List(context.Background(), listOption)
	if err != nil {
		fmt.Println("Err:", err)
		return "false"
	}

	// 遍历 Deployment 列表
	for _, deployment := range deploymentList.Items {
		//获取期望个数
		map_variable[deployment.ObjectMeta.Name] = int32(deployment.Status.Replicas)
		//获取健康状态的个数
		map_variable2[deployment.ObjectMeta.Name] = int32(deployment.Status.AvailableReplicas)
	}
	//判断期望个数与健康个数是否相等
	reflect.DeepEqual(map_variable, map_variable2)

	if reflect.DeepEqual(map_variable, map_variable2) {
		//如果相同返回true
		return "true"
	} else {
		//如果相同返回flase
		return "flase"

	}

}

func homeDir() string {
	//获取家目录
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}

func init() {

	if home := homeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// 使用 ServiceAccount 创建集群配置（InCluster模式）
	if config, err = rest.InClusterConfig(); err != nil {
		// 使用 KubeConfig 文件创建集群配置
		if config, err = clientcmd.BuildConfigFromFlags("", *kubeconfig); err != nil {
			panic(err.Error())
		}
	}
}
