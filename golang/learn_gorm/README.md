https://app.pluralsight.com/library/courses/gorm-go-object-relational-mapper/table-of-contents

### GORM's features

- CRUD
- Schema management
- Transactions
- Migrations
- Event Hooks
- Chainable API
- Logger

```bash
go get -u github.com/jinzhu/gorm
go get -u github.com/lib/pq
```

```golang
type User struct {
  ID uint
  FirstName string
  LastName string
}
```

```
tablename: users
fields:
  id serial
  first_name varchar(255)
  last_name varchar(255)
```

### Customizing table creation

#### Method-1

```golang
db.SingularTable(true)
```

#### Method-2

```golang
type User struct {
  ID uint
  FirstName string
  LastName string
}

func (u User) TableName() string {
  return "stakeholders"
}
```

### Gorm model

```golang
model := gorm.Model()

type User struct {
  gorm.Model
  UserName string
}
```

### Field types

```golang
type User struct {
  gorm.Model
  Username string `sql:"type:VARCHAR(15); not null"`
  FirstName string `sql:"size:100"`
  LastName string `sql:"unique;unique_index;not null;DEFAULT: 'Smith'"`
  Count int `gorm:"AUTO_INCREMENT"`
  TempField bool `sql:"-"`
}
```

Primary keys

```golang
type User struct {
  UserID int `gorm:"primary_key"`
  Username string `sql:"type:VARCHAR(15); not null"`
  FirstName string `sql:"size:100"`
  LastName string `sql:"unique;unique_index;not null;DEFAULT: 'Smith'"`
  Count int `gorm:"AUTO_INCREMENT"`
  TempField bool `sql:"-"`
}
```

Customizing column names

```golang
  UserID int `gorm:"primary_key"`
  Username string `sql:"type:VARCHAR(15); not null"`
  FirstName string `sql:"size:100" gorm:"column:FirstName"`
  LastName string `sql:"unique;unique_index;not null;DEFAULT: 'Smith'" gorm:"column:LastName"`
  Count int `gorm:"AUTO_INCREMENT"`
  TempField bool `sql:"-"`
```

Embedding child objects

```golang

type User struct {
  Model gorm.Model `gorm:"embedded"`
  FirstName string
  LastName string
}

for _, f:= range db.NewScope(&User{}).Firlds() {
  println(f.Name)
}
```

Indexes

```golang
db.CreateTable(&User{})
// name of the field as db sees.
// It is first_name not FirstName
db.Model(&User{}).AddIndex("idx_first_name", "first_name")
db.Model(&User{}).AddUniqueIndex("idx_last_name", "last_name")
type User struct {
  Model gorm.Model
  FirstName string
  LastName string
}
// deleting indexes
db.Model(&User{}).RemoveIndex("idx_first_name")
```

![gorm tag summary]("https://github.com/goutham2027/notes/blob/master/golang/learn_gorm/gorm-tag-summary.png")

### Relationships/Associations

- one-to-one
- foreign keys
- one-to-many
- many-to-many
- polymorphism

```golang
db.Debug().Model(&Calendar{}).AddForeignKey(
  "user_id", "users(id)", "CASCADE", "CASCADE"
)
// user has-one calendar
type User struct {
  gorm.Model
  Username string
  FirstName string
  LastName string
  Calendar Calendar
}

type Calendar struct {
  gorm.Model
  UserID uint
}

db.Debug().Save(
  &User{
    Username: "adent",
    Calendar: Calendar{
      Name: "Improbable events",
    },
  }
)

u := User{}
c := Calendar{}
db.First(&u).Related(&c, "calendar")


// user belongs to calendar
// doesn't make sense
type User struct {
  gorm.Model
  Username string
  FirstName string
  LastName string
  Calendar Calendar
  CalendarId uint
}

type Calendar struct {
  gorm.Model
}
```

#### One-to-Many

```golang
# calendar has many appointments
type Calendar struct {
  gorm.Model
  Name string
  Appointments []Appointment
}

type Appointment struct {
  gorm.Model
  CalendarID uint
}

db.Debug().Save(&User{
  Username: "adent",
  Calendar: Calendar {
    Name: "Improbable Events",
    Appointments: []Appointment {
      {Subject: "Spontaneous Whale Generation"},
      {Subject: "Saved from Vaccuum of Space"},
    }
  }
})
```

#### Many to Many relationship

```golang
type User struct {
  gorm.Model
  Username string
}

type Calendar struct {
  gorm.Model
  Name string
  UserID uint
  Appointments []Appointment
}

type Appointment struct {
  gorm.Model
  Subject string
  Description string
  CalendarID uint
  Attendees []User `gorm:"many2many:appointment_user"`
}

users := []User {
  {Username: "fprefect"},
  {Username: "tmacmillan"},
  {Username: "mrobot"},
}

for i := range users {
  db.Save(&users[i])
}


db.Debug().Save(&User{
  Username: "adent",
  Calendar: Calendar{
    Name: "Improbable events",
    Appointments: []Appointment{
      {Subject: "Spontaneous Whale Generation", Attendees: users},
      {Subject: "Saved from Vaccuum of space", Attendees: users},
    }
  }
})
```

#### Polymorphism

```golang
type Cat struct {
  gorm.Model
  Toy Toy `gorm:"polymorphic:owner"`
}


type Dog struct {
  form.Model
  Toy Toy `gorm:"polymorphic:owner"`
}

type Toy struct {
  gorm.Model
  Type string
  OwnerID uint
  OwnerType string
}
```

#### Methods from db object

```golang
db.Model(&Calendar{}).Association("Appointments")

.Find(&appointments)
.Append(&appointemntsToAdd)
.Delete(&appointemntsToDelete)
.Replace(&appointemntsToSubstitute)
.Count()
.Clear()
```

### Creating, Updating and Deleting records
