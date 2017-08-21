package main

import (
	"fmt"
	"os"
	"bytes"
	"os/exec"
	"errors"
	"bufio"
	"io"
)

func main(){
	fmt.Println(os.Args[0])
	out, err := OpCommandDirect("ls /root")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(out)
}

//OpCommandDirect :Use the shell command to get the result
func OpCommandDirect(cmd string) (string, error) {
	fmt.Printf("%s", cmd)
	outputs, err := OpCommandWithStdErr("/bin/bash", "-c", cmd)
	if err != nil {
		return "", errors.New(err.Error() + ": " + outputs)
	}
	return outputs, nil
}

//OpCommandDirectQuiet :Quietly  use the shell command to get the result
func OpCommandDirectQuiet(cmd string) (string, error) {
	//log.Debugf(cmd)
	outputs, err := OpCommandWithStdErr("/bin/bash", "-c", cmd)
	if err != nil {
		return "", errors.New(err.Error() + ": " + outputs)
	}
	return outputs, nil
}

//OpCommandWithStdin :Use the shell command to get the result,and put the result to somewhere
func OpCommandWithStdin(cmdstr string, in io.Reader) error {
	cmd := exec.Command("/bin/bash", "-c", cmdstr)
	cmd.Stdin = in
	b, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("%v (%s)", err, b)
	}

	return nil
}

//OpCommandDirectByte :Quietly  use the shell command to get the result
func OpCommandDirectByte(cmd string) ([]byte, error) {
	fmt.Printf(cmd)
	retByte, err := OpCommandWithByteErr("/bin/bash", "-c", cmd)
	return retByte, err
}

func OpCommandWithStdErr(bash string, arg ...string) (string, error) {
	//var stdout bytes.Buffer
	//var stderr bytes.Buffer
	//cmd := exec.Command(bash, arg...)
	//cmd.Stderr = &stderr
	//cmd.Stdout = &stdout
	//err := cmd.Run()
	//if err != nil {
	//	return string(stderr.Bytes()) + " " + string(stdout.Bytes()), err
	//}
	//return string(stdout.Bytes()), nil
	stdout, err := OpCommandWithByteErr(bash, arg...)
	return string(stdout), err
}

//OpCommandWithByteErr :this op will return stderr , op err and stdout.
func OpCommandWithByteErr(bash string, arg ...string) ([]byte, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(bash, arg...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return stderr.Bytes(), err
	}

	return stdout.Bytes(), nil
}

// OpCommandWithPipe : this op will return pipe handler, stdout and err
func OpCommandWithPipe(cmdstr string) (*bufio.Reader, *io.ReadCloser, error) {
	cmd := exec.Command("/bin/bash", "-c", cmdstr)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Printf("err:%s", err)
		return nil, nil, err
	}

	inputReader := bufio.NewReader(stdout)
	err = cmd.Start()
	if err != nil {
		fmt.Printf("Run %s failed, error: %v", cmdstr, err)
		stdout.Close()
		return nil, nil, err
	}
	go cmd.Wait()
	return inputReader, &stdout, nil
}
