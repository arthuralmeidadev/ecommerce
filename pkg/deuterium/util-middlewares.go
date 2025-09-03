package deuterium

type CorsConfig struct {
	Origin               string
	Methods              []string
	Credentials          bool
	MaxAge               uint
	PreflightContinue    bool
	OptionsSuccessStatus uint
}
