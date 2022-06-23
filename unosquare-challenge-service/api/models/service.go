package models

import "time"

type Service struct {
	Metadata Metadata `json:"metadata"`
	Spec     Spec     `json:"spec"`
}

type Metadata struct {
	Name              string    `json:"name"`
	Namespace         string    `json:"namespace"`
	Uid               string    `json:"uid"`
	ResourceVersion   string    `json:"resource_version"`
	CreationTimestamp time.Time `json:"creation_timestamp"`
}

type Spec struct {
	ClusterIP             string   `json:"cluster_ip"`
	ClusterIPs            []string `json:"cluster_ips"`
	Type                  string   `json:"type"`
	SessionAffinity       string   `json:"session_affinity"`
	IpFamilies            []string `json:"ip_families"`
	InternalTrafficPolicy string   `json:"internal_traffic_policy"`
}
