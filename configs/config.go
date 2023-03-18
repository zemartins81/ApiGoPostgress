package configs

import "github.com/spf13/viper" //Importa a biblioteca Viper para trabalhar com a configuração

var cfg *config //Define um ponteiro para a estrutura config

type config struct { //Define a estrutura config, que contém duas estruturas embutidas APIConfig e DBConfig
	API APIConfig
	DB  DBConfig
}

type APIConfig struct { //Estrutura para armazenar a configuração da porta do servidor API
	Port string
}

type DBConfig struct { //Estrutura para armazenar a configuração do banco de dados
	Host     string
	Port     string
	User     string
	Pass     string
	Database string
}

func init() { //Função init é executada antes da função main() e configura os valores padrão para as configurações
	viper.SetDefault("api.port", "9000")           //Define a porta padrão do servidor como 9000
	viper.SetDefault("database.host", "localhost") //Define o host padrão do banco de dados como localhost
	viper.SetDefault("database.port", "5432")      //Define a porta padrão do banco de dados como 5432
}

func Load() error { //Carrega as configurações a partir do arquivo de configuração TOML
	viper.SetConfigName("config") //Define o nome do arquivo de configuração como "config.toml"
	viper.SetConfigType("toml")   //Define o tipo de arquivo de configuração como TOML
	viper.AddConfigPath(".")      //Adiciona o diretório atual como um caminho de pesquisa para o arquivo de configuração

	err := viper.ReadInConfig() //Lê o arquivo de configuração TOML
	if err != nil {             //Verifica se houve um erro durante a leitura do arquivo de configuração
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok { //Se o arquivo de configuração não foi encontrado, o erro é ignorado e os valores padrão são usados
			return err //Se houver outro erro ao ler o arquivo de configuração, retorna um erro
		}
	}
	cfg = new(config)    //Inicializa uma nova estrutura config
	cfg.API = APIConfig{ //Define as configurações da estrutura APIConfig a partir das configurações do arquivo de configuração
		Port: viper.GetString("api.port"),
	}
	cfg.DB = DBConfig{ //Define as configurações da estrutura DBConfig a partir das configurações do arquivo de configuração
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.port"),
		User:     viper.GetString("database.user"),
		Pass:     viper.GetString("database.pass"),
		Database: viper.GetString("database.name"),
	}

	return nil //Retorna nil se não houver erro
}

func GetDB() DBConfig { //Retorna a estrutura DBConfig contendo as configurações do banco de dados
	return cfg.DB
}

func GetServerPort() string { //Retorna a porta do servidor
	return cfg.API.Port
}
