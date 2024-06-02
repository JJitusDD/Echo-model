package configs

type Config struct {
	Port           string         `json:"port" mapstructure:"port"`
	ProjectRootDir string         `json:"project_root_dir" mapstructure:"project_root_dir"`
	Kafka          KafkaConfig    `json:"kafka" mapstructure:"kafka"`
	Redis          RedisConfig    `json:"redis" mapstructure:"redis"`
	Postgres       PostgresConfig `json:"postgres" mapstructure:"postgres"`
	ScheduleJob    string         `json:"schedule_job" mapstructure:"schedule_job"`
	TemporalConfig TemporalConfig `json:"temporal_config" mapstructure:"temporal_config"`
	SFTP           SFTPConfig     `json:"sftp" mapstructure:"sftp"`
	ENV            string         `json:"env" mapstructure:"env"`
}

type SFTPConfig struct {
	SftpAddr string `json:"sftp_addr" mapstructure:"sftp_addr"`
	SftpUser string `json:"sftp_user" mapstructure:"sftp_user"`
	SftpPass string `json:"sftp_pass" mapstructure:"sftp_pass"`
}

type KafkaConfig struct {
	Server    string `json:"server" mapstructure:"server"`
	Username  string `json:"username" mapstructure:"username"`
	Password  string `json:"password" mapstructure:"password"`
	Protocol  string `json:"protocol" mapstructure:"protocol"`
	Mechanism string `json:"mechanism" mapstructure:"mechanism"`
	GroupID   string `json:"group_id" mapstructure:"group_id"`
}
type RedisConfig struct {
	Addr     string `json:"addr" mapstructure:"addr"`
	Password string `json:"password" mapstructure:"password"`
	DB       int    `json:"db" mapstructure:"db"`
}

type PostgresConfig struct {
	Host      string `json:"host" mapstructure:"host"`
	Port      int    `json:"port" mapstructure:"port"`
	DbName    string `json:"dbname" mapstructure:"dbname"`
	User      string `json:"user" mapstructure:"user"`
	Pass      string `json:"password" mapstructure:"password"`
	SSLMode   string `json:"sslmode" mapstructure:"sslmode"`
	Prefix    string `json:"prefix" mapstructure:"prefix"`
	DebugMode bool   `json:"debug_mode" mapstructure:"debug_mode"`
}

type TemporalConfig struct {
	Uri                              string `json:"uri" mapstructure:"uri"`
	Namespace                        string `json:"namespace" mapstructure:"namespace"`
	MaxConcurrentWorkflowTaskPollers int    `json:"max_concurrent_workflow_task_pollers" mapstructure:"max_concurrent_workflow_task_pollers"`
	MaxConcurrentActivityTaskPollers int    `json:"max_concurrent_activity_task_pollers" mapstructure:"max_concurrent_activity_task_pollers"`
}
