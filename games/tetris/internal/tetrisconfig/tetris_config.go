package tetrisconfig

type Config struct {
	NextQueueLength int
	GhostEnabled    bool
	MaxLevel        int
	EndOnMaxLevel   bool
	Theme           *Theme
	Keys            *TetrisKeys
}

func CreateConfig() *Config {
	config := Config{
		NextQueueLength: 5,
		GhostEnabled:    true,
		MaxLevel:        15,
		EndOnMaxLevel:   false,
		Theme:           CreateTetrisTheme(),
		Keys:            SetTetrisKeyBindings(),
	}

	return &config
}
