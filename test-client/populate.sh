#!/bin/bash

if [ $# -lt 1 ]; then
    echo "Please provide valid argument"
    echo -e ${HELP}
    exit 1
fi

ENDPOINT=${1}

HELP="\nHELP: This is a simple bash script for the inventory (etcd) population. \n
You need to provide the following argument:\n
\t 1) etcd endpoint in the form: SCHEMA://HOST:PORT\n

Usage example: bash ${0##*/} http://127.0.0.1:2379\n"

#clean up
etcdctl --endpoints $ENDPOINT  rm --recursive /inventory

# create region directories
etcdctl --endpoints $ENDPOINT  mkdir /inventory
etcdctl --endpoints $ENDPOINT  mkdir /inventory/regions

#create region keys-values
etcdctl --endpoints $ENDPOINT  set /inventory/regions/reg1 \
'{"id":"reg1","description":"reg1 description","location":"Munich","tier":0}'
etcdctl --endpoints $ENDPOINT  set /inventory/regions/reg2 \
'{"id":"reg2","description":"reg2 description","location":"Trento","tier":1,"prices":{"cpu":{"min_price":10,"max_price":1,"scarcity":1,"unit_price":5},
"memory":{"min_price":10,"max_price":1,"scarcity":1,"unit_price":5},"disk":{"min_price":10,"max_price":1,"scarcity":1,"unit_price":5}}}'
etcdctl --endpoints $ENDPOINT  set /inventory/regions/reg3 \
'{"id":"reg3","description":"reg3 description","location":"Povo","tier":2,"prices":{"cpu":{"min_price":11,"max_price":2,"scarcity":2,"unit_price":6},
"memory":{"min_price":11,"max_price":2,"scarcity":2,"unit_price":6},"disk":{"min_price":11,"max_price":2,"scarcity":2,"unit_price":6}}}'

# create deployment directories
etcdctl --endpoints $ENDPOINT  mkdir /inventory/deployments

#create deployment keys-values
etcdctl --endpoints $ENDPOINT  set /inventory/deployments/"Face recognition appl" \
'{"name":"Face recognition appl","description":"appl description","externalendpoint_id":"thing1","status":"todeploy",
"dataflows":[{"source_id":"thing1","destination_id":"Camera driver","bandwidth_required":10000000,"latency_required": 100},
{"source_id":"Camera driver","destination_id":"Face detector","bandwidth_required":9000000,"latency_required": 100},
{"source_id":"Face detector","destination_id":"Face recognition","bandwidth_required":500000,"latency_required": 550}],
"microservices":[{"name":"Camera driver","description":"msr description","cpu_required":"1000m","memory_required":"1Gi","disk_required":"0Gi","region_required":"reg3",
"price_required":10,"price_computed":5,"deployment_descriptor":"deployment descriptor for ms 1"},
{"name":"Face detector","description":"msr description 2","cpu_required":"2000m","memory_required":"2Gi","disk_required":"2Gi","region_required":"",
"price_required":11,"price_computed":6,"deployment_descriptor":"deployment descriptor for ms 2"},
{"name":"Face recognition","description":"msr description 3","cpu_required":"4000m","memory_required":"4Gi","disk_required":"4Gi","region_required":"",
"price_required":12,"price_computed":7,"deployment_descriptor":"deployment descriptor for ms 3"}]}'
etcdctl --endpoints $ENDPOINT  set /inventory/deployments/"Face recognition appl1" \
'{"name":"Face recognition appl1","description":"appl description","externalendpoint_id":"thing1","status":"todeploy",
"dataflows":[{"source_id":"thing1","destination_id":"Camera driver","bandwidth_required":10000000,"latency_required": 100},
{"source_id":"Camera driver","destination_id":"Face detector","bandwidth_required":9000000,"latency_required": 100},
{"source_id":"Face detector","destination_id":"Face recognition","bandwidth_required":500000,"latency_required": 550}],
"microservices":[{"name":"Camera driver","description":"msr description","cpu_required":"1000m","memory_required":"1Gi","disk_required":"0Gi","region_required":"reg3",
"price_required":100,"price_computed":50,"deployment_descriptor":"deployment descriptor for ms 1"},
{"name":"Face detector","description":"msr description 2","cpu_required":"2000m","memory_required":"2Gi","disk_required":"2Gi","region_required":"",
"price_required":101,"price_computed":51,"deployment_descriptor":"deployment descriptor for ms 2"},
{"name":"Face recognition","description":"msr description 3","cpu_required":"4000m","memory_required":"4Gi","disk_required":"4Gi","region_required":"",
"price_required":102,"price_computed":52,"deployment_descriptor":"deployment descriptor for ms 3"}]}'


