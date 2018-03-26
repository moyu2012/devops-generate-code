package service

import (
	"os/exec"
	"strconv"
	"time"
)

type ApplicationConfigService struct {
}

func (s *ApplicationConfigService) CMDGenerateVue(name string) (file, dirName string, err error) {
	time := time.Now().Unix()
	dirName = name + "." + strconv.FormatInt(time, 10)
	println("mkdir -p /tmp/" + dirName)
	_, err = exec.Command("/bin/bash", "-c", "mkdir -p /tmp/"+dirName).Output()
	if err != nil {
		return "", "", err
	}
	println("cd /tmp/" + dirName + " && /usr/bin/vue-generate " + name)
	_, err = exec.Command("/bin/bash", "-c", "cd /tmp/"+dirName+" && /usr/bin/vue-generate "+name).Output()
	if err != nil {
		exec.Command("/bin/bash", "-c", "rm -fr /tmp/"+dirName).Output()
		return "", "", err
	}
	println("rm -fr /tmp/" + dirName + "/" + name + "/node_modules")
	_, err = exec.Command("/bin/bash", "-c", "rm -fr /tmp/"+dirName+"/"+name+"/node_modules").Output()
	if err != nil {
		exec.Command("/bin/bash", "-c", "rm -fr /tmp/"+dirName).Output()
		return "", "", err
	}
	println("cd /tmp/" + dirName + "/" + name + " && tar -zcvf generate-vue.tar.gz *")
	_, err = exec.Command("/bin/bash", "-c", "cd /tmp/"+dirName+"/"+name+" && tar -zcvf generate-vue.tar.gz *").Output()
	if err != nil {
		exec.Command("/bin/bash", "-c", "rm -fr /tmp/"+dirName).Output()
		return "", "", err
	}
	return "/tmp/" + dirName + "/" + name + "/generate-vue.tar.gz", dirName, nil
}

func (s *ApplicationConfigService) DeleteGZ(dirName string) error {
	println("rm -fr /tmp/" + dirName)
	_, err := exec.Command("/bin/bash", "-c", "rm -fr /tmp/"+dirName).Output()
	if err != nil {
		return err
	}
	return nil
}

func (s *ApplicationConfigService) CMDGenerateJava(name, group_id, artifact_id, version, package_name string) (file, dirName string, err error) {
	time := time.Now().Unix()
	dirName = name + "." + strconv.FormatInt(time, 10)
	println("mkdir -p /tmp/" + dirName)
	_, err = exec.Command("/bin/bash", "-c", "mkdir -p /tmp/"+dirName).Output()
	if err != nil {
		return "", "", err
	}
	println("cd /tmp/" + dirName + " && mvn archetype:generate -DgroupId=" + group_id + " -DartifactId=" + artifact_id + " -Dversion=" + version + "  -Dpackage=" + package_name + " -DarchetypeArtifactId=maven-archetype-quickstart  -DinteractiveMode=false")
	_, err = exec.Command("/bin/bash", "-c", "cd /tmp/"+dirName+" && mvn archetype:generate -DgroupId="+group_id+" -DartifactId="+artifact_id+" -Dversion="+version+"  -Dpackage="+package_name+" -DarchetypeArtifactId=maven-archetype-quickstart  -DinteractiveMode=false").Output()
	if err != nil {
		exec.Command("/bin/bash", "-c", "rm -fr /tmp/"+dirName).Output()
		return "", "", err
	}
	println("cd /tmp/" + dirName + "/" + artifact_id + " && tar -zcvf generate-java.tar.gz *")
	_, err = exec.Command("/bin/bash", "-c", "cd /tmp/"+dirName+"/"+artifact_id+" && tar -zcvf generate-java.tar.gz *").Output()
	if err != nil {
		exec.Command("/bin/bash", "-c", "rm -fr /tmp/"+dirName).Output()
		return "", "", err
	}
	return "/tmp/" + dirName + "/" + artifact_id + "/generate-java.tar.gz", dirName, nil
}

func (s *ApplicationConfigService) CMDGenerateAngular(name string) (file, dirName string, err error) {
	time := time.Now().Unix()
	dirName = name + "." + strconv.FormatInt(time, 10)
	println("mkdir -p /tmp/" + dirName)
	_, err = exec.Command("/bin/bash", "-c", "mkdir -p /tmp/"+dirName).Output()
	if err != nil {
		return "", "", err
	}
	println("cd /tmp/" + dirName + " && ng new " + name)
	_, err = exec.Command("/bin/bash", "-c", "cd /tmp/"+dirName+" && ng new "+name).Output()
	if err != nil {
		exec.Command("/bin/bash", "-c", "rm -fr /tmp/"+dirName).Output()
		return "", "", err
	}
	println("rm -fr /tmp/" + dirName + "/" + name + "/node_modules")
	_, err = exec.Command("/bin/bash", "-c", "rm -fr /tmp/"+dirName+"/"+name+"/node_modules").Output()
	if err != nil {
		exec.Command("/bin/bash", "-c", "rm -fr /tmp/"+dirName).Output()
		return "", "", err
	}
	println("cd /tmp/" + dirName + "/" + name + " && tar -zcvf generate-angular.tar.gz *")
	_, err = exec.Command("/bin/bash", "-c", "cd /tmp/"+dirName+"/"+name+" && tar -zcvf generate-angular.tar.gz *").Output()
	if err != nil {
		exec.Command("/bin/bash", "-c", "rm -fr /tmp/"+dirName).Output()
		return "", "", err
	}
	return "/tmp/" + dirName + "/" + name + "/generate-angular.tar.gz", dirName, nil
}
