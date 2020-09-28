# omctl

## Introduction of omctl

> Command Language Interface for Member Clusters for Cluster Join/Unjoin Function of OpenMCP Platform developed by KETI
>

## Requirement
1. [External Server to Store Cluster Information](https://github.com/openmcp/external)
1. go 1.14.2 Installation
1. nfs-common Installation (apt-get install nfs-common)


## How to Install
1.Build after setting environment variables at build.sh
```
$ vim 1.build.sh
...
EXTERNAL_IP="10.0.3.12"    # Specifying an External(nfs) server
...

$ ./1.build.sh
```

## How to Use
Join process using [omcpctl](https://github.com/openmcp/openmcp/tree/master/omcpctl) for OpenMCP master cluster and [omcpctl](https://github.com/openmcp/openmcp-cli) for OpneMCP member cluster after [installation of OpenMCP master](https://github.com/openmcp/openmcp)

1. Registering Openmcp in an OpenMCP Master (performing in OpenMCP)
```
$ omcpctl register master
```
2. Register an OpenMCP Member Cluster with OpenMCP (performing in a subcluster) 
```
$ omctl register member <OpenMCP_Master_IP>
```
> Check nfs-common installation if error occurs when mount

3. Query currently available clusters with OpenMCP (performing in OpenMCP)
```
$ omcpctl joinable list

  CLUSTERNAME |           APIENDPOINT           | PLATFORM  
+-------------+---------------------------------+----------+
  cluster3    | https://CLUSTERIP3_IP:6443      |   
```
4. Join Member Cluster in OpenMCP and deploy default modules (performing in OpenMCP)
```
$ omcpctl join cluster <OpenMCP_Member_IP>
```
5. Query clusters currently joined to OpenMCP (performing in OpenMCP)
```
$ omcpctl get cluster -A

Cluster :  openmcp
            NS           | CLUSTERNAME  | STATUS |   REGION    |         ZONES         |           APIENDPOINT               | PLATFORM |  AGE   
+------------------------+--------------+--------+-------------+-----------------------+-------------------------------------+----------+-------+
  kube-federation-system | cluster1     | True   | AS          | KR                    | https://10.0.3.50:6443              |          | 7d2h   
  kube-federation-system | cluster2     | True   | AS          | KR                    | https://10.0.3.70:6443              |          | 35d    
  kube-federation-system | cluster3     | True   | AS          | KR                    | https://10.0.3.80:6443              |          | 20s 
```


## Governance

This project was supported by Institute of Information & communications Technology Planning & evaluation (IITP) grant funded by the Korea government (MSIT)
(No.2019-0-00052, Development of Distributed and Collaborative Container Platform enabling Auto Scaling and Service Mobility)
