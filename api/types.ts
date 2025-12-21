// types.ts
import { Service } from './service';

// Interfaces
interface Org {
  id: string;
  name: string;
  state: string;
  created: Date;
  updated: Date;
}

interface Project {
  id: string;
  name: string;
  slug: string;
  description: string;
  org_id: string;
  org: Org;
}

interface Job {
  id: string;
  name: string;
  slug: string;
  description: string;
  org_id: string;
  org: Org;
  project_id: string;
  project: Project;
}

interface Trigger {
  id: string;
  name: string;
  type: string;
  job_id: string;
  job: Job;
  enabled: boolean;
  active: boolean;
  created: Date;
  updated: Date;
}

interface Event {
  id: string;
  name: string;
  description: string;
  trigger_id: string;
  trigger: Trigger;
  created: Date;
  updated: Date;
}

interface Config {
  token: string;
  url: string;
  org_id: string;
  project_id: string;
  job_id: string;
}

// Services
interface JenkinsService extends Service {
  getOrgs(): Promise<Org[]>;
  getProjects(orgId: string): Promise<Project[]>;
  getJobs(projectId: string): Promise<Job[]>;
  getTriggers(jobId: string): Promise<Trigger[]>;
  getEvents(triggerId: string): Promise<Event[]>;
}

interface GitLabService extends Service {
  getOrgs(): Promise<Org[]>;
  getProjects(orgId: string): Promise<Project[]>;
  getJobs(projectId: string): Promise<Job[]>;
  getTriggers(jobId: string): Promise<Trigger[]>;
  getEvents(triggerId: string): Promise<Event[]>;
}

type Service = JenkinsService | GitLabService;

// Enums
enum Provider {
  JENKINS = 'Jenkins',
  GITLAB = 'GitLab',
}

// Constants
const DEFAULT_TOKEN = 'your-token-here';
const DEFAULT_URL = 'your-url-here';