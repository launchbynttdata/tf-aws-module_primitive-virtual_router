# User needs to populate all the fields in <> in order to use this with `terraform plan|apply`

name          = "<virtual_router_name>"
app_mesh_name = "<service_mesh_name>"
listeners = [
  {
    port     = 8080
    protocol = "http"
  },
  {
    port     = 443
    protocol = "http"
  }
]
