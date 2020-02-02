package main

type Configuration struct {
	api_version string;
	server_port string;
	find_country_api string;
	data_store_path  string;
}

func (this *Configuration) ApiVersion() string {
	return "asd"
}