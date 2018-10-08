package main

import (
	"encoding/json"
	"log"
	"os"
	"os/exec"
	"testing"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	apiclient "github.com/fogatlas/client-go/client"
	"github.com/fogatlas/client-go/client/operations"
	"github.com/fogatlas/client-go/models"
)

var client *apiclient.Fogatlas

func TestMain(m *testing.M) {
	//initialize etcd
	err := exec.Command("./populate.sh", "http://127.0.0.1:2379").Run()
	if err != nil {
		log.Printf("Unable to perform testing (%s)", err)
	}
	schemes := []string{"http"}
	transport := httptransport.New("127.0.0.1:8080", "/api/v2.0.0", schemes)
	client = apiclient.New(transport, strfmt.Default)
	retCode := m.Run()
	os.Exit(retCode)
}

func TestDeleteApplicationsID(t *testing.T) {
	params := operations.NewDeleteApplicationsIDParams()
	params.ID = "application1"
	_, err := client.Operations.DeleteApplicationsID(params)
	if err != nil {
		t.Errorf("TestDeleteApplicationsID failed(%s)", err)
	}
}

func TestDeleteMicroservicesID(t *testing.T) {
	params := operations.NewDeleteMicroservicesIDParams()
	params.ID = "microservice1"
	_, err := client.Operations.DeleteMicroservicesID(params)
	if err != nil {
		t.Errorf("TestDeleteMicroservicesID failed(%s)", err)
	}
}

func TestDeleteNodesID(t *testing.T) {
	params := operations.NewDeleteNodesIDParams()
	params.ID = "node11"
	_, err := client.Operations.DeleteNodesID(params)
	if err != nil {
		t.Errorf("TestDeleteNodesID failed(%s)", err)
	}
}

func TestDeleteRegionsID(t *testing.T) {
	params := operations.NewDeleteRegionsIDParams()
	params.ID = "reg1"
	_, err := client.Operations.DeleteRegionsID(params)
	if err != nil {
		t.Errorf("TestDeleteRegionsID failed(%s)", err)
	}
}

func TestDeleteRelationshipsID(t *testing.T) {
	params := operations.NewDeleteRelationshipsIDParams()
	params.ID = "rel1"
	_, err := client.Operations.DeleteRelationshipsID(params)
	if err != nil {
		t.Errorf("TestDeleteRelationshipsID failed(%s)", err)
	}
}

func TestDeleteExternalendpointsID(t *testing.T) {
	params := operations.NewDeleteExternalendpointsIDParams()
	params.ID = "thing1"
	_, err := client.Operations.DeleteExternalendpointsID(params)
	if err != nil {
		t.Errorf("TestDeleteExternalendpointsID failed(%s)", err)
	}
}

func TestDeleteDynamicnodesID(t *testing.T) {
	params := operations.NewDeleteDynamicnodesIDParams()
	params.ID = "dyn1"
	_, err := client.Operations.DeleteDynamicnodesID(params)
	if err != nil {
		t.Errorf("TestDeleteDynamicnodesID failed(%s)", err)
	}
}

func TestDeleteDeploymentsName(t *testing.T) {
	params := operations.NewDeleteDeploymentsNameParams()
	params.Name = "Face recognition appl1"
	_, err := client.Operations.DeleteDeploymentsName(params)
	if err != nil {
		t.Errorf("TestDeleteDeploymentsNameID failed(%s)", err)
	}
}

func TestGetApplications(t *testing.T) {
	params := operations.NewGetApplicationsParams()
	_, err := client.Operations.GetApplications(params)
	if err != nil {
		t.Errorf("TestGetApplications failed(%s)", err)
	}
}

func TestGetApplicationsID(t *testing.T) {
	params := operations.NewGetApplicationsIDParams()
	params.ID = "application2"
	_, err := client.Operations.GetApplicationsID(params)
	if err != nil {
		t.Errorf("TestGetApplicationsID failed(%s)", err)
	}
}

func TestGetDeployments(t *testing.T) {
	params := operations.NewGetDeploymentsParams()
	_, err := client.Operations.GetDeployments(params)
	if err != nil {
		t.Errorf("TestGetDeployments failed(%s)", err)
	}
}

func TestGetDeploymentsName(t *testing.T) {
	params := operations.NewGetDeploymentsNameParams()
	params.Name = "Face recognition appl"
	_, err := client.Operations.GetDeploymentsName(params)
	if err != nil {
		t.Errorf("TestGetDeploymentsName failed(%s)", err)
	}
}

func TestGetMicroservices(t *testing.T) {
	params := operations.NewGetMicroservicesParams()
	_, err := client.Operations.GetMicroservices(params)
	if err != nil {
		t.Errorf("TestGetMicroservices failed(%s)", err)
	}
}

func TestGetMicroservicesID(t *testing.T) {
	params := operations.NewGetMicroservicesIDParams()
	params.ID = "microservice2"
	_, err := client.Operations.GetMicroservicesID(params)
	if err != nil {
		t.Errorf("TestGetMicroservicesID failed(%s)", err)
	}
}

