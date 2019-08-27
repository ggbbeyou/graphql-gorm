package gen

type key int

const (
	KeyPrincipalID      key    = iota
	KeyLoaders          key    = iota
	KeyExecutableSchema key    = iota
	KeyJWTClaims        key    = iota
	SchemaSDL           string = `scalar Time

type Query {
  user(id: ID, q: String, filter: UserFilterType): User
  users(current_page: Int = 1, per_page: Int = 20, q: String, sort: [UserSortType!], filter: UserFilterType): UserResultType
  task(id: ID, q: String, filter: TaskFilterType): Task
  tasks(current_page: Int = 1, per_page: Int = 20, q: String, sort: [TaskSortType!], filter: TaskFilterType): TaskResultType
  admin(id: ID, q: String, filter: AdminFilterType): Admin
  admins(current_page: Int = 1, per_page: Int = 20, q: String, sort: [AdminSortType!], filter: AdminFilterType): AdminResultType
  group(id: ID, q: String, filter: GroupFilterType): Group
  groups(current_page: Int = 1, per_page: Int = 20, q: String, sort: [GroupSortType!], filter: GroupFilterType): GroupResultType
  role(id: ID, q: String, filter: RoleFilterType): Role
  roles(current_page: Int = 1, per_page: Int = 20, q: String, sort: [RoleSortType!], filter: RoleFilterType): RoleResultType
}

type Mutation {
  createUser(input: UserCreateInput!): User!
  updateUser(id: ID!, input: UserUpdateInput!): User!
  deleteUser(id: ID!): User!
  deleteAllUsers: Boolean!
  createTask(input: TaskCreateInput!): Task!
  updateTask(id: ID!, input: TaskUpdateInput!): Task!
  deleteTask(id: ID!): Task!
  deleteAllTasks: Boolean!
  createAdmin(input: AdminCreateInput!): Admin!
  updateAdmin(id: ID!, input: AdminUpdateInput!): Admin!
  deleteAdmin(id: ID!): Admin!
  deleteAllAdmins: Boolean!
  createGroup(input: GroupCreateInput!): Group!
  updateGroup(id: ID!, input: GroupUpdateInput!): Group!
  deleteGroup(id: ID!): Group!
  deleteAllGroups: Boolean!
  createRole(input: RoleCreateInput!): Role!
  updateRole(id: ID!, input: RoleUpdateInput!): Role!
  deleteRole(id: ID!): Role!
  deleteAllRoles: Boolean!
}

scalar Any

type User {
  id: ID!
  email: String
  firstName: String
  lastName: String
  tasks: [Task!]!
  state: Int
  del: Int
  updatedAt: Int
  createdAt: Int
  deletedBy: ID
  updatedBy: ID
  createdBy: ID
  tasksIds: [ID!]!
}

type Task {
  id: ID!
  title: String
  completed: Boolean
  dueDate: Time
  assignee: User
  assigneeId: ID
  state: Int
  del: Int
  updatedAt: Int
  createdAt: Int
  deletedBy: ID
  updatedBy: ID
  createdBy: ID
}

type Admin {
  id: ID!
  Phone: String
  Password: String
  Username: String
  Money: Int
  Sex: Int
  Super: Int
  LoginCount: Int
  LoginIp: String
  LastIp: String
  groups: [Group!]
  roles: [Role!]
  state: Int
  del: Int
  updatedAt: Int
  createdAt: Int
  deletedBy: ID
  updatedBy: ID
  createdBy: ID
  groupsIds: [ID!]!
  rolesIds: [ID!]!
}

type Group {
  id: ID!
  Name: String
  admin: [Admin!]
  roles: [Role!]
  state: Int
  del: Int
  updatedAt: Int
  createdAt: Int
  deletedBy: ID
  updatedBy: ID
  createdBy: ID
  adminIds: [ID!]!
  rolesIds: [ID!]!
}

type Role {
  id: ID!
  Name: String
  admin: [Admin!]
  group: [Admin]
  Pid: String
  state: Int
  del: Int
  updatedAt: Int
  createdAt: Int
  deletedBy: ID
  updatedBy: ID
  createdBy: ID
  adminIds: [ID!]!
  groupIds: [ID!]!
}

extend type Mutation {
  login(email: String!): Any
}

input UserCreateInput {
  id: ID
  email: String
  firstName: String
  lastName: String
  state: Int
  del: Int
  tasksIds: [ID!]
}

input UserUpdateInput {
  email: String
  firstName: String
  lastName: String
  state: Int
  del: Int
  tasksIds: [ID!]
}

enum UserSortType {
  ID_ASC
  ID_DESC
  EMAIL_ASC
  EMAIL_DESC
  FIRST_NAME_ASC
  FIRST_NAME_DESC
  LAST_NAME_ASC
  LAST_NAME_DESC
  STATE_ASC
  STATE_DESC
  DEL_ASC
  DEL_DESC
  UPDATED_AT_ASC
  UPDATED_AT_DESC
  CREATED_AT_ASC
  CREATED_AT_DESC
  DELETED_BY_ASC
  DELETED_BY_DESC
  UPDATED_BY_ASC
  UPDATED_BY_DESC
  CREATED_BY_ASC
  CREATED_BY_DESC
  TASKS_IDS_ASC
  TASKS_IDS_DESC
}

input UserFilterType {
  AND: [UserFilterType!]
  OR: [UserFilterType!]
  id: ID
  id_ne: ID
  id_gt: ID
  id_lt: ID
  id_gte: ID
  id_lte: ID
  id_in: [ID!]
  email: String
  email_ne: String
  email_gt: String
  email_lt: String
  email_gte: String
  email_lte: String
  email_in: [String!]
  email_like: String
  email_prefix: String
  email_suffix: String
  firstName: String
  firstName_ne: String
  firstName_gt: String
  firstName_lt: String
  firstName_gte: String
  firstName_lte: String
  firstName_in: [String!]
  firstName_like: String
  firstName_prefix: String
  firstName_suffix: String
  lastName: String
  lastName_ne: String
  lastName_gt: String
  lastName_lt: String
  lastName_gte: String
  lastName_lte: String
  lastName_in: [String!]
  lastName_like: String
  lastName_prefix: String
  lastName_suffix: String
  state: Int
  state_ne: Int
  state_gt: Int
  state_lt: Int
  state_gte: Int
  state_lte: Int
  state_in: [Int!]
  del: Int
  del_ne: Int
  del_gt: Int
  del_lt: Int
  del_gte: Int
  del_lte: Int
  del_in: [Int!]
  updatedAt: Int
  updatedAt_ne: Int
  updatedAt_gt: Int
  updatedAt_lt: Int
  updatedAt_gte: Int
  updatedAt_lte: Int
  updatedAt_in: [Int!]
  createdAt: Int
  createdAt_ne: Int
  createdAt_gt: Int
  createdAt_lt: Int
  createdAt_gte: Int
  createdAt_lte: Int
  createdAt_in: [Int!]
  deletedBy: ID
  deletedBy_ne: ID
  deletedBy_gt: ID
  deletedBy_lt: ID
  deletedBy_gte: ID
  deletedBy_lte: ID
  deletedBy_in: [ID!]
  updatedBy: ID
  updatedBy_ne: ID
  updatedBy_gt: ID
  updatedBy_lt: ID
  updatedBy_gte: ID
  updatedBy_lte: ID
  updatedBy_in: [ID!]
  createdBy: ID
  createdBy_ne: ID
  createdBy_gt: ID
  createdBy_lt: ID
  createdBy_gte: ID
  createdBy_lte: ID
  createdBy_in: [ID!]
  tasks: TaskFilterType
}

type UserResultType {
  data: [User!]!
  total: Int!
  current_page: Int!
  per_page: Int!
  total_page: Int!
}

input TaskCreateInput {
  id: ID
  title: String
  completed: Boolean
  dueDate: Time
  assigneeId: ID
  state: Int
  del: Int
}

input TaskUpdateInput {
  title: String
  completed: Boolean
  dueDate: Time
  assigneeId: ID
  state: Int
  del: Int
}

enum TaskSortType {
  ID_ASC
  ID_DESC
  TITLE_ASC
  TITLE_DESC
  COMPLETED_ASC
  COMPLETED_DESC
  DUE_DATE_ASC
  DUE_DATE_DESC
  ASSIGNEE_ID_ASC
  ASSIGNEE_ID_DESC
  STATE_ASC
  STATE_DESC
  DEL_ASC
  DEL_DESC
  UPDATED_AT_ASC
  UPDATED_AT_DESC
  CREATED_AT_ASC
  CREATED_AT_DESC
  DELETED_BY_ASC
  DELETED_BY_DESC
  UPDATED_BY_ASC
  UPDATED_BY_DESC
  CREATED_BY_ASC
  CREATED_BY_DESC
}

input TaskFilterType {
  AND: [TaskFilterType!]
  OR: [TaskFilterType!]
  id: ID
  id_ne: ID
  id_gt: ID
  id_lt: ID
  id_gte: ID
  id_lte: ID
  id_in: [ID!]
  title: String
  title_ne: String
  title_gt: String
  title_lt: String
  title_gte: String
  title_lte: String
  title_in: [String!]
  title_like: String
  title_prefix: String
  title_suffix: String
  completed: Boolean
  completed_ne: Boolean
  completed_gt: Boolean
  completed_lt: Boolean
  completed_gte: Boolean
  completed_lte: Boolean
  completed_in: [Boolean!]
  dueDate: Time
  dueDate_ne: Time
  dueDate_gt: Time
  dueDate_lt: Time
  dueDate_gte: Time
  dueDate_lte: Time
  dueDate_in: [Time!]
  assigneeId: ID
  assigneeId_ne: ID
  assigneeId_gt: ID
  assigneeId_lt: ID
  assigneeId_gte: ID
  assigneeId_lte: ID
  assigneeId_in: [ID!]
  state: Int
  state_ne: Int
  state_gt: Int
  state_lt: Int
  state_gte: Int
  state_lte: Int
  state_in: [Int!]
  del: Int
  del_ne: Int
  del_gt: Int
  del_lt: Int
  del_gte: Int
  del_lte: Int
  del_in: [Int!]
  updatedAt: Int
  updatedAt_ne: Int
  updatedAt_gt: Int
  updatedAt_lt: Int
  updatedAt_gte: Int
  updatedAt_lte: Int
  updatedAt_in: [Int!]
  createdAt: Int
  createdAt_ne: Int
  createdAt_gt: Int
  createdAt_lt: Int
  createdAt_gte: Int
  createdAt_lte: Int
  createdAt_in: [Int!]
  deletedBy: ID
  deletedBy_ne: ID
  deletedBy_gt: ID
  deletedBy_lt: ID
  deletedBy_gte: ID
  deletedBy_lte: ID
  deletedBy_in: [ID!]
  updatedBy: ID
  updatedBy_ne: ID
  updatedBy_gt: ID
  updatedBy_lt: ID
  updatedBy_gte: ID
  updatedBy_lte: ID
  updatedBy_in: [ID!]
  createdBy: ID
  createdBy_ne: ID
  createdBy_gt: ID
  createdBy_lt: ID
  createdBy_gte: ID
  createdBy_lte: ID
  createdBy_in: [ID!]
  assignee: UserFilterType
}

type TaskResultType {
  data: [Task!]!
  total: Int!
  current_page: Int!
  per_page: Int!
  total_page: Int!
}

input AdminCreateInput {
  id: ID
  Phone: String
  Password: String
  Username: String
  Money: Int
  Sex: Int
  Super: Int
  LoginCount: Int
  LoginIp: String
  LastIp: String
  state: Int
  del: Int
  groupsIds: [ID!]
  rolesIds: [ID!]
}

input AdminUpdateInput {
  Phone: String
  Password: String
  Username: String
  Money: Int
  Sex: Int
  Super: Int
  LoginCount: Int
  LoginIp: String
  LastIp: String
  state: Int
  del: Int
  groupsIds: [ID!]
  rolesIds: [ID!]
}

enum AdminSortType {
  ID_ASC
  ID_DESC
  PHONE_ASC
  PHONE_DESC
  PASSWORD_ASC
  PASSWORD_DESC
  USERNAME_ASC
  USERNAME_DESC
  MONEY_ASC
  MONEY_DESC
  SEX_ASC
  SEX_DESC
  SUPER_ASC
  SUPER_DESC
  LOGIN_COUNT_ASC
  LOGIN_COUNT_DESC
  LOGIN_IP_ASC
  LOGIN_IP_DESC
  LAST_IP_ASC
  LAST_IP_DESC
  STATE_ASC
  STATE_DESC
  DEL_ASC
  DEL_DESC
  UPDATED_AT_ASC
  UPDATED_AT_DESC
  CREATED_AT_ASC
  CREATED_AT_DESC
  DELETED_BY_ASC
  DELETED_BY_DESC
  UPDATED_BY_ASC
  UPDATED_BY_DESC
  CREATED_BY_ASC
  CREATED_BY_DESC
  GROUPS_IDS_ASC
  GROUPS_IDS_DESC
  ROLES_IDS_ASC
  ROLES_IDS_DESC
}

input AdminFilterType {
  AND: [AdminFilterType!]
  OR: [AdminFilterType!]
  id: ID
  id_ne: ID
  id_gt: ID
  id_lt: ID
  id_gte: ID
  id_lte: ID
  id_in: [ID!]
  Phone: String
  Phone_ne: String
  Phone_gt: String
  Phone_lt: String
  Phone_gte: String
  Phone_lte: String
  Phone_in: [String!]
  Phone_like: String
  Phone_prefix: String
  Phone_suffix: String
  Password: String
  Password_ne: String
  Password_gt: String
  Password_lt: String
  Password_gte: String
  Password_lte: String
  Password_in: [String!]
  Password_like: String
  Password_prefix: String
  Password_suffix: String
  Username: String
  Username_ne: String
  Username_gt: String
  Username_lt: String
  Username_gte: String
  Username_lte: String
  Username_in: [String!]
  Username_like: String
  Username_prefix: String
  Username_suffix: String
  Money: Int
  Money_ne: Int
  Money_gt: Int
  Money_lt: Int
  Money_gte: Int
  Money_lte: Int
  Money_in: [Int!]
  Sex: Int
  Sex_ne: Int
  Sex_gt: Int
  Sex_lt: Int
  Sex_gte: Int
  Sex_lte: Int
  Sex_in: [Int!]
  Super: Int
  Super_ne: Int
  Super_gt: Int
  Super_lt: Int
  Super_gte: Int
  Super_lte: Int
  Super_in: [Int!]
  LoginCount: Int
  LoginCount_ne: Int
  LoginCount_gt: Int
  LoginCount_lt: Int
  LoginCount_gte: Int
  LoginCount_lte: Int
  LoginCount_in: [Int!]
  LoginIp: String
  LoginIp_ne: String
  LoginIp_gt: String
  LoginIp_lt: String
  LoginIp_gte: String
  LoginIp_lte: String
  LoginIp_in: [String!]
  LoginIp_like: String
  LoginIp_prefix: String
  LoginIp_suffix: String
  LastIp: String
  LastIp_ne: String
  LastIp_gt: String
  LastIp_lt: String
  LastIp_gte: String
  LastIp_lte: String
  LastIp_in: [String!]
  LastIp_like: String
  LastIp_prefix: String
  LastIp_suffix: String
  state: Int
  state_ne: Int
  state_gt: Int
  state_lt: Int
  state_gte: Int
  state_lte: Int
  state_in: [Int!]
  del: Int
  del_ne: Int
  del_gt: Int
  del_lt: Int
  del_gte: Int
  del_lte: Int
  del_in: [Int!]
  updatedAt: Int
  updatedAt_ne: Int
  updatedAt_gt: Int
  updatedAt_lt: Int
  updatedAt_gte: Int
  updatedAt_lte: Int
  updatedAt_in: [Int!]
  createdAt: Int
  createdAt_ne: Int
  createdAt_gt: Int
  createdAt_lt: Int
  createdAt_gte: Int
  createdAt_lte: Int
  createdAt_in: [Int!]
  deletedBy: ID
  deletedBy_ne: ID
  deletedBy_gt: ID
  deletedBy_lt: ID
  deletedBy_gte: ID
  deletedBy_lte: ID
  deletedBy_in: [ID!]
  updatedBy: ID
  updatedBy_ne: ID
  updatedBy_gt: ID
  updatedBy_lt: ID
  updatedBy_gte: ID
  updatedBy_lte: ID
  updatedBy_in: [ID!]
  createdBy: ID
  createdBy_ne: ID
  createdBy_gt: ID
  createdBy_lt: ID
  createdBy_gte: ID
  createdBy_lte: ID
  createdBy_in: [ID!]
  groups: GroupFilterType
  roles: RoleFilterType
}

type AdminResultType {
  data: [Admin!]!
  total: Int!
  current_page: Int!
  per_page: Int!
  total_page: Int!
}

input GroupCreateInput {
  id: ID
  Name: String
  state: Int
  del: Int
  adminIds: [ID!]
  rolesIds: [ID!]
}

input GroupUpdateInput {
  Name: String
  state: Int
  del: Int
  adminIds: [ID!]
  rolesIds: [ID!]
}

enum GroupSortType {
  ID_ASC
  ID_DESC
  NAME_ASC
  NAME_DESC
  STATE_ASC
  STATE_DESC
  DEL_ASC
  DEL_DESC
  UPDATED_AT_ASC
  UPDATED_AT_DESC
  CREATED_AT_ASC
  CREATED_AT_DESC
  DELETED_BY_ASC
  DELETED_BY_DESC
  UPDATED_BY_ASC
  UPDATED_BY_DESC
  CREATED_BY_ASC
  CREATED_BY_DESC
  ADMIN_IDS_ASC
  ADMIN_IDS_DESC
  ROLES_IDS_ASC
  ROLES_IDS_DESC
}

input GroupFilterType {
  AND: [GroupFilterType!]
  OR: [GroupFilterType!]
  id: ID
  id_ne: ID
  id_gt: ID
  id_lt: ID
  id_gte: ID
  id_lte: ID
  id_in: [ID!]
  Name: String
  Name_ne: String
  Name_gt: String
  Name_lt: String
  Name_gte: String
  Name_lte: String
  Name_in: [String!]
  Name_like: String
  Name_prefix: String
  Name_suffix: String
  state: Int
  state_ne: Int
  state_gt: Int
  state_lt: Int
  state_gte: Int
  state_lte: Int
  state_in: [Int!]
  del: Int
  del_ne: Int
  del_gt: Int
  del_lt: Int
  del_gte: Int
  del_lte: Int
  del_in: [Int!]
  updatedAt: Int
  updatedAt_ne: Int
  updatedAt_gt: Int
  updatedAt_lt: Int
  updatedAt_gte: Int
  updatedAt_lte: Int
  updatedAt_in: [Int!]
  createdAt: Int
  createdAt_ne: Int
  createdAt_gt: Int
  createdAt_lt: Int
  createdAt_gte: Int
  createdAt_lte: Int
  createdAt_in: [Int!]
  deletedBy: ID
  deletedBy_ne: ID
  deletedBy_gt: ID
  deletedBy_lt: ID
  deletedBy_gte: ID
  deletedBy_lte: ID
  deletedBy_in: [ID!]
  updatedBy: ID
  updatedBy_ne: ID
  updatedBy_gt: ID
  updatedBy_lt: ID
  updatedBy_gte: ID
  updatedBy_lte: ID
  updatedBy_in: [ID!]
  createdBy: ID
  createdBy_ne: ID
  createdBy_gt: ID
  createdBy_lt: ID
  createdBy_gte: ID
  createdBy_lte: ID
  createdBy_in: [ID!]
  admin: AdminFilterType
  roles: RoleFilterType
}

type GroupResultType {
  data: [Group!]!
  total: Int!
  current_page: Int!
  per_page: Int!
  total_page: Int!
}

input RoleCreateInput {
  id: ID
  Name: String
  Pid: String
  state: Int
  del: Int
  adminIds: [ID!]
  groupIds: [ID!]
}

input RoleUpdateInput {
  Name: String
  Pid: String
  state: Int
  del: Int
  adminIds: [ID!]
  groupIds: [ID!]
}

enum RoleSortType {
  ID_ASC
  ID_DESC
  NAME_ASC
  NAME_DESC
  PID_ASC
  PID_DESC
  STATE_ASC
  STATE_DESC
  DEL_ASC
  DEL_DESC
  UPDATED_AT_ASC
  UPDATED_AT_DESC
  CREATED_AT_ASC
  CREATED_AT_DESC
  DELETED_BY_ASC
  DELETED_BY_DESC
  UPDATED_BY_ASC
  UPDATED_BY_DESC
  CREATED_BY_ASC
  CREATED_BY_DESC
  ADMIN_IDS_ASC
  ADMIN_IDS_DESC
  GROUP_IDS_ASC
  GROUP_IDS_DESC
}

input RoleFilterType {
  AND: [RoleFilterType!]
  OR: [RoleFilterType!]
  id: ID
  id_ne: ID
  id_gt: ID
  id_lt: ID
  id_gte: ID
  id_lte: ID
  id_in: [ID!]
  Name: String
  Name_ne: String
  Name_gt: String
  Name_lt: String
  Name_gte: String
  Name_lte: String
  Name_in: [String!]
  Name_like: String
  Name_prefix: String
  Name_suffix: String
  Pid: String
  Pid_ne: String
  Pid_gt: String
  Pid_lt: String
  Pid_gte: String
  Pid_lte: String
  Pid_in: [String!]
  Pid_like: String
  Pid_prefix: String
  Pid_suffix: String
  state: Int
  state_ne: Int
  state_gt: Int
  state_lt: Int
  state_gte: Int
  state_lte: Int
  state_in: [Int!]
  del: Int
  del_ne: Int
  del_gt: Int
  del_lt: Int
  del_gte: Int
  del_lte: Int
  del_in: [Int!]
  updatedAt: Int
  updatedAt_ne: Int
  updatedAt_gt: Int
  updatedAt_lt: Int
  updatedAt_gte: Int
  updatedAt_lte: Int
  updatedAt_in: [Int!]
  createdAt: Int
  createdAt_ne: Int
  createdAt_gt: Int
  createdAt_lt: Int
  createdAt_gte: Int
  createdAt_lte: Int
  createdAt_in: [Int!]
  deletedBy: ID
  deletedBy_ne: ID
  deletedBy_gt: ID
  deletedBy_lt: ID
  deletedBy_gte: ID
  deletedBy_lte: ID
  deletedBy_in: [ID!]
  updatedBy: ID
  updatedBy_ne: ID
  updatedBy_gt: ID
  updatedBy_lt: ID
  updatedBy_gte: ID
  updatedBy_lte: ID
  updatedBy_in: [ID!]
  createdBy: ID
  createdBy_ne: ID
  createdBy_gt: ID
  createdBy_lt: ID
  createdBy_gte: ID
  createdBy_lte: ID
  createdBy_in: [ID!]
  admin: AdminFilterType
  group: AdminFilterType
}

type RoleResultType {
  data: [Role!]!
  total: Int!
  current_page: Int!
  per_page: Int!
  total_page: Int!
}`
)
