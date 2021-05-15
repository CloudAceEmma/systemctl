package properties

type Property string

const (
	ActiveEnterTimestamp            Property = "ActiveEnterTimestamp"
	ActiveEnterTimestampMonotonic   Property = "ActiveEnterTimestampMonotonic"
	ActiveExitTimestampMonotonic    Property = "ActiveExitTimestampMonotonic"
	ActiveState                     Property = "ActiveState"
	After                           Property = "After"
	AllowIsolate                    Property = "AllowIsolate"
	AssertResult                    Property = "AssertResult"
	AssertTimestamp                 Property = "AssertTimestamp"
	AssertTimestampMonotonic        Property = "AssertTimestampMonotonic"
	Before                          Property = "Before"
	BlockIOAccounting               Property = "BlockIOAccounting"
	BlockIOWeight                   Property = "BlockIOWeight"
	CPUAccounting                   Property = "CPUAccounting"
	CPUAffinityFromNUMA             Property = "CPUAffinityFromNUMA"
	CPUQuotaPerSecUSec              Property = "CPUQuotaPerSecUSec"
	CPUQuotaPeriodUSec              Property = "CPUQuotaPeriodUSec"
	CPUSchedulingPolicy             Property = "CPUSchedulingPolicy"
	CPUSchedulingPriority           Property = "CPUSchedulingPriority"
	CPUSchedulingResetOnFork        Property = "CPUSchedulingResetOnFork"
	CPUShares                       Property = "CPUShares"
	CPUUsageNSec                    Property = "CPUUsageNSec"
	CPUWeight                       Property = "CPUWeight"
	CacheDirectoryMode              Property = "CacheDirectoryMode"
	CanFreeze                       Property = "CanFreeze"
	CanIsolate                      Property = "CanIsolate"
	CanReload                       Property = "CanReload"
	CanStart                        Property = "CanStart"
	CanStop                         Property = "CanStop"
	CapabilityBoundingSet           Property = "CapabilityBoundingSet"
	CleanResult                     Property = "CleanResult"
	CollectMode                     Property = "CollectMode"
	ConditionResult                 Property = "ConditionResult"
	ConditionTimestamp              Property = "ConditionTimestamp"
	ConditionTimestampMonotonic     Property = "ConditionTimestampMonotonic"
	ConfigurationDirectoryMode      Property = "ConfigurationDirectoryMode"
	Conflicts                       Property = "Conflicts"
	ControlGroup                    Property = "ControlGroup"
	ControlPID                      Property = "ControlPID"
	CoredumpFilter                  Property = "CoredumpFilter"
	DefaultDependencies             Property = "DefaultDependencies"
	DefaultMemoryLow                Property = "DefaultMemoryLow"
	DefaultMemoryMin                Property = "DefaultMemoryMin"
	Delegate                        Property = "Delegate"
	Description                     Property = "Description"
	DevicePolicy                    Property = "DevicePolicy"
	DynamicUser                     Property = "DynamicUser"
	EffectiveCPUs                   Property = "EffectiveCPUs"
	EffectiveMemoryNodes            Property = "EffectiveMemoryNodes"
	ExecMainCode                    Property = "ExecMainCode"
	ExecMainExitTimestampMonotonic  Property = "ExecMainExitTimestampMonotonic"
	ExecMainPID                     Property = "ExecMainPID"
	ExecMainStartTimestamp          Property = "ExecMainStartTimestamp"
	ExecMainStartTimestampMonotonic Property = "ExecMainStartTimestampMonotonic"
	ExecMainStatus                  Property = "ExecMainStatus"
	ExecReload                      Property = "ExecReload"
	ExecReloadEx                    Property = "ExecReloadEx"
	ExecStart                       Property = "ExecStart"
	ExecStartEx                     Property = "ExecStartEx"
	FailureAction                   Property = "FailureAction"
	FileDescriptorStoreMax          Property = "FileDescriptorStoreMax"
	FinalKillSignal                 Property = "FinalKillSignal"
	FragmentPath                    Property = "FragmentPath"
	FreezerState                    Property = "FreezerState"
	GID                             Property = "GID"
	GuessMainPID                    Property = "GuessMainPID"
	IOAccounting                    Property = "IOAccounting"
	IOReadBytes                     Property = "IOReadBytes"
	IOReadOperations                Property = "IOReadOperations"
	IOSchedulingClass               Property = "IOSchedulingClass"
	IOSchedulingPriority            Property = "IOSchedulingPriority"
	IOWeight                        Property = "IOWeight"
	IOWriteBytes                    Property = "IOWriteBytes"
	IOWriteOperations               Property = "IOWriteOperations"
	IPAccounting                    Property = "IPAccounting"
	IPEgressBytes                   Property = "IPEgressBytes"
	IPEgressPackets                 Property = "IPEgressPackets"
	IPIngressBytes                  Property = "IPIngressBytes"
	IPIngressPackets                Property = "IPIngressPackets"
	Id                              Property = "Id"
	IgnoreOnIsolate                 Property = "IgnoreOnIsolate"
	IgnoreSIGPIPE                   Property = "IgnoreSIGPIPE"
	InactiveEnterTimestampMonotonic Property = "InactiveEnterTimestampMonotonic"
	InactiveExitTimestamp           Property = "InactiveExitTimestamp"
	InactiveExitTimestampMonotonic  Property = "InactiveExitTimestampMonotonic"
	InvocationID                    Property = "InvocationID"
	JobRunningTimeoutUSec           Property = "JobRunningTimeoutUSec"
	JobTimeoutAction                Property = "JobTimeoutAction"
	JobTimeoutUSec                  Property = "JobTimeoutUSec"
	KeyringMode                     Property = "KeyringMode"
	KillMode                        Property = "KillMode"
	KillSignal                      Property = "KillSignal"
	LimitAS                         Property = "LimitAS"
	LimitASSoft                     Property = "LimitASSoft"
	LimitCORE                       Property = "LimitCORE"
	LimitCORESoft                   Property = "LimitCORESoft"
	LimitCPU                        Property = "LimitCPU"
	LimitCPUSoft                    Property = "LimitCPUSoft"
	LimitDATA                       Property = "LimitDATA"
	LimitDATASoft                   Property = "LimitDATASoft"
	LimitFSIZE                      Property = "LimitFSIZE"
	LimitFSIZESoft                  Property = "LimitFSIZESoft"
	LimitLOCKS                      Property = "LimitLOCKS"
	LimitLOCKSSoft                  Property = "LimitLOCKSSoft"
	LimitMEMLOCK                    Property = "LimitMEMLOCK"
	LimitMEMLOCKSoft                Property = "LimitMEMLOCKSoft"
	LimitMSGQUEUE                   Property = "LimitMSGQUEUE"
	LimitMSGQUEUESoft               Property = "LimitMSGQUEUESoft"
	LimitNICE                       Property = "LimitNICE"
	LimitNICESoft                   Property = "LimitNICESoft"
	LimitNOFILE                     Property = "LimitNOFILE"
	LimitNOFILESoft                 Property = "LimitNOFILESoft"
	LimitNPROC                      Property = "LimitNPROC"
	LimitNPROCSoft                  Property = "LimitNPROCSoft"
	LimitRSS                        Property = "LimitRSS"
	LimitRSSSoft                    Property = "LimitRSSSoft"
	LimitRTPRIO                     Property = "LimitRTPRIO"
	LimitRTPRIOSoft                 Property = "LimitRTPRIOSoft"
	LimitRTTIME                     Property = "LimitRTTIME"
	LimitRTTIMESoft                 Property = "LimitRTTIMESoft"
	LimitSIGPENDING                 Property = "LimitSIGPENDING"
	LimitSIGPENDINGSoft             Property = "LimitSIGPENDINGSoft"
	LimitSTACK                      Property = "LimitSTACK"
	LimitSTACKSoft                  Property = "LimitSTACKSoft"
	LoadState                       Property = "LoadState"
	LockPersonality                 Property = "LockPersonality"
	LogLevelMax                     Property = "LogLevelMax"
	LogRateLimitBurst               Property = "LogRateLimitBurst"
	LogRateLimitIntervalUSec        Property = "LogRateLimitIntervalUSec"
	LogsDirectoryMode               Property = "LogsDirectoryMode"
	MainPID                         Property = "MainPID"
	ManagedOOMMemoryPressure        Property = "ManagedOOMMemoryPressure"
	ManagedOOMMemoryPressureLimit   Property = "ManagedOOMMemoryPressureLimit"
	ManagedOOMPreference            Property = "ManagedOOMPreference"
	ManagedOOMSwap                  Property = "ManagedOOMSwap"
	MemoryAccounting                Property = "MemoryAccounting"
	MemoryCurrent                   Property = "MemoryCurrent"
	MemoryDenyWriteExecute          Property = "MemoryDenyWriteExecute"
	MemoryHigh                      Property = "MemoryHigh"
	MemoryLimit                     Property = "MemoryLimit"
	MemoryLow                       Property = "MemoryLow"
	MemoryMax                       Property = "MemoryMax"
	MemoryMin                       Property = "MemoryMin"
	MemorySwapMax                   Property = "MemorySwapMax"
	MountAPIVFS                     Property = "MountAPIVFS"
	NFileDescriptorStore            Property = "NFileDescriptorStore"
	NRestarts                       Property = "NRestarts"
	NUMAPolicy                      Property = "NUMAPolicy"
	Names                           Property = "Names"
	NeedDaemonReload                Property = "NeedDaemonReload"
	Nice                            Property = "Nice"
	NoNewPrivileges                 Property = "NoNewPrivileges"
	NonBlocking                     Property = "NonBlocking"
	NotifyAccess                    Property = "NotifyAccess"
	OOMPolicy                       Property = "OOMPolicy"
	OOMScoreAdjust                  Property = "OOMScoreAdjust"
	OnFailureJobMode                Property = "OnFailureJobMode"
	PIDFile                         Property = "PIDFile"
	Perpetual                       Property = "Perpetual"
	PrivateDevices                  Property = "PrivateDevices"
	PrivateIPC                      Property = "PrivateIPC"
	PrivateMounts                   Property = "PrivateMounts"
	PrivateNetwork                  Property = "PrivateNetwork"
	PrivateTmp                      Property = "PrivateTmp"
	PrivateUsers                    Property = "PrivateUsers"
	ProcSubset                      Property = "ProcSubset"
	ProtectClock                    Property = "ProtectClock"
	ProtectControlGroups            Property = "ProtectControlGroups"
	ProtectHome                     Property = "ProtectHome"
	ProtectHostname                 Property = "ProtectHostname"
	ProtectKernelLogs               Property = "ProtectKernelLogs"
	ProtectKernelModules            Property = "ProtectKernelModules"
	ProtectKernelTunables           Property = "ProtectKernelTunables"
	ProtectProc                     Property = "ProtectProc"
	ProtectSystem                   Property = "ProtectSystem"
	RefuseManualStart               Property = "RefuseManualStart"
	RefuseManualStop                Property = "RefuseManualStop"
	ReloadResult                    Property = "ReloadResult"
	RemainAfterExit                 Property = "RemainAfterExit"
	RemoveIPC                       Property = "RemoveIPC"
	Requires                        Property = "Requires"
	Restart                         Property = "Restart"
	RestartKillSignal               Property = "RestartKillSignal"
	RestartUSec                     Property = "RestartUSec"
	RestrictNamespaces              Property = "RestrictNamespaces"
	RestrictRealtime                Property = "RestrictRealtime"
	RestrictSUIDSGID                Property = "RestrictSUIDSGID"
	Result                          Property = "Result"
	RootDirectoryStartOnly          Property = "RootDirectoryStartOnly"
	RuntimeDirectoryMode            Property = "RuntimeDirectoryMode"
	RuntimeDirectoryPreserve        Property = "RuntimeDirectoryPreserve"
	RuntimeMaxUSec                  Property = "RuntimeMaxUSec"
	SameProcessGroup                Property = "SameProcessGroup"
	SecureBits                      Property = "SecureBits"
	SendSIGHUP                      Property = "SendSIGHUP"
	SendSIGKILL                     Property = "SendSIGKILL"
	Slice                           Property = "Slice"
	StandardError                   Property = "StandardError"
	StandardInput                   Property = "StandardInput"
	StandardOutput                  Property = "StandardOutput"
	StartLimitAction                Property = "StartLimitAction"
	StartLimitBurst                 Property = "StartLimitBurst"
	StartLimitIntervalUSec          Property = "StartLimitIntervalUSec"
	StartupBlockIOWeight            Property = "StartupBlockIOWeight"
	StartupCPUShares                Property = "StartupCPUShares"
	StartupCPUWeight                Property = "StartupCPUWeight"
	StartupIOWeight                 Property = "StartupIOWeight"
	StateChangeTimestamp            Property = "StateChangeTimestamp"
	StateChangeTimestampMonotonic   Property = "StateChangeTimestampMonotonic"
	StateDirectoryMode              Property = "StateDirectoryMode"
	StatusErrno                     Property = "StatusErrno"
	StopWhenUnneeded                Property = "StopWhenUnneeded"
	SubState                        Property = "SubState"
	SuccessAction                   Property = "SuccessAction"
	SyslogFacility                  Property = "SyslogFacility"
	SyslogLevel                     Property = "SyslogLevel"
	SyslogLevelPrefix               Property = "SyslogLevelPrefix"
	SyslogPriority                  Property = "SyslogPriority"
	SystemCallErrorNumber           Property = "SystemCallErrorNumber"
	TTYReset                        Property = "TTYReset"
	TTYVHangup                      Property = "TTYVHangup"
	TTYVTDisallocate                Property = "TTYVTDisallocate"
	TasksAccounting                 Property = "TasksAccounting"
	TasksCurrent                    Property = "TasksCurrent"
	TasksMax                        Property = "TasksMax"
	TimeoutAbortUSec                Property = "TimeoutAbortUSec"
	TimeoutCleanUSec                Property = "TimeoutCleanUSec"
	TimeoutStartFailureMode         Property = "TimeoutStartFailureMode"
	TimeoutStartUSec                Property = "TimeoutStartUSec"
	TimeoutStopFailureMode          Property = "TimeoutStopFailureMode"
	TimeoutStopUSec                 Property = "TimeoutStopUSec"
	TimerSlackNSec                  Property = "TimerSlackNSec"
	Transient                       Property = "Transient"
	Type                            Property = "Type"
	UID                             Property = "UID"
	UMask                           Property = "UMask"
	UnitFilePreset                  Property = "UnitFilePreset"
	UnitFileState                   Property = "UnitFileState"
	UtmpMode                        Property = "UtmpMode"
	WantedBy                        Property = "WantedBy"
	WatchdogSignal                  Property = "WatchdogSignal"
	WatchdogTimestampMonotonic      Property = "WatchdogTimestampMonotonic"
	WatchdogUSec                    Property = "WatchdogUSec"
)