#create node directories
etcdctl --endpoints $ENDPOINT  mkdir /inventory/nodes

#create node keys-values
etcdctl --endpoints $ENDPOINT  set /inventory/nodes/node11 \
'{"id":"node11","cpu_capacity":"100000m","cpu_available":"100000m","memory_capacity":"100Gi","memory_available":"100Gi","disk_capacity":"500Gi","disk_available":"500Gi","architecture":"AMD64","distribution":"Linux","version":"Ubuntu 16.04","status":"up","region_id":"reg1"}'
etcdctl --endpoints $ENDPOINT  set /inventory/nodes/node21 \
'{"id":"node21","cpu_capacity":"100000m","cpu_available":"100000m","memory_capacity":"100Gi","memory_available":"100Gi","disk_capacity":"500Gi","disk_available":"500Gi","architecture":"AMD64","distribution":"Linux","version":"Ubuntu 16.04","status":"up","region_id":"reg1"}'
etcdctl --endpoints $ENDPOINT  set /inventory/nodes/node31 \
'{"id":"node31","cpu_capacity":"100000m","cpu_available":"100000m","memory_capacity":"100Gi","memory_available":"100Gi","disk_capacity":"500Gi","disk_available":"500Gi","architecture":"AMD64","distribution":"Linux","version":"Ubuntu 16.04","status":"up","region_id":"reg1"}'
etcdctl --endpoints $ENDPOINT  set /inventory/nodes/node41 \
'{"id":"node41","cpu_capacity":"100000m","cpu_available":"100000m","memory_capacity":"100Gi","memory_available":"100Gi","disk_capacity":"500Gi","disk_available":"500Gi","architecture":"AMD64","distribution":"Linux","version":"Ubuntu 16.04","status":"up","region_id":"reg1"}'
etcdctl --endpoints $ENDPOINT  set /inventory/nodes/node12 \
'{"id":"node12","cpu_capacity":"10000m","cpu_available":"10000m","memory_capacity":"10Gi","memory_available":"10Gi","disk_capacity":"50Gi","disk_available":"50Gi","architecture":"AMD64","distribution":"Linux","version":"Ubuntu 16.04","status":"up","region_id":"reg2"}'
etcdctl --endpoints $ENDPOINT  set /inventory/nodes/node22 \
'{"id":"node22","cpu_capacity":"10000m","cpu_available":"10000m","memory_capacity":"10Gi","memory_available":"10Gi","disk_capacity":"50Gi","disk_available":"50Gi","architecture":"AMD64","distribution":"Linux","version":"Ubuntu 16.04","status":"up","region_id":"reg2"}'
etcdctl --endpoints $ENDPOINT  set /inventory/nodes/node32 \
'{"id":"node32","cpu_capacity":"10000m","cpu_available":"10000m","memory_capacity":"10Gi","memory_available":"10Gi","disk_capacity":"50Gi","disk_available":"50Gi","architecture":"AMD64","distribution":"Linux","version":"Ubuntu 16.04","status":"up","region_id":"reg2"}'
etcdctl --endpoints $ENDPOINT  set /inventory/nodes/node42 \
'{"id":"node42","cpu_capacity":"10000m","cpu_available":"10000m","memory_capacity":"10Gi","memory_available":"10Gi","disk_capacity":"50Gi","disk_available":"50Gi","architecture":"AMD64","distribution":"Linux","version":"Ubuntu 16.04","status":"up","region_id":"reg2"}'
etcdctl --endpoints $ENDPOINT  set /inventory/nodes/node13 \
'{"id":"node13","cpu_capacity":"4000m","cpu_available":"4000m","memory_capacity":"4Gi","memory_available":"4Gi","disk_capacity":"20Gi","disk_available":"20Gi","architecture":"AMD64","distribution":"Linux","version":"Ubuntu 16.04","status":"up","region_id":"reg3"}'

#create microservice directories
etcdctl --endpoints $ENDPOINT  mkdir /inventory/microservices

