package constant

import "github.com/tanveerprottoy/stdlib-go-template/internal/pkg/types"

const ApiPattern = "/api"
const RootPattern = "/"
const V1 = "/v1"
const UsersPattern = "/users"
const ContentsPattern = "/contents"
const FilesPattern = "/files"

// db
const RowsAffected = "rowsAffected"

// misc
const InternalServerError = "internal server error"
const BadRequest = "bad request"
const Unauthorized = "unauthorized"
const OperationNotSuccess = "operation was not successful"

const Error = "error"
const Errors = "errors"

// basic keys
const KeyId = "id"
const KeyPage = "page"
const KeyLimit = "limit"

// context keys
const KeyAuthData types.KeyContext = "AuthData"
const KeyAuthUser types.KeyContext = "AuthUser"
const KeyRBAC types.KeyContext = "rbac"

// remote userservice auth endpoint
const UserServiceAuthEndpoint = "/api/v2/auth/get-user"

// auth data keys
const AuthUser = "authUser"
const AuthRole = "authRole"

const RBACError = "rbac error"
