package terraform

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-version"
	"github.com/hashicorp/hc-install/product"
	"github.com/hashicorp/hc-install/releases"
	"github.com/hashicorp/terraform-exec/tfexec"
	log "github.com/sirupsen/logrus"
	// TODO: How to get rid of this import? https://github.com/unravellingtechnologies/unravelling-cloud/issues/1
	_ "github.com/unravellingtechnologies/unravelling-cloud/cli/lib/logging"
	"net/http"
)

var tf *tfexec.Terraform

var Router *mux.Router

func initRouter() {
	Router = mux.NewRouter()
	Router.HandleFunc("/plan", PlanHandler)
}

func init() {
	initRouter()

	log.Infof("Initializing Terraform...")

	installer := &releases.ExactVersion{
		Product: product.Terraform,
		Version: version.Must(version.NewVersion("1.3.7")),
	}

	execPath, err := installer.Install(context.Background())
	if err != nil {
		log.Fatalf("error installing Terraform: %s", err)
	} else {
		log.Infof("Terraform installed at %s", execPath)
	}

	workingDir := "/Users/tiago/Projects/unravelling/aws-developer-associate/infrastructure"
	tf, err = tfexec.NewTerraform(workingDir, execPath)
	if err != nil {
		log.Fatalf("error running NewTerraform: %s", err)
	} else {
		log.Infof("Terraform initialized at working directory%s", workingDir)
	}
}

func Init() {
	log.Info("Running Terraform Init...")
	err := tf.Init(context.Background(), tfexec.Upgrade(true))
	if err != nil {
		log.Fatalf("error running Init: %s", err)
	} else {
		log.Infof("Terraform init completed successfully!")
	}
}

func Plan() {
	log.Info("Running Terraform Plan...")

}

func PlanHandler(w http.ResponseWriter, r *http.Request) {
	Plan()
	_, err := w.Write([]byte("Planning Terraform execution"))
	if err != nil {
		return
	}
}
