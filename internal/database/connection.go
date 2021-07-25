package database

import (
	"git.pesca.dev/pesca-dev/moneyboy-backend/internal/models"
	"gorm.io/gorm"
)

type Connection struct {
	dialector      gorm.Dialector
	opts           []gorm.Option
	db             *gorm.DB
	userConnection *UserConnection
}

// Create a new connection, but do not connect
func Create(dialector gorm.Dialector, opts ...gorm.Option) *Connection {
	return &Connection{dialector: dialector, opts: opts, db: nil}
}

// Connect to the database
func (c *Connection) Connect() error {
	if c.db != nil {
		panic("already connected to a database")
	}
	db, err := gorm.Open(c.dialector, c.opts...)
	c.db = db
	if err == nil {
		db.AutoMigrate(&models.User{})
		c.init()
	}
	return err
}

func (c *Connection) init() {
	c.userConnection = createUserConnection(c.db)
}

// Query users
func (c *Connection) Users() *UserConnection {
	return c.userConnection
}
