package vsphere

import (
	"github.com/emc-advanced-dev/unik/pkg/config"
	"github.com/emc-advanced-dev/unik/pkg/providers/vsphere/vsphereclient"
	"github.com/emc-advanced-dev/unik/pkg/state"
	"github.com/layer-x/layerx-commons/lxerrors"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

var VsphereStateFile = os.Getenv("HOME") + "/.unik/vsphere/state.json"
var VsphereImagesDirectory = "unik/vsphere/images/"
var VsphereVolumesDirectory = "unik/vsphere/volumes/"

const VsphereInstanceListener = "VsphereInstanceListener"

type VsphereProvider struct {
	config config.Vsphere
	state  state.State
	u      url.URL
}

func NewVsphereProvier(config config.Vsphere) (*VsphereProvider, error) {
	rawUrl := "https://" + config.VsphereUser + ":" + config.VspherePassword + "@" + strings.TrimSuffix(strings.TrimPrefix(strings.TrimPrefix(config.VsphereURL, "http://"), "https://"), "/sdk") + "/sdk"
	u, err := url.Parse(rawUrl)
	if err != nil {
		return nil, lxerrors.New("parsing vsphere url", err)
	}

	return &VsphereProvider{
		config: config,
		state:  state.NewBasicState(VsphereStateFile),
		u:      u,
	}
}

func (p *VsphereProvider) getClient() *vsphereclient.VsphereClient {
	return vsphereclient.NewVsphereClient(p.u, p.config.Datastore)
}

//just for consistency
func getInstanceDatastoreDir(instanceName string) string {
	return instanceName
}

func getImageDatastoreDir(imageName string) string {
	return filepath.Join(VsphereImagesDirectory, imageName + "/")
}

func getImageDatastorePath(imageName string) string {
	return filepath.Join(getImageDatastoreDir(imageName), "boot.vmdk")
}

func getVolumeDatastoreDir(volumeName string) string {
	return filepath.Join(VsphereVolumesDirectory, volumeName + "/")
}

func getVolumeDatastorePath(volumeName string) string {
	return filepath.Join(getVolumeDatastoreDir(volumeName), "data.vmdk")
}
