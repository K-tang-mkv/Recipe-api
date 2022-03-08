## A distributed recipe api based on gin framework with microservice

### **fundament tech**
* database - mongodb
* cache service - redis
* framework - gin
* distributed arch - docker 

### **data model**
| Recipe | ID | Name | Tags | Ingredients | Instructions | PublishedAt | 
| --- | --- | --- | --- | --- | --- | --- | 
| **type** | primitive.ObjectID | string | []string | []string | []string | time.Time |
| **json** | id | name | tags | ingredients | instructions | publishedAt |
| **bson** | _id | name | tags | ingredients | instructions | publishedAt |

```golang
type Recipe struct {
	// swagger:ignore
	ID           primitive.ObjectID `json:"id" bson:"_id"`
	Name         string             `json:"name" bson:"name"`
	Tags         []string           `json:"tags" bson:"tags"`
	Ingredients  []string           `json:"ingredients" bson:"ingredients"`
	Instructions []string           `json:"instructions" bson:"instructions"`
	PublishedAt  time.Time          `json:"publishedAt" bson:"publishedAt"`
}

```

| User | Username | Password |
| --- | --- | --- |
| **type** | string | string |
| **json** | usename | password |

```golang
type User struct {
	Password string `json:"password"`
	Username string `json:"username"`
}
```

### Reference
**gin** - https://github.com/gin-gonic/gin

**book** - https://github.com/PacktPublishing/Building-Distributed-Applications-in-Gin

