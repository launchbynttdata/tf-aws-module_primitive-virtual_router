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

resource "aws_appmesh_virtual_router" "this" {
  name       = var.name
  mesh_name  = var.app_mesh_name
  mesh_owner = var.app_mesh_owner_id

  spec {
    dynamic "listener" {
      for_each = [for listener_val in var.listeners : {
        port     = listener_val.port
        protocol = listener_val.protocol
      }]
      content {
        port_mapping {
          port     = listener.value.port
          protocol = listener.value.protocol
        }
      }
    }
  }

  tags = local.tags
}
