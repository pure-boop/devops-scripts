// Types
import { Volume, Container } from 'dockerode';

// Docker Volume Types
export enum DockerVolumeType {
  LOCAL = 'local',
  BIND = 'bind',
  TEMPORARY = 'temporary',
  NAMED = 'named',
}

// Docker Container Options
export interface IDockerContainerOptions {
  cpuCount: number;
  cpuPercent: number;
  memoryLimit: number;
  memorySwap: number;
  cpuSet: string;
  memSwap: string;
  memLimit: string;
  memReservation: string;
  oomKillDisable: boolean;
  pidsLimit: number;
  pidsCount: number;
  memSwapLimit: number;
  memSwapPercent: number;
  memLimitPercent: number;
  pidsLimitPercent: number;
  cpusetMems: string;
  cpuShares: number;
  cpuQuota: number;
  cpusetCpus: string;
  memSwappiness: number;
  runtime: string;
  blkmemorySize: number;
  deviceReadBps: string;
  deviceWriteBps: string;
  deviceReadIOPS: string;
  deviceWriteIOPS: string;
  cpus: string;
  cpuset: string;
}

// Docker Container Create Options
export interface IDockerCreateOptions {
  image: string;
  env: { [key: string]: string };
  volumes: {
    type: DockerVolumeType;
    source: string;
  }[];
  ports: { hostPort: number; containerPort: number }[];
  deploy: {
    resources: {
      limits: {
        cpus: number;
        memory: number;
      };
      reservations: {
        cpus: number;
        memory: number;
      };
    };
    placement: {
      constraints: [
        {
          type: string;
          key: string;
          value: string;
        }
      ];
    };
  };
}

// Docker Container Stats
export interface IDockerContainerStats {
  readIOCounters: {
    count: number;
    total: {
      total: number;
      eta: {
        seconds: number;
      };
    };
  };
  cpuStats: {
    cpuUsage: number;
    totalUsage: number;
    percpuUsage: number;
    systemUsage: number;
    cpuQuota: number;
  };
  memoryStats: {
    usage: number;
    limit: number;
    stats: {
      usage: number;
    };
  };
  networkThrottlingDataReadBps: number;
  networkThrottlingDataWriteBps: number;
}

// Docker Container Volumes
export interface IDockerVolumes {
  DockerVolume: Volume;
  BindVolume: Volume;
  TempVolume: Volume;
  NamedVolume: Volume;
}

// Docker Container
export interface IDockerContainer extends Container {
  stats(): Promise<IDockerContainerStats>;
  logs(): Promise<string>;
  rm(force: boolean): Promise<void>;
  pause(): Promise<void>;
  unpause(): Promise<void>;
  top(): Promise<{ [containerId: string]: string[] >>;
  topContainerId(): Promise<string>;
}

// Docker
export interface IDocker {
  createContainer(options: IDockerCreateOptions): Promise<IDockerContainer>;
  getContainer(id: string): Promise<IDockerContainer | null>;
  getContainerAll(): Promise<IDockerContainer[]>;
  getNetwork(id: string): Promise<any>;
  getNetworks(): Promise<any[]>;
  getVolumes(): Promise<IDockerVolumes>;
  stop(id: string): Promise<void>;
  start(id: string): Promise<void>;
}

// Dockerode Volume
export interface IDockerodeVolume {
  ls(): Promise<Volume[]>;
  rm(id: string): Promise<void>;
}