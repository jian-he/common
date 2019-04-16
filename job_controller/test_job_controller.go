package job_controller

import (
	commonv1 "github.com/kubeflow/common/operator/v1"
	testv1 "github.com/kubeflow/common/test_job/v1"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type TestJobController struct {
	JobController
	job            *testv1.TestJob
	createdService *corev1.Service
}

func (TestJobController) ControllerName() string {
	return "test-operator"
}

func (TestJobController) GetAPIGroupVersionKind() schema.GroupVersionKind {
	return testv1.SchemeGroupVersionKind
}

func (TestJobController) GetAPIGroupVersion() schema.GroupVersion {
	return testv1.SchemeGroupVersion
}

func (TestJobController) GetGroupNameLabelKey() string {
	return "group-name"
}

func (TestJobController) GetJobNameLabelKey() string {
	return "test-replica-type"
}

func (TestJobController) GetGroupNameLabelValue() string {
	return testv1.GroupName
}

func (TestJobController) GetReplicaTypeLabelKey() string {
	return "test-replica-type"
}

func (TestJobController) GetReplicaIndexLabelKey() string {
	return "test-replica-index"
}

func (TestJobController) GetJobRoleKey() string {
	return "test-job-role"
}

func (t *TestJobController) GetJobFromInformerCache(namespace, name string) (v1.Object, error) {
	return t.job, nil
}

func (t *TestJobController) GetJobFromAPIClient(namespace, name string) (v1.Object, error) {
	return t.job, nil
}

func (t *TestJobController) DeleteJob(job interface{}) error {
	t.job = nil
	return nil
}

func (t *TestJobController) UpdateJobStatus(job interface{}) error {
	*t.job = job.(testv1.TestJob)
	return nil
}

func (t *TestJobController) CreateNewService(job v1.Object, rtype commonv1.ReplicaType, spec *commonv1.ReplicaSpec, index string) error {
	labels := t.GenLabels(job.GetName())
	labels[t.GetReplicaTypeLabelKey()] = string(rtype)
	labels[t.GetReplicaIndexLabelKey()] = index

	service := &corev1.Service{
		Spec: corev1.ServiceSpec{
			ClusterIP: "None",
			Selector:  labels,
			Ports: []corev1.ServicePort{
				{
					Name:"tfjob-port",
					Port: 2222,
				},
			},
		},
	}
	t.createdService = service
	return nil
}

func (TestJobController) ReconcilePods(
	job v1.Object,
	pods []*corev1.Pod,
	rtype commonv1.ReplicaType,
	spec *commonv1.ReplicaSpec,
	rstatus map[string]corev1.PodPhase) error {
	panic("implement me")
}



