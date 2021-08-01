package database

import (
	"git.pesca.dev/pesca-dev/moneyboy-backend/internal/models"
	"gorm.io/gorm"
)

type Connection struct {
	dialector         gorm.Dialector
	opts              []gorm.Option
	db                *gorm.DB
	userConnection    *UserConnection
	sessionConnection *SessionConnection
}

func New(dialector gorm.Dialector, opts ...gorm.Option) *Connection {
	con := &Connection{dialector: dialector, opts: opts, db: nil}
	if err := con.connect(); err != nil {
		panic("cannot connect to database")
	}
	return con
}

// Connect to the database
func (c *Connection) connect() error {
	if c.db != nil {
		panic("already connected to a database")
	}
	db, err := gorm.Open(c.dialector, c.opts...)
	c.db = db
	if err == nil {
		db.AutoMigrate(&models.Session{}, &models.User{})
		c.init()
	}
	return err
}

func (c *Connection) init() {
	c.userConnection = createUserConnection(c.db)
	c.sessionConnection = createSessionConnection(c.db)
}

// Query users
func (c *Connection) Users() *UserConnection {
	return c.userConnection
}

func (c *Connection) Sessions() *SessionConnection {
	return c.sessionConnection
}
