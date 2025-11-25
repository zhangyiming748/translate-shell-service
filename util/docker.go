package util

import (
	"os"
	"runtime"
	"strings"
	"syscall"
)

// IsRunningInContainer 跨平台判定是否运行在容器中
func IsRunningInContainer() bool {
	switch runtime.GOOS {
	case "linux":
		return isLinuxContainer()
	case "darwin": // Mac OS
		return isMacContainer()
	case "windows":
		return isWindowsContainer()
	default:
		return false
	}
}

// isLinuxContainer Linux 容器检测（延续之前可靠方案，补充 Podman 支持）
func isLinuxContainer() bool {
	// 1. cgroup 检测（兼容 Docker/K8s/Containerd/LXC/Podman）
	cgroupPath := "/proc/self/cgroup"
	if content, err := os.ReadFile(cgroupPath); err == nil {
		cgroupTags := []string{"docker", "kubepods", "containerd", "lxc", "podman", "cri-o"}
		for _, tag := range cgroupTags {
			if strings.Contains(string(content), tag) {
				return true
			}
		}
	}

	// 2. 根目录 inode 检测
	if rootStat, err := os.Stat("/"); err == nil {
		if inode := rootStat.Sys().(*syscall.Stat_t).Ino; inode != 2 {
			return true
		}
	}

	// 3. 环境变量兜底
	containerEnvs := []string{"DOCKER_ENV", "KUBERNETES_SERVICE_HOST", "CONTAINER_ID", "PODMAN_ID"}
	for _, env := range containerEnvs {
		if os.Getenv(env) != "" {
			return true
		}
	}

	return false
}

// isMacContainer Mac OS 容器检测（主要针对 Docker Desktop 容器）
func isMacContainer() bool {
	// 1. Docker Desktop 容器会挂载 /Users 目录，且根目录 inode 非 2
	if rootStat, err := os.Stat("/"); err == nil {
		if inode := rootStat.Sys().(*syscall.Stat_t).Ino; inode != 2 {
			// 同时检查是否存在 Docker 特征文件
			if _, err := os.Stat("/.dockerenv"); err == nil {
				return true
			}
		}
	}

	// 2. 环境变量检测（Docker Desktop 会注入相关变量）
	macContainerEnvs := []string{"DOCKER_DESKTOP", "DOCKER_HOST", "CONTAINER_ID"}
	for _, env := range macContainerEnvs {
		if os.Getenv(env) != "" {
			return true
		}
	}

	return false
}

// isWindowsContainer Windows 容器检测（支持 Docker 容器、WSL2 容器）
func isWindowsContainer() bool {
	// 1. 检查 Windows 容器专属环境变量
	winContainerEnvs := []string{"COMSPEC", "CONTAINER_SID", "DOCKER_CONTAINER"}
	for _, env := range winContainerEnvs {
		if os.Getenv(env) != "" {
			return true
		}
	}

	// 2. 检查是否存在 Docker 容器特征文件（Docker for Windows 注入）
	if _, err := os.Stat("C:\\.dockerenv"); err == nil {
		return true
	}

	// 3. 检查 WSL2 容器（WSL2 容器会挂载 /mnt/c 目录，且存在 /etc/wsl.conf）
	if _, err := os.Stat("/etc/wsl.conf"); err == nil {
		if _, err := os.Stat("/mnt/c"); err == nil {
			return true
		}
	}

	return false
}
