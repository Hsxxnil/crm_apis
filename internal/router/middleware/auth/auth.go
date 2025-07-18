package auth

import (
	"net/http"

	roleModel "crm/internal/interactor/models/roles"
	"crm/internal/interactor/service/role"

	"github.com/casbin/casbin/v2/model"
	"gorm.io/gorm"

	_ "crm/internal/interactor/pkg/connect"
	"crm/internal/interactor/pkg/util/log"

	"github.com/casbin/casbin/v2"
	jsonAdapter "github.com/casbin/json-adapter/v2"
	"github.com/gin-gonic/gin"
	_ "gorm.io/driver/postgres"
)

// CasbinBind struct is used to create or delete the policy rule from front.
type CasbinBind struct {
	Ptype    string `json:"ptype" binding:"required" validate:"required"`
	RoleName string `json:"role_name" binding:"required" validate:"required"`
	Path     string `json:"path" binding:"required" validate:"required"`
	Method   string `json:"method" binding:"required" validate:"required"`
}

// CasbinModel includes CasbinBind and an automatically generated id.
type CasbinModel struct {
	ID int `json:"id"`
	CasbinBind
}

// CasbinOutput is used to return all policies.
type CasbinOutput struct {
	RoleName string `json:"role_name"`
	Path     string `json:"path"`
	Method   string `json:"method"`
}

func adapter() *jsonAdapter.Adapter {
	//db, err := connect.PostgresSQL()
	//if err != nil {
	//	panic(err)
	//}

	//a, err := gormAdapter.NewAdapterByDB(db)
	//if err != nil {
	//	panic(err)
	//}

	cb := []byte(_CASBIN_RULES)

	a := jsonAdapter.NewAdapter(&cb)

	return a
}

func Casbin(a *jsonAdapter.Adapter) *casbin.Enforcer {
	m, err := model.NewModelFromString(`[request_definition]
	r = sub, obj, act

	[policy_definition]
	p = sub, obj, act

	[policy_effect]
	e = some(where (p.eft == allow))

	[matchers]
	m = r.sub == p.sub && keyMatch(r.obj, p.obj) && regexMatch(r.act, p.act)

	#[matchers]
	#m = r.sub == p.sub && r.obj == p.obj && r.act == p.act`)
	if err != nil {
		log.Info("model error: %s;", err)
	}

	e, err := casbin.NewEnforcer(m, a)
	if err != nil {
		panic(err)
	}

	return e
}

var Enforcer *casbin.Enforcer

func init() {
	a := adapter()
	Enforcer = Casbin(a)
}

func AddPolicy(cm CasbinModel) (bool, error) {
	return Enforcer.AddPolicy(cm.RoleName, cm.Path, cm.Method)
}

func DeletePolicy(cm CasbinModel) (bool, error) {
	return Enforcer.RemovePolicy(cm.RoleName, cm.Path, cm.Method)
}

func GetAllPolicies() [][]string {
	return Enforcer.GetPolicy()
}

func AuthCheckRole(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		checkRole, err := role.Init(db).GetBySingle(&roleModel.Field{
			RoleID: c.MustGet("role_id").(string),
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": -1,
				"msg":    err.Error(),
			})
			c.Abort()
			return
		}

		// 去除path中的uuid
		//array := strings.Split(c.Request.URL.Path, "/")
		//var path string
		//if _, err := uuid.Parse(array[len(array)-1]); err == nil {
		//	path = strings.Join(array[:len(array)-1], "/")
		//} else {
		//	path = c.Request.URL.Path
		//}

		log.Info("Casbin policy:", *checkRole.Name, c.Request.URL.Path, c.Request.Method)
		res, err := Enforcer.Enforce(*checkRole.Name, c.Request.URL.Path, c.Request.Method)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": -1,
				"msg":    err.Error(),
			})
			c.Abort()
			return
		}

		if res {
			c.Next()
		} else {
			c.JSON(http.StatusNonAuthoritativeInfo, gin.H{
				"status": 203,
				"msg":    "Sorry, you don't have permission.",
			})
			c.Abort()
			return
		}
	}
}
