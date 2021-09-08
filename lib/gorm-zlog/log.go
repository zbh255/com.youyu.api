package gorm_zlog

import (
	gLog "gorm.io/gorm/logger"
)

// gorm logger interface
type InterFace gLog.Interface

type logger gLog.Config
