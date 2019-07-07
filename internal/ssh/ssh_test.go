// Copyright © 2019 Alvaro Saurin
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ssh

import (
	"fmt"
	"io"
	"log"
	"time"

	"github.com/hashicorp/terraform/communicator/remote"
	"github.com/hashicorp/terraform/terraform"
)

type DummyOutput struct{}

func (_ DummyOutput) Output(s string) {
	fmt.Print(s)
}

type DummyCommunicator struct {
	startFunction func(cmd *remote.Cmd) error
}

func (_ DummyCommunicator) Connect(terraform.UIOutput) error {
	log.Printf("[DEBUG] [KUBEADM] DummyCommunicator: Connect()")
	return nil
}

func (_ DummyCommunicator) Disconnect() error {
	log.Printf("[DEBUG] [KUBEADM] DummyCommunicator: Disconnect()")
	return nil
}

func (_ DummyCommunicator) Timeout() time.Duration {
	log.Printf("[DEBUG] [KUBEADM] DummyCommunicator: Timeout()")
	return 1 * time.Hour
}

func (_ DummyCommunicator) ScriptPath() string {
	log.Printf("[DEBUG] [KUBEADM] DummyCommunicator: ScriptPath()")
	return ""
}

func (dc DummyCommunicator) Start(cmd *remote.Cmd) error {
	log.Printf("[DEBUG] [KUBEADM] DummyCommunicator: Start(%s)", cmd.Command)
	if dc.startFunction != nil {
		return dc.startFunction(cmd)
	}
	return nil
}

func (_ DummyCommunicator) Upload(string, io.Reader) error {
	log.Printf("[DEBUG] [KUBEADM] DummyCommunicator: Upload()")
	return nil
}

func (_ DummyCommunicator) UploadScript(string, io.Reader) error {
	log.Printf("[DEBUG] [KUBEADM] DummyCommunicator: UploadScript()")
	return nil
}

func (_ DummyCommunicator) UploadDir(string, string) error {
	log.Printf("[DEBUG] [KUBEADM] DummyCommunicator: UploadDir()")
	return nil
}
