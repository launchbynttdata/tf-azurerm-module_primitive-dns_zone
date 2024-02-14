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

# Variables used for public DNS Zone
variable "resource_group_name" {
  description = "The resource group that resources will be created in."
  type        = string
}

variable "domain_names" {
  description = "The list of domain names of the DNS Zones."
  type        = list(string)
}

variable "tags" {
  description = "A mapping of tags to assign to the resource."
  type        = map(string)
  default     = {}
}

variable "resource_name_tag" {
  description = "The tag to use for the resource name."
  type        = string
  default     = null
}
