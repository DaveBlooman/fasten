package appmeta

type LightSailInstance struct {
	Instances []struct {
		Arn           string `json:"arn"`
		BlueprintID   string `json:"blueprintId"`
		BlueprintName string `json:"blueprintName"`
		BundleID      string `json:"bundleId"`
		CreatedAt     string `json:"createdAt"`
		Hardware      struct {
			CPUCount string `json:"cpuCount"`
			Disks    []struct {
				Arn             string `json:"arn"`
				AttachedTo      string `json:"attachedTo"`
				AttachmentState string `json:"attachmentState"`
				CreatedAt       string `json:"createdAt"`
				GbInUse         string `json:"gbInUse"`
				Iops            string `json:"iops"`
				IsAttached      string `json:"isAttached"`
				IsSystemDisk    string `json:"isSystemDisk"`
				Location        struct {
					AvailabilityZone string `json:"availabilityZone"`
					RegionName       string `json:"regionName"`
				} `json:"location"`
				Name         string `json:"name"`
				Path         string `json:"path"`
				ResourceType string `json:"resourceType"`
				SizeInGb     string `json:"sizeInGb"`
				SupportCode  string `json:"supportCode"`
			} `json:"disks"`
			RAMSizeInGb string `json:"ramSizeInGb"`
		} `json:"hardware"`
		Ipv6Address string `json:"ipv6Address"`
		IsStaticIP  string `json:"isStaticIp"`
		Location    struct {
			AvailabilityZone string `json:"availabilityZone"`
			RegionName       string `json:"regionName"`
		} `json:"location"`
		Name       string `json:"name"`
		Networking struct {
			MonthlyTransfer struct {
				GbPerMonthAllocated string `json:"gbPerMonthAllocated"`
			} `json:"monthlyTransfer"`
			Ports []struct {
				AccessDirection string `json:"accessDirection"`
				AccessFrom      string `json:"accessFrom"`
				AccessType      string `json:"accessType"`
				CommonName      string `json:"commonName"`
				FromPort        string `json:"fromPort"`
				Protocol        string `json:"protocol"`
				ToPort          string `json:"toPort"`
			} `json:"ports"`
		} `json:"networking"`
		PrivateIPAddress string `json:"privateIpAddress"`
		PublicIPAddress  string `json:"publicIpAddress"`
		ResourceType     string `json:"resourceType"`
		SSHKeyName       string `json:"sshKeyName"`
		State            struct {
			Code string `json:"code"`
			Name string `json:"name"`
		} `json:"state"`
		SupportCode string `json:"supportCode"`
		Username    string `json:"username"`
	} `json:"instances"`
	NextPageToken string `json:"nextPageToken"`
}
