package structures

// All struct we will add here

type DBConnect struct {
	DBName     string
	DBUser     string
	DBPassword string
	DBPort     string
	DBHost     string
}

type Blog struct {
	BlogId          string
	Title           string
	Content         string
	Author          string
	PublicationDate string
	Tags            string
}
