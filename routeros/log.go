package routeros

import (
	"context"
	"os"

	"github.com/fatih/color"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ColorizedDebug Used to display provider log color messages.
// Please set the environment variable
func ColorizedDebug(ctx context.Context, msg string, args ...map[string]interface{}) {
	if _, set := os.LookupEnv("ROS_LOG_COLOR"); set {
		color.NoColor = false
	}
	tflog.Debug(ctx, color.GreenString(msg), args...)
}
