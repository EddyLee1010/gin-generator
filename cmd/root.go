package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"log/slog"
	"net/http"
	"os"
	"time"
)

var rootCmd = &cobra.Command{
	Use:   "gin-generator",
	Short: "🕹️一款自动创建基于gin+GORM项目的cli小工具😂",
	Long: `💡 gin-generator🐔 可以快速把你建立起一个gin+GORM的项目
💡 包括自动生成项目结构、数据库模型+query、service+DTO、controller、router的代码`,
}

var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "生成命令 🔑help 获取使用方法",
}

// 获取版本号子命令
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "获取当前版本号",
	Run: func(cmd *cobra.Command, args []string) {
		// 使用http请求获取GitHub上的版本号
		releaseInfo, err := getLatestRelease("eddylee1010", "gin-generator")
		if err != nil {
			slog.Error("获取版本号失败")
			return
		}
		fmt.Printf("🫆当前版本号:%s\n", releaseInfo.TagName)
		fmt.Printf("🤕发布时间:%s\n", releaseInfo.PublishedAt.Local().Format("2006-01-02 15:04:05"))
		fmt.Printf("😆更新详情:%s\n", releaseInfo.HtmlUrl)
	},
}

type GitHubRelease struct {
	TagName     string    `json:"tag_name"`
	PublishedAt time.Time `json:"published_at"`
	HtmlUrl     string    `json:"html_url"`
}

func getLatestRelease(owner, repo string) (GitHubRelease, error) {
	var release GitHubRelease
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/releases/latest", owner, repo)
	resp, err := http.Get(url)
	if err != nil {
		return GitHubRelease{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return GitHubRelease{}, fmt.Errorf("GitHub API returned status: %s", resp.Status)
	}

	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		return GitHubRelease{}, err
	}
	return release, nil
}
func Execute() {
	rootCmd.AddCommand(genCmd) // 添加gen子命令
	rootCmd.AddCommand(versionCmd)

	if err := rootCmd.Execute(); err != nil {
		slog.Error("", err)
		os.Exit(1)
	}
}
