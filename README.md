# omctl

## Introduction of omctl

> KETI에서 개발한 OpenMCP 플랫폼의 Cluster Join/Unjoin 기능을 위한 Member Cluster용 명령어 인터페이스
>

## Requirement
1. [클러스터 정보를 저장할 External 서버](https://github.com/openmcp/external)
1. go 1.14.2 설치
1. nfs-common 설치


## How to Install
1.build.sh 에서 환경변수 설정 후 빌드
```
$ vim 1.build.sh
...
EXTERNAL_IP="10.0.3.12"    # External(nfs) 서버 지정
...

$ ./1.build.sh
```

## How to Use
[OpenMCP Master 설치](https://github.com/openmcp/openmcp) 후 OpenMCP Master Cluster용 [omcpctl](https://github.com/openmcp/openmcp/tree/master/omcpctl)과 OpneMCP Member Cluster용 [omctl](https://github.com/openmcp/openmcp-cli)을 이용한 Join 과정


```
1. OpenMCP Master Cluster에서 Openmcp 등록 (OpenMCP에서 수행)
$ omcpctl register master

2. OpenMCP Member Cluster 등록 (하위 클러스터에서 수행) 
$ omctl register member <OpenMCP_Master_IP>

* mount 오류 시 nfs-common 설치 확인

3. 현재 OpenMCP Join된 클러스터 조회 (OpenMCP에서 수행)
$ omcpctl get cluster -A

Cluster :  openmcp
            NS           | CLUSTERNAME  | STATUS |   REGION    |         ZONES         |           APIENDPOINT               | PLATFORM |  AGE   
+------------------------+--------------+--------+-------------+-----------------------+-------------------------------------+----------+-------+
  kube-federation-system | cluster1     | True   | AS          | KR                    | https://10.0.3.50:6443              |          | 7d2h   
  kube-federation-system | cluster2     | True   | AS          | KR                    | https://10.0.3.70:6443              |          | 35d    
  kube-federation-system | cluster3     | True   | AS          | KR                    | https://10.0.3.80:6443              |          | 20d      


4. 현재 OpenMCP 조인가능한 클러스터 조회 (OpenMCP에서 수행)
$ omcpctl joinable list

  CLUSTERNAME |           APIENDPOINT           | PLATFORM  
+-------------+---------------------------------+----------+
  cluster3    | https://CLUSTERIP3_IP:6443      |   

5. Cluster Join 및 기본 모듈 배포 (OpenMCP에서 수행)
$ omcpctl join cluster <OpenMCP_Member_IP>

6. 현재 OpenMCP Join된 클러스터 조회 (OpenMCP에서 수행)
$ omcpctl get cluster -A

Cluster :  openmcp
            NS           | CLUSTERNAME  | STATUS |   REGION    |         ZONES         |           APIENDPOINT               | PLATFORM |  AGE   
+------------------------+--------------+--------+-------------+-----------------------+-------------------------------------+----------+-------+
  kube-federation-system | cluster1     | True   | AS          | KR                    | https://10.0.3.50:6443              |          | 7d2h   
  kube-federation-system | cluster2     | True   | AS          | KR                    | https://10.0.3.70:6443              |          | 35d    
  kube-federation-system | cluster3     | True   | AS          | KR                    | https://10.0.3.80:6443              |          | 20d  
```


## Governance

본 프로젝트는 정보통신기술진흥센터(IITP)에서 지원하는 '19년 정보통신방송연구개발사업으로, "컴퓨팅 자원의 유연한 확장 및 서비스 이동을 제공하는 분산·협업형 컨테이너 플랫폼 기술 개발 과제" 임.
