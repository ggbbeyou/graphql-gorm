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
}

type User {
  id: ID!
  email: String
  firstName: String
  lastName: String
  tasks: [Task!]!
  state: Int
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
  updatedAt: Int
  createdAt: Int
  deletedBy: ID
  updatedBy: ID
  createdBy: ID
}

extend type Mutation {
  token(input: String): String!
}

input UserCreateInput {
  id: ID
  email: String
  firstName: String
  lastName: String
  state: Int
  tasksIds: [ID!]
}

input UserUpdateInput {
  email: String
  firstName: String
  lastName: String
  state: Int
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
}

input TaskUpdateInput {
  title: String
  completed: Boolean
  dueDate: Time
  assigneeId: ID
  state: Int
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
}`
)