func TestGetNodes(t *testing.T) {
	params := operations.NewGetNodesParams()
	_, err := client.Operations.GetNodes(params)
	if err != nil {
		t.Errorf("TestGetNodes failed(%s)", err)
	}
}

func TestGetNodesID(t *testing.T) {
	params := operations.NewGetNodesIDParams()
	params.ID = "node12"
	_, err := client.Operations.GetNodesID(params)
	if err != nil {
		t.Errorf("TestGetNodesID failed(%s)", err)
	}
}

func TestGetRegions(t *testing.T) {
	params := operations.NewGetRegionsParams()
	_, err := client.Operations.GetRegions(params)
	if err != nil {
		t.Errorf("TestGetRegions failed(%s)", err)
	}
}

func TestGetRegionsID(t *testing.T) {
	params := operations.NewGetRegionsIDParams()
	params.ID = "reg3"
	_, err := client.Operations.GetRegionsID(params)
	if err != nil {
		t.Errorf("TestGetRegionsID failed(%s)", err)
	}
}

func TestGetRelationships(t *testing.T) {
	params := operations.NewGetRelationshipsParams()
	_, err := client.Operations.GetRelationships(params)
	if err != nil {
		t.Errorf("TestGetRelationships failed(%s)", err)
	}
}

func TestGetRelationshipsID(t *testing.T) {
	params := operations.NewGetRelationshipsIDParams()
	params.ID = "rel3"
	_, err := client.Operations.GetRelationshipsID(params)
	if err != nil {
		t.Errorf("TestGetRelationshipsID failed(%s)", err)
	}
}

func TestGetExternalendpoints(t *testing.T) {
	params := operations.NewGetExternalendpointsParams()
	_, err := client.Operations.GetExternalendpoints(params)
	if err != nil {
		t.Errorf("TestGetExternalendpoints failed(%s)", err)
	}
}

func TestGetExternalendpointsID(t *testing.T) {
	params := operations.NewGetExternalendpointsIDParams()
	params.ID = "thing2"
	_, err := client.Operations.GetExternalendpointsID(params)
	if err != nil {
		t.Errorf("TestGetExternalendpointsID failed(%s)", err)
	}
}

func TestGetDynamicnodes(t *testing.T) {
	params := operations.NewGetDynamicnodesParams()
	_, err := client.Operations.GetDynamicnodes(params)
	if err != nil {
		t.Errorf("TestGetDynamicnodes failed(%s)", err)
	}
}

func TestGetDynamicnodesID(t *testing.T) {
	params := operations.NewGetDynamicnodesIDParams()
	params.ID = "dyn2"
	_, err := client.Operations.GetDynamicnodesID(params)
	if err != nil {
		t.Errorf("TestGetDynamicnodesID failed(%s)", err)
	}
}

func TestPutApplicationsID(t *testing.T) {
	params := operations.NewPutApplicationsIDParams()
	var d models.Application
	str := `{"id":"application100","name":"Face recognition appl","description":"this is the Face recognition application",
    "microservices":[{"microservice_id":"microservice1"},{"microservice_id":"microservice2"},{"microservice_id":"microservice3"}],
    "status":"up"}`
	b := []byte(str)
	err := json.Unmarshal(b, &d)
	if err != nil {
		t.Errorf("TestPutApplicationsID failed(%s)", err)
	} else {
		params.Application = &d
		params.ID = "application100"
		_, err := client.Operations.PutApplicationsID(params)
		if err != nil {
			t.Errorf("TestPutApplicationsID failed(%s)", err)
		}
	}
}

func TestPutDeploymentsName(t *testing.T) {
	params := operations.NewPutDeploymentsNameParams()
	var d models.Deployment
	str := `{"name":"yet another appl","description":"appl description","externalendpoint_id":"thing1","status":"toundeploy",
  "dataflows":[{"source_id":"thing1","destination_id":"Camera driver","bandwidth_required":10000000,"latency_required": 100},
  {"source_id":"Camera driver","destination_id":"Face detector","bandwidth_required":9000000,"latency_required": 100},
  {"source_id":"Face detector","destination_id":"Face recognition","bandwidth_required":500000,"latency_required": 550}],
  "microservices":[{"name":"Camera driver","description 1":"msr description","cpu_required":"1000m","memory_required":"1Gi","disk_required":"0Gi","region_required":"reg3",
	"price_required":10,"price_computed":4,"deployment_descriptor":"deployment descriptor for ms 1"},
  {"name":"Face detector","description":"msr description 2","cpu_required":"2000m","memory_required":"2Gi","disk_required":"2Gi","region_required":"",
	"price_required":10,"price_computed":5,"deployment_descriptor":"deployment descriptor for ms 2"},
  {"name":"Face recognition","description":"msr description 3","cpu_required":"4000m","memory_required":"4Gi","disk_required":"4Gi","region_required":"",
	"price_required":15,"price_computed":55,"deployment_descriptor":"deployment descriptor for ms 3"}]}`
	b := []byte(str)
	err := json.Unmarshal(b, &d)
	if err != nil {
		t.Errorf("TestPutDeploymentsName failed(%s)", err)
	} else {
		params.Deployment = &d
		params.Name = "yet another appl"
		_, err := client.Operations.PutDeploymentsName(params)
		if err != nil {
			t.Errorf("TestPutDeploymentsName failed(%s)", err)
		}
	}
}

