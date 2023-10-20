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

variable "name" {
  description = "Name to use for the virtual router. Must be between 1 and 255 characters in length."
  type        = string
}

variable "app_mesh_name" {
  description = "Name of the service mesh in which to create the virtual router. Must be between 1 and 255 characters in length."
  type        = string
}

variable "listeners" {
  description = "Listeners that the virtual router is expected to receive inbound traffic from. Currently only one listener is supported per virtual router."
  type        = any
  default     = null
}

variable "tags" {
  description = "A map of custom tags to be attached to this resource"
  type        = map(string)
  default     = {}
}
