package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var dataConfigFile string

type Recipe struct {
	Connection ConnectionConfig `yaml:"connection"`
	Records    []Record         `yaml:"records"`
}

type ConnectionConfig struct {
	DSN      string `yaml:"dsn"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
	SSLMode  string `yaml:"sslmode"`
}

type Record struct {
	ID     string           `yaml:"id"`
	Table  string           `yaml:"table"`
	Count  int              `yaml:"count"`
	Fields map[string]Field `yaml:"fields"`
}

type Field struct {
	Type            string   `yaml:"type"`
	Value           string   `yaml:"value"`
	Unique          bool     `yaml:"unique"`
	Nullable        bool     `yaml:"nullable"`
	NullProbability float64  `yaml:"null_probability"`
	Pattern         string   `yaml:"pattern"`
	Min             int      `yaml:"min"`
	Max             int      `yaml:"max"`
	Values          []string `yaml:"values"`
	From            string   `yaml:"from"`
	Field           string   `yaml:"field"`
	Start           string   `yaml:"start"`
	End             string   `yaml:"end"`
	Format          string   `yaml:"format"`
}

type Engine struct {
	recipe        Recipe
	dbConnection  *pgx.Conn
	generatedData map[string][]map[string]interface{}
}

func NewEngine(recipe Recipe, dbConnection *pgx.Conn) *Engine {
	return &Engine{
		recipe:        recipe,
		dbConnection:  dbConnection,
		generatedData: make(map[string][]map[string]interface{}),
	}
}

func (e *Engine) Generate(ctx context.Context) error {
	for _, record := range e.recipe.Records {
		fmt.Printf("Generating %d records for table %s...\n", record.Count, record.Table)
	}
	return nil
}

var generateDataCmd = &cobra.Command{
	Use:   "generate-data",
	Short: "Generates and inserts test data into a database based on a recipe file",
	Long: `Reads a YAML recipe file to connect to a PostgreSQL database, generate realistic fake data, and insert it
into the specified tables. This is ideal for setting up a clean state for your integration or E2E tests.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("Reading data recipe file: %s\n", dataConfigFile)
		yamlFile, err := os.ReadFile(dataConfigFile)
		if err != nil {
			return fmt.Errorf("failed to read recipe file: %w", err)
		}

		var recipe Recipe

		err = yaml.Unmarshal(yamlFile, &recipe)
		if err != nil {
			return fmt.Errorf("failed to parse YAML recipe file: %w", err)
		}

		fmt.Println("Recipe file successfully parsed!")
		fmt.Println("\n--- Parsed Recipe ---")
		spew.Dump(recipe)
		fmt.Println("----------------------")
		return nil
	},
}

func (e *Engine) generateFieldValue(field Field, recordIndex int) (interface{}, error) {
	switch field.Type {
	case "uuid":
		return uuid.New().String(), nil
	case "faker":
		return "fake_value", nil
	case "reference":
		return "referenced_value", nil
	case "sequence":
		return "sequence_value", nil

	default:
		return nil, fmt.Errorf("unknown field type: %s, field.Type")
	}
}

func init() {
	rootCmd.AddCommand(generateDataCmd)

	generateDataCmd.Flags().StringVarP(&dataConfigFile, "file", "f", "", "Path to the data recipe YAML file (required)")
	if err := generateDataCmd.MarkFlagRequired("file"); err != nil {
		panic(err)
	}
}