func TestPutMicroservicesID(t *testing.T) {
	params := operations.NewPutMicroservicesIDParams()
	var d models.Microservice
	str := `{"id":"microservice100","name": "Camera driver","description": "Camera Driver description",
    "status": "up","node_id": "node13","application_id": "application1","region_id": "reg3"}`
	b := []byte(str)
	err := json.Unmarshal(b, &d)
	if err != nil {
		t.Errorf("TestPutMicroservicesID failed(%s)", err)
	} else {
		params.Microservice = &d
		params.ID = "microservice100"
		_, err := client.Operations.PutMicroservicesID(params)
		if err != nil {
			t.Errorf("TestPutMicroservicesID failed(%s)", err)
		}
	}
}

func TestPutNodesID(t *testing.T) {
	params := operations.NewPutNodesIDParams()
	var d models.Node
	str := `{"id":"node100","cpu_capacity":"100000m","cpu_available":"100000m","memory_capacity":"100Gi",
    "memory_available":"100Gi","disk_capacity":"500Gi","disk_available":"500Gi","architecture":"AMD64",
    "distribution":"Linux","version":"Ubuntu 16.04","status":"up","region_id":"reg1"}`
	b := []byte(str)
	err := json.Unmarshal(b, &d)
	if err != nil {
		t.Errorf("TestPutNodesID failed(%s)", err)
	} else {
		params.Node = &d
		params.ID = "node100"
		_, err := client.Operations.PutNodesID(params)
		if err != nil {
			t.Errorf("TestPutNodesID failed(%s)", err)
		}
	}
}

func TestPutRegionsID(t *testing.T) {
	params := operations.NewPutRegionsIDParams()
	var d models.Region
	str := `{"id":"region100","description":"reg100 description","location":"Trento","tier":1,"relationships":[{"relationship_id":"rel1"},{"relationship_id":"rel2"}]}`
	b := []byte(str)
	err := json.Unmarshal(b, &d)
	if err != nil {
		t.Errorf("TestPutRegionsID failed(%s)", err)
	} else {
		params.Region = &d
		params.ID = "region100"
		_, err := client.Operations.PutRegionsID(params)
		if err != nil {
			t.Errorf("TestPutRegionsID failed(%s)", err)
		}
	}
}

func TestPutRelationshipsID(t *testing.T) {
	params := operations.NewPutRelationshipsIDParams()
	var d models.Relationship
	str := `{"id":"rel100","endpoint_a":"reg2","endpoint_b":"reg3","bandwidth_capacity":10000000,
    "bandwidth_available":10000000,"latency":50,"status":"up","region_id":"reg2"}`
	b := []byte(str)
	err := json.Unmarshal(b, &d)
	if err != nil {
		t.Errorf("TestPutRelationshipsID failed(%s)", err)
	} else {
		params.Relationship = &d
		params.ID = "rel100"
		_, err := client.Operations.PutRelationshipsID(params)
		if err != nil {
			t.Errorf("TestPutRelationshipsID failed(%s)", err)
		}
	}
}

func TestPutExternalendpointsID(t *testing.T) {
	params := operations.NewPutExternalendpointsIDParams()
	var d models.ExternalEndpoint
	str := `{"id":"thing100","description":"centesima thing","name":"Camera 1","ip_address": "1.2.3.4","location": "Povo","region_id":"reg3", "type": "Camera"}`
	b := []byte(str)
	err := json.Unmarshal(b, &d)
	if err != nil {
		t.Errorf("TestPutExternalendpointsID failed(%s)", err)
	} else {
		params.Externalendpoint = &d
		params.ID = "thing100"
		_, err := client.Operations.PutExternalendpointsID(params)
		if err != nil {
			t.Errorf("TestPutExternalendpointsID failed(%s)", err)
		}
	}
}

func TestPutDynamicnodesID(t *testing.T) {
	params := operations.NewPutDynamicnodesIDParams()
	var d models.DynamicNode
	str := `{"id":"dyn100","ip_address": "1.2.3.4","region_id":"reg3333"}`
	b := []byte(str)
	err := json.Unmarshal(b, &d)
	if err != nil {
		t.Errorf("TestPutDynamicnodesID failed(%s)", err)
	} else {
		params.Dynamicnode = &d
		params.ID = "dyn100"
		_, err := client.Operations.PutDynamicnodesID(params)
		if err != nil {
			t.Errorf("TestPutDynamicnodesID failed(%s)", err)
		}
	}
}

func TestPatchDeploymentsName(t *testing.T) {
	params := operations.NewPatchDeploymentsNameParams()
	var stat models.PatchStatus
	stat.Status = "todeploy"

	params.PatchStatus = &stat
	params.Name = "yet another appl"
	_, err := client.Operations.PatchDeploymentsName(params)
	if err != nil {
		t.Errorf("TestPatchDeploymentsName failed(%s)", err)
	}
}