#create microservice keys-values
etcdctl --endpoints $ENDPOINT  set /inventory/microservices/microservice1 \
'{"id":"microservice1","name": "Camera driver","description": "Camera Driver description","status": "up","node_id": "node13","application_id": "application1","region_id": "reg3"}'
etcdctl --endpoints $ENDPOINT  set /inventory/microservices/microservice2 \
'{"id":"microservice2","name": "Face detector","description": "Face detector description","status": "up","node_id": "node12","application_id": "application1","region_id": "reg2"}'
etcdctl --endpoints $ENDPOINT  set /inventory/microservices/microservice3 \
'{"id":"microservice3","name": "Face recognition","description": "Face recognition description","status": "up","node_id": "node12","application_id": "application1","region_id": "reg2"}'
etcdctl --endpoints $ENDPOINT  set /inventory/microservices/microservice4 \
'{"id":"microservice4","name": "mysql","description": "text","status": "up","node_id": "node11","application_id": "application2","region_id": "reg1"}'

#create application directories
etcdctl --endpoints $ENDPOINT  mkdir /inventory/applications

#create application keys-values
etcdctl --endpoints $ENDPOINT  set /inventory/applications/application1 \
'{"id":"application1","name":"Face recognition appl","description":"this is the Face recognition application","microservices":[{"microservice_id":"microservice1"},{"microservice_id":"microservice2"},{"microservice_id":"microservice3"}],"status":"up"}'
etcdctl --endpoints $ENDPOINT  set /inventory/applications/application2 \
'{"id":"application2","name":"your_application","description":"this is your application","microservices":[{"microservice_id":"microservice3"},{"microservice_id":"microservice4"}],"status":"up"}'

#create external endpoint directories
etcdctl --endpoints $ENDPOINT  mkdir /inventory/externalendpoints

#create external endpoint keys-values
etcdctl --endpoints $ENDPOINT  set /inventory/externalendpoints/thing1 \
'{"id":"thing1","description":"una thing","name":"Camera 1","ip_address": "1.2.3.4","location": "Povo","region_id":"reg3", "type": "Camera"}'
etcdctl --endpoints $ENDPOINT  set /inventory/externalendpoints/thing2 \
'{"id":"thing2","description":"seconda thing","name":"Camera 1","ip_address": "1.2.3.4","location": "Povo","region_id":"reg3", "type": "Camera"}'

#create dynamicnode directories
etcdctl --endpoints $ENDPOINT  mkdir /inventory/dynamicnodes

#create dynamic node keys-values
etcdctl --endpoints $ENDPOINT  set /inventory/dynamicnodes/dyn1 \
'{"id":"dyn1","ip_address": "1.2.3.4","region_id":"reg3", "node_id": "12344"}'
etcdctl --endpoints $ENDPOINT  set /inventory/dynamicnodes/dyn2 \
'{"id":"dyn2","ip_address": "1.2.3.4","region_id":"reg3"}'

#create relationship directories
etcdctl --endpoints $ENDPOINT  mkdir /inventory/relationships

#create relationship keys-values
etcdctl --endpoints $ENDPOINT  set /inventory/relationships/rel1 \
'{"id":"rel1","endpoint_a":"reg2","endpoint_b":"reg3","bandwidth_capacity":10000000,"bandwidth_available":10000000,"latency":50,"status":"up",
"prices":{"bandwidth":{"min_price":10,"max_price":1,"scarcity":1,"unit_price":5},"latency":{"min_price":10,"max_price":1,"scarcity":1,"unit_price":5}},
"region_id":"reg2"}'
etcdctl --endpoints $ENDPOINT  set /inventory/relationships/rel2 \
'{"id":"rel2","endpoint_a":"reg2","endpoint_b":"reg1","bandwidth_capacity":1000000,"bandwidth_available":1000000,"latency":500,"status":"up",
"prices":{"bandwidth":{"min_price":10,"max_price":1,"scarcity":1,"unit_price":5},"latency":{"min_price":10,"max_price":1,"scarcity":1,"unit_price":56}},
"region_id":"reg2"}'
etcdctl --endpoints $ENDPOINT  set /inventory/relationships/rel3 \
'{"id":"rel3","endpoint_a":"reg3","endpoint_b":"reg1","bandwidth_capacity":1000000,"bandwidth_available":1000000,"latency":600,"status":"up",
"region_id":"reg3"}'
