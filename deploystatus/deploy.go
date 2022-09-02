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
	// var err error
	// var config *rest.Config
	// var kubeconfig *string

	// if home := homeDir(); home != "" {
	// 	kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	// } else {
	// 	kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	// }
	// flag.Parse()

	// // 使用 ServiceAccount 创建集群配置（InCluster模式）
	// if config, err = rest.InClusterConfig(); err != nil {
	// 	// 使用 KubeConfig 文件创建集群配置
	// 	if config, err = clientcmd.BuildConfigFromFlags("", *kubeconfig); err != nil {
	// 		panic(err.Error())
	// 	}
	// }

	// 创建 clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// 根据 deploymentGetName 参数是否为空来决定显示单个 Deployment 信息还是所有 Deployment 信息
	listOption := metav1.ListOptions{}
	// 如果指定了 Deployment Name，那么只获取单个 Deployment 的信息

	// 调用 List 接口获取 Deployment 列表
	deploymentList, err := clientset.AppsV1().Deployments(deployna).List(context.Background(), listOption)
	if err != nil {
		fmt.Println("Err:", err)
		return "false"
	}

	// 遍历 Deployment 列表
	for _, deployment := range deploymentList.Items {
		//获取期望状态以及健康状态
		map_variable[deployment.ObjectMeta.Name] = int32(deployment.Status.Replicas)
		map_variable2[deployment.ObjectMeta.Name] = int32(deployment.Status.AvailableReplicas)
	}
	reflect.DeepEqual(map_variable, map_variable2)
	// fmt.Println("第一个", map_variable)
	// fmt.Println("第二个", map_variable2)

	// fmt.Printf("%T", reflect.DeepEqual(map_variable, map_variable2))

	if reflect.DeepEqual(map_variable, map_variable2) {

		return "true"
	} else {
		return "flase"

	}

}

func homeDir() string {
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

//func init() {
//	file := "./" + "read_log" + ".log"
//	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
//	if err != nil {
//		panic(err)
//	}
//	log.SetOutput(logFile) // 将文件设置为log输出的文件
//	log.SetPrefix("[read_log]")
//	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
//	return
//}
