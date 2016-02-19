package config

const DatabaseUser 				= `postgres`
const DatabasePassword 		= `1234`
const DatabaseName 				= `kart_development`
const SslMode 						= `disable`

func GetDbConnectionString() string {
	return `user=` + DatabaseUser +
				` password=` + DatabasePassword +
				` dbname=` + DatabaseName +
				` sslmode=` + SslMode
}