# ketikubecli

## Introduction of KetiKubeCli

> KETI���� ������ OpenMCP �÷����� Cluster �ڵ� Join/Unjoin ����� ȿ�������� ������ �� �ִ� ��ɾ� �������̽�

## Requirement
OpenMCP Master Node���� ketikubecli�� ����Ѵٸ�, openmcp ��ġ�� �Ϸ�Ǿ�� ��.

OpenMCP ���� Cluster Node�鿡�� ketikubecli�� ����Ѵٸ�, ������ openmcp ��ġ�� �ʿ����� ����. 


## How to Install
```
# ���� ���α׷� �����
1.build.sh

# ���� ���α׷� ��� ���� �� Config ���� ��� ����
2.install.sh
```

## Config ���� ����

> KetiKubeCli�� ������ ���� ������(/var/lib/ketikubecli/config.yaml)�� �ʿ��մϴ�.
```
# OpenMCP ��ġ ��� ����(OpenMCP Master�� ��츸)
openmcpDir: "/root/workspace/openmcp"

# External(nfs) ���� ����
nfsServer: "10.0.3.12"
```

## Governance

�� ������Ʈ�� ������ű�����＾��(IITP)���� �����ϴ� '19�� ������Ź�ۿ������߻������, "��ǻ�� �ڿ��� ������ Ȯ�� �� ���� �̵��� �����ϴ� �лꡤ������ �����̳� �÷��� ��� ���� ����" ��.
