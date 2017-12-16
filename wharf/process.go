package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"path"
	"io/ioutil"
	"strconv"
)

const fullNS = syscall.CLONE_NEWNS|syscall.CLONE_NEWUTS|syscall.CLONE_NEWIPC|syscall.CLONE_NEWPID|syscall.CLONE_NEWUSER|syscall.CLONE_NEWNET

const cgroupMemoryHierarchyMount = "/sys/fs/cgroup/memory"

func main() {
	fmt.Println("欢迎体验Wharf容器平台！")

/*	fmt.Println("===== namespace =====")
	cmd := exec.Command("sh")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags:fullNS,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
*/

	fmt.Println("====== cgroups ======")
	if os.Args[0] == "/proc/self/exe" {
		fmt.Println("===== container =====")
		fmt.Println("current pid is", syscall.Getpid())
		cmd := exec.Command("sh", "-c", `stress --vm-bytes 256m --vm-keep -m 1 --timeout 30`)
		cmd.SysProcAttr = &syscall.SysProcAttr{}
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Println("failed to run command in container")
			fmt.Println(err)
			os.Exit(1)
		}
	}

	cmd := exec.Command("/proc/self/exe")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags:fullNS,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err != nil {
		fmt.Println("failed to fork container process")
		fmt.Println(err)
		os.Exit(1)
	} else {
		fmt.Println("the global pid of container process is", cmd.Process.Pid)
		fmt.Println("在系统默认创建的、挂在了memory subsystem的hierarchy上创建cgroup")
		memoryLimit := path.Join(cgroupMemoryHierarchyMount, "testmemorylimit")
		os.Mkdir(memoryLimit, 0755)
		fmt.Println("将container process加入到此cgroup")
		ioutil.WriteFile(path.Join(memoryLimit, "tasks"), []byte(strconv.Itoa(cmd.Process.Pid)),0644)
		fmt.Println("限制memory")
		ioutil.WriteFile(path.Join(memoryLimit, "memory.limit_in_bytes"), []byte("128m"), 0644)
	}
	cmd.Process.Wait()

	fmt.Println("======== git ========")

/*
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
	}
*/
}
